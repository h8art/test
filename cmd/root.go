package cmd

import (
  "fmt"
  "github.com/spf13/cobra"
  "os"
)

var rootCmd = &cobra.Command{
  Use:   "health",
  Short: "Rushydro.Health",
  Long:  `Rushydro.Health Application`,
}

func Execute() {
  if err := rootCmd.Execute(); err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
}