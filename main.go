package main

import (
	"github.com/gpmgo/gopm/modules/cli"

)

const  useage = `mydocker is a simple container runtime implementation. The purpose of this project is to learn how docker works and how to write a docker by ourselves Enjoy it, just for fun.`

func main()  {
	app := cli.NewApp()
	app.Name = "myker"
	app.Usage = useage

	app.Commands = []cli.Command{
		initCommand,
		runCommand,
	}
}