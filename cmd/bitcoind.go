/*
   Copyright © 2022 Johann Suarez <johann@gohash.tech>
*/

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"loader"
    "log"
	"server"
	"validation"
)

// rootCmd represents the base command when called without any subcommands
var (
	// These two vars store the provided inputs for the flags.
	auth_str string
	port_num uint16 = 4000 // Port number can go up to 65535.

	/*

    Information on the Use field of the Command struct ( taken from official documentation )

    Use is the one-line usage message.
    Recommended syntax is as follow:
      [ ] identifies an optional argument. Arguments that are not enclosed in brackets are required.
      ... indicates that you can specify multiple values for the previous argument.
      |   indicates mutually exclusive information. You can use the argument to the left of the separator or the
          argument to the right of the separator. You cannot use both arguments in a single use of the command.
      { } delimits a set of mutually exclusive arguments when one of the arguments is required. If the arguments are
          optional, they are enclosed in brackets ([ ]).
    Example: add [-F file | -D dir]... [-f format] profile
	*/

	bitcoindCmd = &cobra.Command{
		Use:   "bitcoind",
		Short: "A bitcoind simulator",
		Long:  `A bitcoind simulator that is queries for block templates as part of testing`,
		Run:   bitcoind,
	}
)

func bitcoind(cmd *cobra.Command, args []string) {

	fmt.Println("bitcoind called")
	auth_name, auth_pass, err := validation.ValidateAuth(&auth_str)

	if err != nil {
		log.Println("Authentication string format invalid.")
		log.Fatal(err)
	} else {
		log.Println("Authentication string format is valid.")

		// Load the block templates.
		// Trying to be descriptive about what block_templates is.
        log.Println("Loading block templates")
		var block_templates, err = loader.LoadFromBlockTemplatesDir()

        if err != nil {
            log.Fatal(err)
        }

		bss := server.BitcoindSimulationServer{}
		bss.StartServer(&port_num, &auth_name, &auth_pass, block_templates)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Default string is empty.
	rootCmd.AddCommand(bitcoindCmd)
	bitcoindCmd.Flags().StringVarP(&auth_str, "auth", "a", "", `Authentication in the form: (name:pass)`)
	// Default port num is 4000
	bitcoindCmd.Flags().Uint16VarP(&port_num, "port", "p", 4000, "Port number")
	bitcoindCmd.MarkFlagRequired("auth")
}
