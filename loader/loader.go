package loader

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

// This function loads the environment variable first
// then makes a call to load
func LoadFromBlockTemplatesDir() (*[][]byte, error) {

	dir := os.Getenv("BLOCK_TEMPLATE_DIR")

	// If path doesn't end with a "/", we add it ( to rectify a possible small mistake )
	if !(strings.HasSuffix(dir, "/")) {
		dir = dir + "/"
	}

	return loadJSONs(dir, "")
}

// Function to be used by the Miner sim, this one takes an argument
func LoadFromMinerTemplatesDir(miner string) (*[][]byte, error) {

	dir := os.Getenv("MINER_TEMPLATE_DIR")

	if !(strings.HasSuffix(dir, "/")) {
		dir = dir + "/"
	}
	return loadJSONs(dir, miner)
}

// Upon receiving the data directory environment variable,
// this function checks the directory for JSONs, and loads them from disk.
// ioutil reads a file and turns it into a slice of bytes.
// We'll be getting several of those... a slice of slice of bytes.
// Finally, we return a pointer to that.
func loadJSONs(path string, file string) (*[][]byte, error) {

	var json_files [][]byte
	files, err := os.ReadDir(path)

	if err != nil {
		return nil, err
	}

	// If no specific file was provided, or we specify all.
	if file == "" || file == "all" {
		for _, f := range files {
			if strings.HasSuffix(f.Name(), ".json") {
				absolute_file_address := fmt.Sprint(path, f.Name())
				content, err := ioutil.ReadFile(absolute_file_address)

				if err != nil {
					return nil, err
				}
				json_files = append(json_files, content)
			}
		}

	} else {
		// If we're here, the user specified a miner type.
		// So we try to look for that in the MINER_TEMPLATE_DIR and load it.
		var found_jsons []string
		files, err := os.ReadDir(path)

		if err != nil {
			return nil, err
		}

		for _, f := range files {
			// If name ends in .json (lowercased for more robustness) add it to a list.
			if strings.HasSuffix(strings.ToLower(f.Name()), ".json") {
				// We trim the suffix so we won't have to expect
				// the user to explictly specify the ".json" part.
				found_jsons = append(found_jsons, strings.TrimSuffix(f.Name(), ".json"))
			}
		}

		// At this point if found_jsons is empty, we return an error.
		if len(found_jsons) == 0 {
			log.Println(len(found_jsons))
			err := errors.New("Specified JSON not found.")
			return nil, err
		}

		// Othwerwise we look for the specified miner's JSON and load it. If it's not there,
		// return an error.
		if contains(found_jsons, file) {
			absolute_file_address := fmt.Sprint(path, file, ".json")
			content, err := ioutil.ReadFile(absolute_file_address)

			if err != nil {
				return nil, err
			}
			json_files = append(json_files, content)

		} else {
			fmt.Println(found_jsons)
			err := errors.New(fmt.Sprint("Could not find ", file))
			return nil, err
		}
	}
	return &json_files, err
}

// A helper function to check for a string within a slice.
func contains(s []string, str string) bool {

	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}
