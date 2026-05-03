package cmd

import (
	mathematics "calc/math"
	"strconv"
	"fmt"
	"github.com/spf13/cobra"
	"math"
)

var (
	// Whether Cos(x) accepts radian or degrees as input.
	cosUsesRad bool = false
	// Whether Sin(x) accepts radian or degree as input.
	sinUsesRad bool = false
	// Base of Log(x, b).
	base int
)

var cosCmd = &cobra.Command{
	Use: "cos",
	Short: "Cosine function",
	Args: cobra.ExactArgs(1),
	Run: func (cmd *cobra.Command, args []string)  {
		x, err := strconv.ParseFloat(args[0], 64)
		if err != nil {
			fmt.Println(err)
			return
		}
		if !cosUsesRad {
			x = math.Pi / 180 * x
		}
		result := mathematics.Cos(x)
		fmt.Println(result)
	},
}

var sinCmd = &cobra.Command{
	Use: "sin",
	Short: "Sine function",
	Args: cobra.ExactArgs(1),
	Run: func (cmd *cobra.Command, args []string)  {
		x, err := strconv.ParseFloat(args[0], 64)
		if err != nil {
			fmt.Println(err)
			return
		}
		if !sinUsesRad {
			x = math.Pi / 180 * x
		}
		result := mathematics.Sin(x)
		fmt.Println(result)
	},
}

var expCmd = &cobra.Command{
	Use: "exp",
	Short: "Exponential function",
	Args: cobra.ExactArgs(1),
	Run: func (cmd *cobra.Command, args []string)  {
		x, err := strconv.ParseFloat(args[0], 64)
		if err != nil {
			fmt.Println(err)
			return
		}
		result := mathematics.Exp(x)
		fmt.Println(result)
	},
}

var logCmd = &cobra.Command{
	Use: "log",
	Short: "Natural logarithm function",
	Args: cobra.ExactArgs(1),
	Run: func (cmd *cobra.Command, args []string)  {
		x, err := strconv.ParseFloat(args[0], 64)
		if err != nil {
			fmt.Println(err)
			return
		}
		result := mathematics.Log(x, base)
		fmt.Println(result)
	},
}

func init() {
	cosCmd.Flags().BoolVar(&cosUsesRad, "rad", false, "Indicates whether the input is degree or radian")
	sinCmd.Flags().BoolVar(&sinUsesRad, "rad", false, "Indicates whether the input is degree or radian")
	logCmd.Flags().IntVar(&base, "base", 0, "Defines a basis for the logarithm function | 0 stands for Natural Logarithm")
	rootCmd.AddCommand(cosCmd)
	rootCmd.AddCommand(sinCmd)
	rootCmd.AddCommand(expCmd)
	rootCmd.AddCommand(logCmd)
}