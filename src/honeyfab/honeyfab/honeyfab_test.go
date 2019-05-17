package honeyfab_test

import (
	"io/ioutil"
	"os"
	"path/filepath"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/njbennett/honeytail-release/src/honeyfab/honeyfab"
)

var _ = Describe("Honeyfab", func() {
	var (
		fixturesDir string
		logPaths    []string
	)

	BeforeEach(func() {
		workingDir, err := os.Getwd()
		Expect(err).NotTo(HaveOccurred())

		fixturesDir = filepath.Join(workingDir, "fixtures")

		logPaths = []string{
			"/var/vcap/sys/log/web/web.stdout.log",
			"/var/vcap/sys/log/ncp/ncp.stdout.log",
			"/var/vcap/sys/log/bpm/bpm.stdout.log",
			"/var/vcap/sys/log/uaa/uaa.stdout.log",
		}
	})

	Describe("FindLogs", func() {
		It("finds all structured logs in the given directory", func() {
			logPaths, err := honeyfab.FindLogs(fixturesDir)
			Expect(err).NotTo(HaveOccurred())

			Expect(logPaths).To(ConsistOf(
				filepath.Join(fixturesDir, "dir1/structured.log"),
				filepath.Join(fixturesDir, "dir2/structured.log"),
				filepath.Join(fixturesDir, "dir3/structured.log"),
			))
		})
	})

	Describe("WriteHoneytailConf", func() {
		var (
			baseHoneytailConfPath   string
			newHoneytailConfPath    string
			expectedHoneytailConf   string
			brokenHoneytailConfPath string
		)

		BeforeEach(func() {
			baseHoneytailConfPath = filepath.Join(fixturesDir, "honeytail.conf")
			newHoneytailConf, err := ioutil.TempFile("", "honeytail-*.conf")
			Expect(err).NotTo(HaveOccurred())
			newHoneytailConfPath = newHoneytailConf.Name()

			expectedHoneytailConf = `[Required Options]
ParserName = json

WriteKey = some-write-key

Dataset = some-dataset

LogFiles = /var/vcap/sys/log/web/web.stdout.log
LogFiles = /var/vcap/sys/log/ncp/ncp.stdout.log
LogFiles = /var/vcap/sys/log/bpm/bpm.stdout.log
LogFiles = /var/vcap/sys/log/uaa/uaa.stdout.log

[ Application Options ]
# add some BOSH-provided fields to all events
AddFields = bosh.spec.deployment=concourse
`
		})

		AfterEach(func() {
		})

		It("writes new honeytail conf to the provided path", func() {
			err := honeyfab.WriteHoneytailConf(baseHoneytailConfPath, newHoneytailConfPath, logPaths)
			Expect(err).NotTo(HaveOccurred())

			contents, err := ioutil.ReadFile(newHoneytailConfPath)
			Expect(err).NotTo(HaveOccurred())
			Expect(string(contents)).To(Equal(expectedHoneytailConf))
		})

		Context("when the base honeytail conf does not exist", func() {
			BeforeEach(func() {
				brokenHoneytailConfPath = filepath.Join(fixturesDir, "honeytail-nonexistent.conf")
			})

			It("returns an error", func() {
				err := honeyfab.WriteHoneytailConf(brokenHoneytailConfPath, newHoneytailConfPath, logPaths)
				Expect(err).To(HaveOccurred())

				Expect(err.Error()).To(ContainSubstring(brokenHoneytailConfPath))
				Expect(err.Error()).To(ContainSubstring("no such file or directory"))
			})
		})

		Context("when the base honeytail conf is not a parseable Go template", func() {
			BeforeEach(func() {
				brokenHoneytailConfPath = filepath.Join(fixturesDir, "honeytail-broken.conf")
			})

			It("returns an error", func() {
				err := honeyfab.WriteHoneytailConf(brokenHoneytailConfPath, newHoneytailConfPath, logPaths)
				Expect(err).To(HaveOccurred())

				Expect(err.Error()).To(ContainSubstring(brokenHoneytailConfPath))
				Expect(err.Error()).To(ContainSubstring("is not a valid Go template"))
			})
		})

		Context("when the new honeytail conf cannot be created", func() {
			BeforeEach(func() {
				brokenHoneytailConfPath = "/sbin/no-permission-to-write-here"
			})

			It("returns an error", func() {
				err := honeyfab.WriteHoneytailConf(baseHoneytailConfPath, brokenHoneytailConfPath, logPaths)
				Expect(err).To(HaveOccurred())

				Expect(err.Error()).To(ContainSubstring(brokenHoneytailConfPath))
				Expect(err.Error()).To(ContainSubstring("creating new honeytail.conf file"))
				Expect(err.Error()).To(ContainSubstring("operation not permitted"))
			})
		})
	})
})
