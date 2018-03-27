package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "Sentinel"
	app.Usage = "Status checking tool for OpenHack - DevOps"
	//	app.UsageText = "This is the usage text" // default will be nice
	app.Version = "0.0.1"
	app.Commands = []cli.Command{
		{
			Name:    "init",
			Aliases: []string{"i"},
			Usage:   "Initialize Sentinel. Insert initial Data for an Open Hack",
			Action: func(c *cli.Context) error {
				fmt.Println("***** I do initialize!")
				fmt.Printf("Flag: %v \n", c.Bool("t"))
				return nil
			},
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:  "template, t",
					Usage: "Initialize template data set for Testing",
				},
			},
		},
		{
			Name:    "start",
			Aliases: []string{"s"},
			Usage:   "Start Sentinel App and monitor the endpoint",
			Action: func(c *cli.Context) error {
				fmt.Println("**** server started ***** ")
				return nil
			},
		},
	}

	app.Run(os.Args)
}
