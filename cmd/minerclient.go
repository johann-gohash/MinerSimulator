/*
    Copyright © 2022 Johann Suarez <johann@gohash.tech>
*/

package cmd

import (
	"github.com/spf13/cobra"
	"log"
    "fmt"
	"miner_sim"
    //"time"
)

// minerClientCmd represents the minerClient command
var (
	server_str string
	port       uint16
	load       uint32
	miner_type string
	all        bool

	minerClientCmd = &cobra.Command{
		Use:   "miner",
		Short: "A TCP client-driven ASIC miner simulator.",
		Run:   root_cmd,
	}
)

// This is where the simulation is set up and ran.
// If a miner type is specified, the we instantiate a factory
// that produces only that type.
// The mine contains those miner instances, and it can then
// make the final call to make those miner instances connect through TCP.
func root_cmd(cmd *cobra.Command, args []string) {

	// Debugging statements, should be fixed/removed before production.
	log.Println("Miner simulator starting...")

    pool_server_address := server_str + ":" + fmt.Sprint(port)

	// Instantiate a Miner Factory.
	miner_factory := miner_sim.NewMinerFactory("s17")

    // Instantiate a MinerClient Factory.
    mcf := miner_sim.NewMinerClientFactory()

    // Instantiating a Mine
    mine := miner_sim.Mine{}

    // Filling the Mine with Miners based on given load.
    for i := uint32(0); i < load; i++ {
        // Adding Clients directly from the builder(s)
        mine.AddMiner(mcf.BuildMinerClient(miner_factory.BuildMiner()))
    }

    // Having the mine connect its clients.
    mine.ConnectMiners(pool_server_address)
    // Keep the programming running until CTRL + C
    for {  } 
}


func init() {
	rootCmd.AddCommand(minerClientCmd)
	// Here you will define your flags and configuration settings.
	minerClientCmd.Flags().StringVarP(&server_str, "server", "s", "127.0.0.1",
		`Stratum URL or public IP address, defaults to localhost.`)

	minerClientCmd.Flags().Uint16VarP(&port, "port", "p", 60000,
		`Corresponding port, defaults to 60000`)
	minerClientCmd.Flags().Uint32VarP(&load, "load", "l", 10, "The number of miners you want to connect, defaults to 10.")
	minerClientCmd.Flags().StringVarP(&miner_type, "miner-type", "m", "all", `Name of the miner to target (s9, s17, s19, m30)`)
	minerClientCmd.Flags().BoolVar(&all, "all", true, `Will load all miner types`) 
}
