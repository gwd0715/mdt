/*
Copyright © 2020 Chris

*/
package cmd

import (
	"github.com/gwd0715/mdt/util"
	"github.com/spf13/cobra"
)

// cnCmd represents the cn command
var cnCmd = &cobra.Command{
	Use:   "cn",
	Short: "Convert markdown file from zh-tw to zh-cn",
	Run: func(cmd *cobra.Command, args []string) {
		for _, filepath := range args {
			util.Translate(filepath, "cn")
		}
	},
}

func init() {
	rootCmd.AddCommand(cnCmd)
}
