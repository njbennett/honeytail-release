package honeyfab_test

import (
	"io/ioutil"
	"os"
	"path/filepath"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/pivotal/honeytail-release/src/honeyfab/honeyfab"
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
			baseHoneytailConfPath string
			newHoneytailConfPath  string
			expectedHoneytailConf string
		)

		BeforeEach(func() {
			baseHoneytailConfPath = filepath.Join(fixturesDir, "honeytail.conf")
			newHoneytailConf, err := ioutil.TempFile("", "honeytail-*.conf")
			Expect(err).NotTo(HaveOccurred())
			newHoneytailConfPath = newHoneytailConf.Name()

			expectedHoneytailConf = `[Required Options]
ParserName = <%= properties.parser %>

WriteKey = <%= properties.write_key %>

Dataset = <%= properties.dataset %>

LogFiles = /var/vcap/sys/log/web/web.stdout.log
LogFiles = /var/vcap/sys/log/ncp/ncp.stdout.log
LogFiles = /var/vcap/sys/log/bpm/bpm.stdout.log
LogFiles = /var/vcap/sys/log/uaa/uaa.stdout.log

[ Application Options ]
# add some BOSH-provided fields to all events
AddFields = bosh.spec.address=<%=spec.address%>
AddFields = bosh.spec.az=<%=spec.az%>
AddFields = bosh.spec.bootstrap=<%=spec.bootstrap%>
AddFields = bosh.spec.deployment=<%=spec.deployment%>
AddFields = bosh.spec.id>=<%=spec.id%>
AddFields = bosh.spec.index=<%=spec.index%>
AddFields = bosh.spec.name=<%=spec.name%>
AddFields = bosh.spec.network=<%=spec.network%>
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
	})
})
