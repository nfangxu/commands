package main

import (
	"github.com/nfangxu/tools/internal/starter"
	"github.com/spf13/cobra"
	"log"
)

var (
	version = "v0.1.0"

	rootCmd = &cobra.Command{
		Use:     "fx",
		Short:   "fx: fangx tools.",
		Long:    `fx: fangx tools.`,
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
