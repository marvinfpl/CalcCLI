package cmd

import (
	"strconv"
	"fmt"
	"github.com/spf13/cobra"
	"calc/math"
)

var additionCmd = &cobra.Command{
	Use: "add",
	Short: "Outputs the addition of two numbers",
	Args: cobra.ExactArgs(2),
	Run: func(cmd* cobra.Command, args []string) {
		a, err := strconv.ParseFloat(args[0], 64)
		if err != nil {
			fmt.Println("invalid number:", args[0])
			return
		}
		b, err := strconv.ParseFloat(args[1], 64)
		if err != nil {
			fmt.Println("invalid number:", args[1])
			return
		}
		result := mathematics.Add(a, b)
		fmt.Println(result)
	},
}

var substrationCmd = &cobra.Command{
	Use: "sub",
	Short: "Outputs the substraction of two numbers",
	Args: cobra.ExactArgs(2),
	Run: func(cmd* cobra.Command, args []string) {
		a, err := strconv.ParseFloat(args[0], 64)
		if err != nil {
			fmt.Println("invalid number:", args[0])
			return
		}
		b, err := strconv.ParseFloat(args[1], 64)
		if err != nil {
			fmt.Println("invalid number:", args[1])
			return
		}
		result := mathematics.Subtract(a, b)
		fmt.Println(result)
	},
}

var multiplyCmd = &cobra.Command{
	Use: "mul",
	Short: "Outputs the multiplication of two numbers",
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		a, err := strconv.ParseFloat(args[0], 64)
		if err != nil {
			fmt.Println("invalid number:", args[0])
			return
		}
		b, err := strconv.ParseFloat(args[1], 64)
		if err != nil {
			fmt.Println("invalid number:", args[1])
			return
		}
		result := mathematics.Multiply(a, b)
		fmt.Println(result)
	},
}

var divideCmd = &cobra.Command{
	Use: "div",
	Short: "Outputs the division of two numbers",
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		a, err := strconv.ParseFloat(args[0], 64)
		if err != nil {
			fmt.Println("invalid number:", args[0])
			return
		}
		b, err := strconv.ParseFloat(args[1], 64)
		if err != nil {
			fmt.Println("invalid number:", args[1])
			return
		}
		result, err := mathematics.Divide(a, b)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(result)
	},
}

func init() {
	rootCmd.AddCommand(additionCmd)
	rootCmd.AddCommand(substrationCmd)
	rootCmd.AddCommand(multiplyCmd)
	rootCmd.AddCommand(divideCmd)
}