package main

import (
	"log"

	"github.com/farmani/sharebuy/cmd/api"
	cobra "github.com/spf13/cobra"
)

func main() {
	const description = `Sharebuy Application`
	var rootCmd = &cobra.Command{Use: "app", Short: description, Long: description}
	rootCmd.AddCommand(
		api.Server{}.Command(nil),
	)

	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("failed to execute root command:\n%v", err)
	}
}
