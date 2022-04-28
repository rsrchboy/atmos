package main

import (
	"github.com/cloudposse/atmos/cmd"
	u "github.com/cloudposse/atmos/pkg/utils"

	// be super-lazy and just load any .env
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		u.PrintErrorToStdErrorAndExit(err)
	}
}
