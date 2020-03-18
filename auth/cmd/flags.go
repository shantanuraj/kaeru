package main

import "github.com/urfave/cli/v2"

const (
	// AppVersion is the version of the app
	AppVersion = "0.0.1"
	// AppName is the name of the app
	AppName = "auth"
	// AppUsage is the help text for the app
	AppUsage = "Auth API server for kaeru"
	// FlagDatabase defines the flag name for specifying the db address
	FlagDatabase = "database"
	// FlagPort defines the flag name for specifying the port
	FlagPort = "port"
	// FlagVerbose is the flag name for turning on verbose mode
	FlagVerbose = "verbose"
)

var flags = []cli.Flag{
	&cli.StringFlag{
		Name:     FlagDatabase,
		Usage:    "configure database address",
		FilePath: "/etc/kaeru/auth_db",
	},

	&cli.StringFlag{
		Name:  FlagPort,
		Usage: "configure server port",
		Value: "4551",
	},
}
