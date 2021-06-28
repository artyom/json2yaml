// Command json2yaml converts JSON documents to YAML
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

func main() {
	flag.Parse()
	log.SetFlags(0)
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	dec := json.NewDecoder(os.Stdin)
	enc := yaml.NewEncoder(os.Stdout)
	defer enc.Close()
	for {
		var doc interface{}
		if err := dec.Decode(&doc); err != nil {
			if err == io.EOF {
				return enc.Close()
			}
			return err
		}
		if err := enc.Encode(doc); err != nil {
			return err
		}
	}
}

func init() {
	flag.Usage = func() {
		fmt.Fprintln(flag.CommandLine.Output(), "Usage: json2yaml < input.json > output.yaml")
	}
}
