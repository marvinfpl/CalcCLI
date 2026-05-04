package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	mathematics "calc/math"
)


var additionCmd = &cobra.Command{
	Use: "add",
	Short: "Outputs the addition of two numbers",
	Args: cobra.ExactArgs(2),
	RunE: func(cmd* cobra.Command, args []string) error {
		a, b, err := mathematics.ParseTwoOperands(args)
		if err != nil {
			return fmt.Errorf("Invalid arguments: %v", err)
		}
		result := mathematics.Add(a, b)
		fmt.Println(result)
		return nil
	},
}

var subtractionCmd = &cobra.Command{
	Use: "sub",
	Short: "Outputs the subtraction of two numbers",
	Args: cobra.ExactArgs(2),
	RunE: func(cmd* cobra.Command, args []string) error {
		a, b, err := mathematics.ParseTwoOperands(args)
		if err != nil {
			return fmt.Errorf("Invalid arguments: %v", err)
		}
		result := mathematics.Subtract(a, b)
		fmt.Println(result)
		return nil
	},
}

var multiplyCmd = &cobra.Command{
	Use: "mul",
	Short: "Outputs the multiplication of two numbers",
	Args: cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		a, b, err := mathematics.ParseTwoOperands(args)
		if err != nil {
			return fmt.Errorf("Invalid arguments: %v", err)
		}
		result := mathematics.Multiply(a, b)
		fmt.Println(result)
		return nil
	},
}

var divideCmd = &cobra.Command{
	Use: "div",
	Short: "Outputs the division of two numbers",
	Args: cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		a, b, err := mathematics.ParseTwoOperands(args)
		if err != nil {
			return fmt.Errorf("Invalid arguments: %v", err)
		}
		result, err := mathematics.Divide(a, b)
		if err != nil {
			return fmt.Errorf("Invalid arguments: %v", err)
		}
		fmt.Println(result)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(additionCmd)	
	rootCmd.AddCommand(subtractionCmd)
	rootCmd.AddCommand(multiplyCmd)
	rootCmd.AddCommand(divideCmd)
}