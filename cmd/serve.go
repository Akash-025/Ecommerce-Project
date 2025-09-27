package cmd

import (
	"practice/config"
	"practice/rest"
)

func Serve() {

	cnf := config.GetConfig()

	rest.Start(cnf)
}
