package main

import (
	"framework/cfgargs"
	"framework/logger"
)

var (
	BuildVersion, BuildTime, BuildUser, BuildMachine string
)

func init() {
}

func main() {
	build := &cfgargs.Build{
		BuildVersion: BuildVersion,
		BuildTime:    BuildTime,
		BuildUser:    BuildUser,
		BuildMachine: BuildMachine,
	}
	srvConfig, err := cfgargs.InitSrvCfg(build, nil)
	if err != nil {
		panic(err)
	}

	logger.InitLogger(srvConfig)

	logger.Debug("Test for framework")
}
