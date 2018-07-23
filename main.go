package main

// func (cmdHandler *cmdHandler) handleAction(c *cli.Context) error {
// 	fmt.Println("TTL", c.String("TTL"))
// 	fmt.Println("Command:", c.Args()[0])
// 	fmt.Println("Args:", c.Args()[1:])

// 	// args := c.Args()

// 	// stdout, err := executeCommand(
// 	// 	args[0],
// 	// 	args[1:],
// 	// 	os.Environ(),
// 	// )

// 	return nil
// }

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

// func main() {
// 	setupCmd(
// 		os.Args,
// 		os.Stdout,
// 		os.Stderr,
// 	)
// }
