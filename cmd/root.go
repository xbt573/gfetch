package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var verbose bool
var limit int

func init() {
	rootCmd.AddCommand(&gelbooruCmd)
	rootCmd.PersistentFlags().BoolVarP(&verbose,
		"verbose",
		"v",
		false,
		"verbose output",
	)

	rootCmd.PersistentFlags().IntVarP(&limit,
		"limit",
		"l",
		100,
		"post limit",
	)
}

var rootCmd = &cobra.Command{
	Use:              "gfetch [subcommand] [tags]",
	Short:            "gfetch is batch photo downloader from many boorus",
	Long:             "gfetch is batch photo downloader from many boorus",
	TraverseChildren: true,
	Run: func(cmd *cobra.Command, args []string) {
		gelbooruCmd.Run(cmd, args)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
