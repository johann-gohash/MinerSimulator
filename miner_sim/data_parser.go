// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    minerData, err := UnmarshalMinerData(bytes)
//    bytes, err = minerData.Marshal()

package miner_sim

import "bytes"
import "errors"
import "encoding/json"


func UnmarshalMinerData(data []byte) (MinerData, error) {
	var r MinerData
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *MinerData) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type MinerData struct {
	MinerType      string `json:"miner_type"`     
	VardiffTimeout int64  `json:"vardiff_timeout"`
	Mining         Mining `json:"mining"`         
}

type Mining struct {
	Configure Configure `json:"configure"`
	Subscribe Subscribe `json:"subscribe"`
	Authorize Authorize `json:"authorize"`
	Submit    Submit    `json:"submit"`   
}

type Authorize struct {
	Request          AuthorizeRequest          `json:"request"`          
	ExpectedResponse AuthorizeExpectedResponse `json:"expected_response"`
}

type AuthorizeExpectedResponse struct {
	ResultResponse PurpleResultResponse `json:"result_response"`
	MethodResponse ResultResponseClass  `json:"method_response"`
}

type ResultResponseClass struct {
}

type PurpleResultResponse struct {
	ID     int64       `json:"id"`    
	Result bool        `json:"result"`
	Error  interface{} `json:"error"` 
}

type AuthorizeRequest struct {
	ID     int64   `json:"id"`    
	Method string   `json:"method"`
	Params []string `json:"params"`
}

type Configure struct {
	Request          ConfigureRequest          `json:"request"`          
	ExpectedResponse ConfigureExpectedResponse `json:"expected_response"`
}

type ConfigureExpectedResponse struct {
	ResultResponse FluffyResultResponse `json:"result_response"`
	MethodResponse AuthorizeRequest     `json:"method_response"`
}

type FluffyResultResponse struct {
	Error  interface{} `json:"error"` 
	ID     int64       `json:"id"`    
	Result ResultClass `json:"result"`
}

type ResultClass struct {
	VersionRolling     bool   `json:"version-rolling"`     
	VersionRollingMask string `json:"version-rolling.mask"`
}

type ConfigureRequest struct {
	ID     int64          `json:"id"`    
	Method string         `json:"method"`
	Params []ParamElement `json:"params"`
}

type ParamClass struct {
	VersionRollingMask        string `json:"version-rolling.mask"`         
	VersionRollingMinBitCount int64  `json:"version-rolling.min-bit-count"`
}

type Submit struct {
	Request          ResultResponseClass    `json:"request"`          
	ExpectedResponse SubmitExpectedResponse `json:"expected_response"`
}

type SubmitExpectedResponse struct {
	ResultResponse ResultResponseClass `json:"result_response"`
	MethodResponse ResultResponseClass `json:"method_response"`
}

type Subscribe struct {
	Request          AuthorizeRequest          `json:"request"`          
	ExpectedResponse SubscribeExpectedResponse `json:"expected_response"`
}

type SubscribeExpectedResponse struct {
	MethodResponse MethodResponse          `json:"method_response"`
	ResultResponse TentacledResultResponse `json:"result_response"`
}

type MethodResponse struct {
	ID     interface{} `json:"id"`    
	Method string      `json:"method"`
	Params []int64     `json:"params"`
}

type TentacledResultResponse struct {
	Error  interface{}     `json:"error"` 
	ID     int64           `json:"id"`    
	Result []ResultElement `json:"result"`
}

type ParamElement struct {
	ParamClass  *ParamClass
	StringArray []string
}

func (x *ParamElement) UnmarshalJSON(data []byte) error {
	x.StringArray = nil
	x.ParamClass = nil
	var c ParamClass
	object, err := unmarshalUnion(data, nil, nil, nil, nil, true, &x.StringArray, true, &c, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
		x.ParamClass = &c
	}
	return nil
}

func (x *ParamElement) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, nil, x.StringArray != nil, x.StringArray, x.ParamClass != nil, x.ParamClass, false, nil, false, nil, false)
}

type ResultElement struct {
	Integer          *int64
	String           *string
	StringArrayArray [][]string
}

func (x *ResultElement) UnmarshalJSON(data []byte) error {
	x.StringArrayArray = nil
	object, err := unmarshalUnion(data, &x.Integer, nil, nil, &x.String, true, &x.StringArrayArray, false, nil, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
	}
	return nil
}

func (x *ResultElement) MarshalJSON() ([]byte, error) {
	return marshalUnion(x.Integer, nil, nil, x.String, x.StringArrayArray != nil, x.StringArrayArray, false, nil, false, nil, false, nil, false)
}

func unmarshalUnion(data []byte, pi **int64, pf **float64, pb **bool, ps **string, haveArray bool, pa interface{}, haveObject bool, pc interface{}, haveMap bool, pm interface{}, haveEnum bool, pe interface{}, nullable bool) (bool, error) {
	if pi != nil {
		*pi = nil
	}
	if pf != nil {
		*pf = nil
	}
	if pb != nil {
		*pb = nil
	}
	if ps != nil {
		*ps = nil
	}

	dec := json.NewDecoder(bytes.NewReader(data))
	dec.UseNumber()
	tok, err := dec.Token()
	if err != nil {
		return false, err
	}

	switch v := tok.(type) {
	case json.Number:
		if pi != nil {
			i, err := v.Int64()
			if err == nil {
				*pi = &i
				return false, nil
			}
		}
		if pf != nil {
			f, err := v.Float64()
			if err == nil {
				*pf = &f
				return false, nil
			}
			return false, errors.New("Unparsable number")
		}
		return false, errors.New("Union does not contain number")
	case float64:
		return false, errors.New("Decoder should not return float64")
	case bool:
		if pb != nil {
			*pb = &v
			return false, nil
		}
		return false, errors.New("Union does not contain bool")
	case string:
		if haveEnum {
			return false, json.Unmarshal(data, pe)
		}
		if ps != nil {
			*ps = &v
			return false, nil
		}
		return false, errors.New("Union does not contain string")
	case nil:
		if nullable {
			return false, nil
		}
		return false, errors.New("Union does not contain null")
	case json.Delim:
		if v == '{' {
			if haveObject {
				return true, json.Unmarshal(data, pc)
			}
			if haveMap {
				return false, json.Unmarshal(data, pm)
			}
			return false, errors.New("Union does not contain object")
		}
		if v == '[' {
			if haveArray {
				return false, json.Unmarshal(data, pa)
			}
			return false, errors.New("Union does not contain array")
		}
		return false, errors.New("Cannot handle delimiter")
	}
	return false, errors.New("Cannot unmarshal union")

}

func marshalUnion(pi *int64, pf *float64, pb *bool, ps *string, haveArray bool, pa interface{}, haveObject bool, pc interface{}, haveMap bool, pm interface{}, haveEnum bool, pe interface{}, nullable bool) ([]byte, error) {
	if pi != nil {
		return json.Marshal(*pi)
	}
	if pf != nil {
		return json.Marshal(*pf)
	}
	if pb != nil {
		return json.Marshal(*pb)
	}
	if ps != nil {
		return json.Marshal(*ps)
	}
	if haveArray {
		return json.Marshal(pa)
	}
	if haveObject {
		return json.Marshal(pc)
	}
	if haveMap {
		return json.Marshal(pm)
	}
	if haveEnum {
		return json.Marshal(pe)
	}
	if nullable {
		return json.Marshal(nil)
	}
	return nil, errors.New("Union must not be null")
}
