package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "mdscli",
	Short: "MDS CLI interacts with the Confluent Platform's Metadata Service",
	Long: `A command-line tool to manage Confluent Platform's role bindings
	from a YAML file using the Metadata Service API.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
