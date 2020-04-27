package cmd

import (
  "github.com/spf13/cobra"
  "health/api"
)

var apiCmd = &cobra.Command{
  Use:   "api",
  Short: "Run Api",
  Long:  `Run Api server SIRES Tasks`,
  Run: func(cmd *cobra.Command, args []string) {
    api.StartApi()
  },
}

func init() {
  rootCmd.AddCommand(apiCmd)
}
