package miner_sim

import (
	"encoding/json"
	"loader"
	"log"
	"math/rand"
	"time"
)

type Miner struct {
	Name            string
	data            *Mining
	vardiff_timeout int64
}

// Returns the Vardiff timeout
func (m Miner) GetVardiffTimeout() int64 {
	return m.vardiff_timeout
}

//-------------=| Mining Request Methods |=-------------//

func (m Miner) GetMiningConfigure(id int64) *[]byte {
	// Convert into string of bytes with MarshalIndent
	m.data.Configure.Request.ID = id
	config_bytes, _ := json.Marshal(m.data.Configure.Request)
	return &config_bytes
}

func (m Miner) GetExpectedConfigureResultResponse() *[]byte {
	result_response_bytes, _ := json.Marshal(m.data.Configure.ExpectedResponse.ResultResponse)
	return &result_response_bytes

}


func (m Miner) GetExpectedConfigureMethodResponse() *[]byte {
	method_response_bytes, _ := json.Marshal(m.data.Configure.ExpectedResponse.MethodResponse)
	return &method_response_bytes
}

//-------------=| Subscribe Request Methods |=---------//

func (m Miner) GetMiningSubscribe(id int64) *[]byte {
	// Convert into string of bytes with MarshalIndent
	m.data.Subscribe.Request.ID = id
	config_bytes, _ := json.Marshal(m.data.Subscribe.Request)
	return &config_bytes
}

func (m Miner) GetExpectedSubscribeResultResponse() *[]byte {
	result_response_bytes, _ := json.Marshal(m.data.Subscribe.ExpectedResponse.ResultResponse)
	return &result_response_bytes
}

func (m Miner) GetExpectedSubscribeMethodResponse() *[]byte {
	//method_response_bytes, _ := json.MarshalIndent(m.data.Subscribe.ExpectedResponse.MethodResponse, "", "")
	method_response_bytes, _ := json.Marshal(m.data.Subscribe.ExpectedResponse.MethodResponse)
	return &method_response_bytes
}

//-------------=| Authorize Request Methods |=---------//
func (m Miner) GetMiningAuthorize(id int) *[]byte {
	// Returns the mining authorize expected server response
	authorize_bytes, _ := json.Marshal(m.data.Authorize.Request)
	return &authorize_bytes
}

func (m Miner) GetExpectedAuthorizeResultResponse() *[]byte {
	result_response_bytes, _ := json.Marshal(m.data.Authorize.ExpectedResponse.ResultResponse)
	return &result_response_bytes
}

func (m Miner) GetExpectedAuthorizeMethodResponse() *[]byte {
	method_response_bytes, _ := json.Marshal(m.data.Authorize.ExpectedResponse.MethodResponse)
	return &method_response_bytes
}

// On hold..
func (m Miner) GetMiningSubmit(id int) *[]byte {
	str := []byte("Placeholder")
	return &str
}

// The Factory is only concerned with properly creating
// the desired miner instances. Miners have to be instantiated with
// the proper data from the data dir
type MinerFactory struct {
    // This is an array of array of bytes.
	MinerDataSet *[][]byte
}

// This will be called every time a Miner is to be built.
// If MinerFactory is specified to build all miner types, it will
// choose their types randomly, otherwise it will stick to giving just
// one typ.
func (f *MinerFactory) giveMinerData() (string, *Mining, int64, error) {

	// Get the length of MinerDataSet, assign it to x.
	dataset_len := len(*f.MinerDataSet)

	// If length is one, just access the first element of MinerDataSet and return it.
	if dataset_len == 1 {
		unmarshed_mdata, err := UnmarshalMinerData((*f.MinerDataSet)[0])

		// Remarshall ( No immediate way to convert to array of bytes otherwise )
		if err != nil {
			return "", &Mining{}, 0, err
		}

		return unmarshed_mdata.MinerType,
			&unmarshed_mdata.Mining,
			unmarshed_mdata.VardiffTimeout,
			nil
	} else {

		// Otherwise, choose a random int between 0 and x.
		// Use that as an index to MinerDataSet, return what slice we got.
		// Get a seed. It has to be Unix Nano because Unix() is too slow.
		rand.Seed(time.Now().UnixNano())
		index := rand.Intn(dataset_len)

		unmarshed_mdata, err := UnmarshalMinerData((*f.MinerDataSet)[index])

		if err != nil {
			return "", &Mining{}, 0, err
		}

		return unmarshed_mdata.MinerType,
			&unmarshed_mdata.Mining,
			unmarshed_mdata.VardiffTimeout, nil
	}
}

// This method is what creates Miner instances.
func (f *MinerFactory) BuildMiner() *Miner {
	name, data, timeout, _ := f.giveMinerData()
	return &Miner{name, data, timeout}
}

// Function that instantiates a new factory according to our desired Miner data and vardiff
// First argument should be an enum for the only valid types of miners
// Use the loader to load in the miner type(s)
func NewMinerFactory(miner_type string) *MinerFactory {
	var loading_err error
	var miner_templates *[][]byte

	if len(miner_type) > 0 {
		// We specified a miner.
		miner_templates, loading_err = loader.LoadFromMinerTemplatesDir(miner_type)

	} else {
		// Else "all" was true
		miner_templates, loading_err = loader.LoadFromMinerTemplatesDir("all")
	}

	if loading_err != nil {
		log.Fatal(loading_err)
	}

	return &MinerFactory{miner_templates}
}
