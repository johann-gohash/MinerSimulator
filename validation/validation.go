package validation

/*
    A collection of methods for validating data for the simulation's bitcoind mode.
*/

import (
    "strings"
    "encoding/json"
    "log"
    "errors"
    "github.com/go-playground/validator/v10"
)

var (
    ErrInvalidNamePassString = errors.New(`Invalid authentication details. 
        Please follow the format "username:pass"`)
)

// When the app is launched, one of the flags takes a string
// in the form "user:name:". ValidateAuth verifies it by trying to split
// the string into this struct.
type User struct {
    Name string `json:"name" validate:"required,excludes= "`
    Pass string `json:"pass" validate:"required"`
}

// A client makes a request to the bitcoind simulator. With the request comes a piece 
// of JSON RPC data that can be broken down into this struct for verification.
type ClientRPC struct {
    Method string `json:"name"`
    Params []ClientRPCRules `json:"params"`
    id int `json: "id"`
}

// A parameter inside the RPC is an array of strings for "rules"
type ClientRPCRules struct {
    Rules []string `json:"rules"`
}

//  We expect to receive a string in the form name:pass
//  The first step is to split them by the ":" symbol,
//  put it in a User struct, and let the validator module take
//  care of the rest.
func ValidateAuth(auth_str *string) (name string, pass string, err error) {

    separated := strings.Split(*auth_str, ":")

    if len(separated) != 2 {
        err := ErrInvalidNamePassString
        return name, pass, err
    }

    v := validator.New()
    user_struct := User{
        Name: separated[0],
        Pass: separated[1],
    }
    err = v.Struct(user_struct)

    if err != nil {
        for _, e := range err.(validator.ValidationErrors) {
            log.Println(e)
            err := ErrInvalidNamePassString
            return name, pass, err
        }
    }

    return separated[0], separated[1], err
}

//  For now we have no actual use for the body,
//  we just check if it's valid.
//  We expect a nil error from this if the JSON is good.
func ValidateRPC(rpc_str * []byte) error {
    var err error
    var rpc_body ClientRPC

    err = json.Unmarshal([]byte(*rpc_str), &rpc_body)
    return err
}
