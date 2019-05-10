package main

import (
	"log"
	"os/exec"

	"github.com/pivotal/honeytail-release/src/honeyfab/honeyfab"
)

func main() {
	logPaths, err := honeyfab.FindLogs("/var/vcap/sys/log")
	if err != nil {
		log.Fatalf("finding logs: %v", err)
	}

	err = honeyfab.WriteHoneytailConf(
		"/var/vcap/jobs/honeytail/config/honeytail.base.conf",
		"/var/vcap/jobs/honeytail/config/honeytail.conf",
		logPaths,
	)

	if err != nil {
		log.Fatalf("writing honeytail.conf: %v", err)
	}

	cmd := exec.Command("/var/vcap/packages/honeytail/bin/honeytail",
		"-c", "/var/vcap/jobs/honeytail/config/honeytail.conf")

	cmd.Run()
}
