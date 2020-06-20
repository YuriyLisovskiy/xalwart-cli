package commands

import (
	"flag"
)

const newAppCmdName = "new-app"

var (
	NewAppCmd = flag.NewFlagSet(newAppCmdName, flag.ExitOnError)

	naNameFlag = NewAppCmd.String("name", "", "Name of a new application")
)

// TODO: implement new-app command
