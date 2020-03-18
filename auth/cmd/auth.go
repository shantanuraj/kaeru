package main

import (
	"log"
	"os"
	"sixth-io/kaeru/auth/server"
	"sixth-io/kaeru/db"
	"strings"

	"github.com/urfave/cli/v2"
)

func main() {
	app := cli.NewApp()
	app.Name = AppName
	app.Usage = AppUsage
	app.Version = AppVersion
	app.Flags = flags
	app.Action = run
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal("Unexpected runtime error", err)
	}
}

func run(c *cli.Context) error {
	dbAddress := strings.Trim(c.String(FlagDatabase), " \t\n")
	port := strings.Trim(c.String(FlagPort), "\t\n")
	dao := db.New(dbAddress)

	if err := dao.Start(); err != nil {
		log.Fatal("Could not connect to database")
	}
	defer dao.Close()

	server.Start(port, dao)
	return nil
}
