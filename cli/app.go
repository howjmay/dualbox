package cli

import (
	"github.com/urfave/cli/v2"
)

func App() *cli.App {
	return &cli.App{
		Name:  "dualbox",
		Usage: "encrypt files into one encrypted file and decrypt them depends on the given key",
		Commands: []*cli.Command{
			Enc(),
			Dec(),
		},
	}
}
