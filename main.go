// read JSON from stdin, pretty-printed to stdout
package main

import (
	"encoding/json"
	"os"
)

func main() {
	dec := json.NewDecoder(os.Stdin)
	var doc map[string]interface{}
	if err := dec.Decode(&doc); err != nil {
		panic(err)
	}
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "    ")
	if err := enc.Encode(doc); err != nil {
		panic(err)
	}
}
