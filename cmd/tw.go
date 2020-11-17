/*
Copyright © 2020 Chris

*/
package cmd

import (
    "github.com/gwd0715/mdt/util"
    "github.com/spf13/cobra"
)

// twCmd represents the tw command
var twCmd = &cobra.Command{
    Use:   "tw [path]",
    Short: "convert markdown file from zh-ch to zh-tw",
    Run: func(cmd *cobra.Command, args []string) {
        for _, filePath := range args {
            util.Translate(filePath, "tw")
        }
    },
}

func init() {
    rootCmd.AddCommand(twCmd)
}
