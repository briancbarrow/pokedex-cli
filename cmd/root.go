package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "pokedex",
	Short: "Pokedex CLI ",
	Long: `Pokedex CLI is a CLI tool to interact with the Pokedex API. It allows you to search for Pokemon by name, ID, or type. 
	You can also view a list of all Pokemon, or view a list of all Pokemon types.`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
