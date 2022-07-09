package server

import (
	"fmt"
	"io/ioutil"
	"log"
    "math/rand"
	"net/http"
	"strconv"
	"validation"
)

// This struct is needed for two purposes:
// To store the data for the actual credentials of the server, and
// to store the data for the received credentials sent from a client.
// If these two structs match, that means the client got the name and pass right
// and thus will be given a block template.
type NamePassPair struct {
	name string
	pass string
}

// This struct aims to represent a server class
// that has both methods and data members to keep the
// name and password provided through the command-line.
type BitcoindSimulationServer struct {
	authPair       NamePassPair // The real name-pass pair that clients must get right.
	blockTemplates *[][]byte    // Pre-loaded block templates from disk.

}

// This method is called after the HTTP server gets a request to the root path.
// It first checks the authentication data provided by the client, and if correct
// will send back a block header.
func (bss BitcoindSimulationServer) InspectRequest(w http.ResponseWriter,
	req *http.Request) {

	u, a, _ := req.BasicAuth()
	client_auth_pair := NamePassPair{u, a}

	// Buffer (buf) is an array of bytes.
	buf, err := ioutil.ReadAll(req.Body)

	if err != nil {
		log.Println("Invalid request received by client.")
		log.Println(err)
	}

	if client_auth_pair != bss.authPair {
		log.Println("Credentials don't match.")
        fmt.Fprintf(w, "{'error':'Invalid credentials.'}")
		return
	}

	err = validation.ValidateRPC(&buf)

	if err != nil {
		log.Println("Received RPC call from miner is invalid!")
        fmt.Fprintf(w, "{'error':'Invalid RPC'}")
	} else {
        // Randomly select a block template from the slice.
        random_index := rand.Intn(len(*bss.blockTemplates))
		fmt.Fprintf(w, string((*bss.blockTemplates)[random_index]))
	}
}

// We try to simulate a bitcoind server that must first receive
// the correct credentials (auth_name & auth_pass) before it can
// reply with a block header.
func (bss *BitcoindSimulationServer) StartServer(port_num *uint16,
	auth_name *string,
	auth_pass *string,
	block_templates *[][]byte) {

	var err error

	bss.authPair.name = *auth_name
	bss.authPair.pass = *auth_pass
	bss.blockTemplates = block_templates

	str_port_num := strconv.Itoa(int(*port_num))
	log.Println("Starting server with port " + str_port_num)
	http.HandleFunc("/", bss.InspectRequest)
	err = http.ListenAndServe(":"+str_port_num, nil)

	if err != nil {
		log.Fatal(err)
	}
}
