/*
    Copyright © 2022 Johann Suarez <johann@gohash.ca>
*/
package cmd

import (
	"os"
    "log"
    "fmt"
	"github.com/spf13/cobra"
    "errors"
)


var rootCmd = &cobra.Command{
	Use:   "gohash_miner_simulator",
	Short: "Testing software that simulates both bitcoind and miner clients.",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// Make this return an error
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {

    env_vars := []string{"MINER_TEMPLATE_DIR", "BLOCK_TEMPLATE_DIR"}

    for _, env_var := range(env_vars) {
        err := check_env_var(env_var)
        if err != nil {
            // Propagating the error
            log.Println(err)
            os.Exit(1)
        }
    }
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

// Check the validity of environment variables.
// They should exist and they recognizable as directories.
func check_env_var(key string) error {

    // Checks if the environment variable has been set ( and not just if it's empty. )
    val, ok := os.LookupEnv(key)

    if !ok {
        message := fmt.Sprint("ERROR: ", key, " environment variable is not set.")
        return errors.New(message)

    } else {
        // If the environment variable has been set, we check if it's a valid directory
        // using os.Open
        _, err := os.Open(val)

        if err != nil {
            message := fmt.Sprint("ERROR: ", val, " is not a valid directory. From env var: ", key)
            return errors.New(message)
        } 
    }
    return nil
}

func init() {

    // Check for the environment variables here.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
