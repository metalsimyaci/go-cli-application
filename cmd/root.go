package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var ( rootCmd = &cobra.Command{
	Use: "go-cli-application",
	Short: "Example Go CLI Application",
	Long: "Example go CLI Application Long",
})

func Execute(){
	if err := rootCmd.Execute(); err !=nil {
		fmt.Fprintln(os.Stderr,err)
		os.Exit(1)
	}
}

func init(){
	rootCmd.CompletionOptions.DisableDefaultCmd = true
}