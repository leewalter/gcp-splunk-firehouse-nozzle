package main

import (
	"flag"
	"os"
	"os/signal"
	"syscall"

	"code.cloudfoundry.org/cflager"
	"github.com/cloudfoundry-community/splunk-firehose-nozzle/splunknozzle"
)

var (
	version string
	branch  string
	commit  string
	buildos string
)

func main() {
	cflager.AddFlags(flag.CommandLine)

	logger, _ := cflager.New("splunk-nozzle-logger")
	logger.Info("Running splunk-firehose-nozzle")

	shutdownChan := make(chan os.Signal, 2)
	signal.Notify(shutdownChan, syscall.SIGINT, syscall.SIGTERM)

	config := splunknozzle.NewConfigFromCmdFlags(version, branch, commit, buildos)

	splunkNozzle := splunknozzle.NewSplunkFirehoseNozzle(config)
	err := splunkNozzle.Run(shutdownChan, logger)
	if err != nil {
		logger.Error("Failed to run splunk-firehose-nozzle", err)
	}
}
