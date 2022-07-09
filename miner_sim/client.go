package miner_sim

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"reflect"
	"strings"
)

type MinerClient struct {
	MinerData *Miner
	StratumID int64
}

// Increment MinerClient's stratumID
// then we return the current value of it after the increment.
func (mc *MinerClient) incrementID() int64 {
	mc.StratumID += 1
	return mc.StratumID
}

// Establish a TCP connection to the PoolServer
// wait for replies from the server and handle each of them accordingly
func (mc *MinerClient) ConnectToPool(ip string) {

	d := net.Dialer{}
	c, err := d.Dial("tcp", ip)

	if err != nil {
		log.Println(err)
		return
	}

	// Send mining.configure()
	str_dat := string(*mc.MinerData.GetMiningConfigure(mc.incrementID()))
	fmt.Fprintf(c, str_dat+"\n")

	r := bufio.NewReader(c)
	for {
		message, _ := r.ReadString('\n')
		fmt.Println(message)

		//First reply from Client to Server will be a subscribe request.
		reply, _ := mc.verifyResponse(message)

		// After verifying the response, if there is a reply, it means it's the next step
		// of the protocol "handshake". Often we will get JSONs with "id": null, that means
		// to NOT reply back.
		if reply != nil {
			str_dat = string(*reply)
			fmt.Fprintf(c, str_dat+"\n")
		}
	}
}

// A helper method for the ConnectToPool loop that does verification
// as well as tailoring the appropriate replies.
func (mc *MinerClient) verifyResponse(response string) (*[]byte, error) {
	unmarshed := mc.unmarshall(response)

	// The string can sometimes contain hidden delimiters/line-breaks.
	// This is to clean it and prevent bugs.
	cleaned_response := strings.Trim(response, " \r\n")

	// Check whether it has an indexable "result" or a "method" field.
	if _, ok := unmarshed["result"]; ok {

		// Grabbing the reflect.Type for map of interfaces, where the keys are strings.
		map_type := reflect.TypeOf(make(map[string]interface{}))

		// Grabbing the reflect.Type for array of interfaces.
		array_type := reflect.TypeOf(make([]interface{}, 0))

		// Grabbing the reflect.Type of bool
		bool_type := reflect.TypeOf(true)

		// TO-DO: Make a custom error for these cases.
		switch reflect.TypeOf(unmarshed["result"]) {
		case map_type:
			// If it's a map ( the first is ),
			// we verify it and send out the "mining.subscribe" method.
			expected := string(*mc.MinerData.GetExpectedConfigureResultResponse())
			if cleaned_response != expected {
				return nil, fmt.Errorf("Received incorrect configure result response.")
			}

			return mc.MinerData.GetMiningSubscribe(mc.incrementID()), nil
		case array_type:
			// If it's an array type, verify it (The mining subscribe result).
			expected := string(*mc.MinerData.GetExpectedSubscribeResultResponse())
			if cleaned_response != expected {
				return nil, fmt.Errorf("Received incorrect subscribe result response.")
			}
		case bool_type:
			log.Println("Result -> bool")
		}

	} else if val, ok := unmarshed["method"]; ok {

		// TO-DO: Make a custom error for these cases.
		switch val {
		case "mining.set_version_mask":
			expected := string(*mc.MinerData.GetExpectedSubscribeMethodResponse())
			if cleaned_response != expected {
				return nil, fmt.Errorf("Received incorrect configure method response.")
			}
		case "mining.set_difficulty":
			expected := string(*mc.MinerData.GetExpectedSubscribeMethodResponse())
			if cleaned_response != expected {
				return nil, fmt.Errorf("Received incorrect subscribe method response.")
			}
		case "mining.notify":
			// Currently no data available for verification
			log.Println("[mining.notify] Simulating work")
		}
	}

	return nil, nil
}

// A helper function for verifyResponse.
// This turns raw string JSONs into indexable maps.
func (mc *MinerClient) unmarshall(jsonString string) (res map[string]interface{}) {
	// Create the container for the string that will be unmarshalled
	var reading map[string]interface{}

	// The act of unmarshalling.
	err := json.Unmarshal([]byte(jsonString), &reading)
	if err != nil {
		fmt.Println(err)
	}
	return reading
}

type MinerClientFactory struct{}

func (f *MinerClientFactory) BuildMinerClient(miner *Miner) *MinerClient {
	return &MinerClient{MinerData: miner, StratumID: 0}
}

func NewMinerClientFactory() *MinerClientFactory {
	return &MinerClientFactory{}
}
