package cmd

import (
	"github.com/spf13/cobra"
	"fmt"
)

var rootCmd = &cobra.Command{
	Use: "calc",
	Short: "Advanced calculator for scientific purpose",
	Run: func (cmd *cobra.Command, args[] string)  {
		fmt.Print(  "__        __   _                           \n",
           			"\\ \\      / /__| | ___ ___  _ __ ___   ___  \n",
           			" \\ \\ /\\ / / _ \\ |/ __/ _ \\| '_ ` _ \\ / _ \\ \n",
           			"  \\ V  V /  __/ | (_| (_) | | | | | |  __/ \n",
           			"   \\_/\\_/ \\___|_|\\___\\___/|_| |_| |_|\\___| \n");
    
    	fmt.Print("\n");

    	fmt.Print(  "__  __      _____                       _         \n",
            		"|  \\/  |_ __|  ___|__  _ __  _ __   ___ | | ___   \n",
            		"| |\\/| | '__| |_ / _ \\| '_ \\| '_ \\ / _ \\| |/ _ \\  \n",
            		"| |  | | |_ |  _| (_) | |_) | |_) | (_) | | (_) |  \n",
            		"|_|  |_|_(_)|_|  \\___/| .__/| .__/ \\___/|_|\\___/  \n",
            		"                       |_|   |_|                    )\n");
	},
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}