// Copyright © 2022 Thiago Sousa Bastos <thiagosbastos@live.com>

package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "splash-cli",
	Short: "A Sprint Planning CLI written in Go",
	Long:  `A Sprint Planning CLI written in Go.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {}
