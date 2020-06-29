package internal

import (
	"os"

	"github.com/spf13/viper"
	"github.com/tendermint/tendermint/libs/cli"
	"github.com/tendermint/tendermint/libs/log"
)

var logger log.Logger

func init() {
	loadLogger()
}

func loadLogger() {
	logger = log.NewTMLogger(log.NewSyncWriter(os.Stdout))
	if viper.GetBool(cli.TraceFlag) {
		logger = log.NewTracingLogger(logger)
	}
}

func GetLogger() log.Logger {
	return logger
}
