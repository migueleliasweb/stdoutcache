package main

// import (
// 	"io"
// 	"log"

// 	cli "gopkg.in/urfave/cli.v1"
// )

// func setupCmd(args []string, stdoutWriter io.WriteCloser, stderrWriter io.WriteCloser) {
// 	app := cli.NewApp()

// 	app.Flags = []cli.Flag{
// 		cli.StringFlag{
// 			Name:   "TTL",
// 			Usage:  "The maximum amount of seconds the output will be cached",
// 			EnvVar: "STDOUTCACHE_TTL",
// 		},
// 	}

// 	handler := cmdHandler{
// 		stdoutWriter: stdoutWriter,
// 		stderrWriter: stderrWriter,
// 	}

// 	app.Action = handler.handleAction

// 	err := app.Run(args)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }
