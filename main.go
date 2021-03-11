package main

import (
    "github.com/nfangxu/commands/internal/starter"
    "github.com/spf13/cobra"
    "log"
)

var (
    version = "v0.1.0-beta01"

    rootCmd = &cobra.Command{
        Use:     "fangx",
        Short:   "fangx: An elegant toolkit for Go microservices.",
        Long:    `fangx: An elegant toolkit for Go microservices.`,
        Version: version,
    }
)

func init() {
    rootCmd.AddCommand(starter.Cmd)
}

func main() {
    if err := rootCmd.Execute(); err != nil {
        log.Fatal(err)
    }
}
