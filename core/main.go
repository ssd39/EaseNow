package main

import (
	"log"
	"os"

	"github.com/ssd39/easenow/core/app"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "easenow-core",
		Usage: "TEE core layer of protocol",
		Action: func(cCtx *cli.Context) error {
			seed, err := app.Init("/data/.seed")
			if err != nil {
				return err
			}
			return app.StartApi(seed)
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
