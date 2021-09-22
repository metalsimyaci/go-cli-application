package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init(){
	rootCmd.AddCommand(exampleCmd)
}

func example(){
	fmt.Println("My example output")
}

var(
	exampleCmd = &cobra.Command{
		Use: "example",
		Short: "Example command",
		Long: "Example command",
		Run: func(cmd *cobra.Command, args []string)  {
			example()
		},
	}
)