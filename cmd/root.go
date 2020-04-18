package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"
)

var (
	// Flags

	// Port specify which port the server should be hosted on
	Port string
	// Tunnel if a tunnel is needed to expose the server to the internet
	Tunnel bool
	// Root cmd
	rootCmd = &cobra.Command{
		Use:   "iplogger",
		Short: "Start up a server whos main purpose is to log client IPs",
		Long:  `Start up a server whos main purpose is to log client IPs`,
		Args:  cobra.ArbitraryArgs,
		Run: func(cmd *cobra.Command, args []string) {
			// print the usage if any arguments are passed
			if len(args) > 0 {
				if err := cmd.Usage(); err != nil {
					log.Fatal(err)
				}

				return
			}

			// lets start logging client IPs
			LogIP(args)
		},
	}
)

// Execute the command
func Execute() {
	// Persistent Flags
	rootCmd.PersistentFlags().BoolVarP(&Tunnel, "tunnel", "t", false, "Specify if a reverse ssh tunnel should be used to expose the webserver to the internet")
	rootCmd.PersistentFlags().StringVar(&Port, "port", "", "Specify which port the webserver should be hosted on")

	// Execute the command :)
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
