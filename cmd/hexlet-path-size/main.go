package main

import (
	"context"
	"fmt"
	"os"

	"github.com/urfave/cli/v3"

	"code"
)

func main() {
	cmd := &cli.Command{
		Name:      "hexlet-path-size",
		Usage:     "print size of a file or directory",
		ArgsUsage: "<path>",

		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "human",
				Aliases: []string{"H"},
				Usage:   "human-readable sizes (auto-select unit)",
			},
			&cli.BoolFlag{
				Name:    "all",
				Aliases: []string{"a"},
				Usage:   "include hidden files and directories",
			},
			&cli.BoolFlag{
				Name:    "recursive",
				Aliases: []string{"r"},
				Usage:   "recursive size of directories",
			},
		},

		Action: func(ctx context.Context, cmd *cli.Command) error {
			//We check to see if the path has been passed.
			if cmd.Args().Len() == 0 {
				return cli.ShowAppHelp(cmd)
			}

			path := cmd.Args().Get(0)

			// Get flag values.
			recursive := cmd.Bool("recursive")
			human := cmd.Bool("human")
			all := cmd.Bool("all")

			// Get the formatted size of the path.
			size, err := code.GetPathSize(path, recursive, human, all)
			if err != nil {
				return cli.Exit(err.Error(), 1)
			}

			fmt.Printf("%s\t%s\n", size, path)

			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
