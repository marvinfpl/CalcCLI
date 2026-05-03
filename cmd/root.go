package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "calc",
	Short: "Advanced calculator for scientific purpose",
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}