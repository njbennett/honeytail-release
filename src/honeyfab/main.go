package main

import (
	"log"

	"github.com/njbennett/honeytail-release/src/honeyfab/honeyfab"
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
}
