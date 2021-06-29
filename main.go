// Command json2yaml converts JSON documents to YAML.
//
// If command is called by yaml2json name, then it converts from YAML to JSON.
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

func main() {
	flag.Parse()
	log.SetFlags(0)
	if err := run(flag.Arg(0)); err != nil {
		log.Fatal(err)
	}
}

func run(name string) error {
	input, output := os.Stdin, os.Stdout
	if name != "" && name != "-" {
		f, err := os.Open(name)
		if err != nil {
			return err
		}
		defer f.Close()
		input = f
	}
	if reverseMode {
		return yaml2json(input, output)
	}
	return json2yaml(input, output)
}

func json2yaml(input io.Reader, output io.Writer) error {
	dec := json.NewDecoder(input)
	enc := yaml.NewEncoder(output)
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

func yaml2json(input io.Reader, output io.Writer) error {
	dec := yaml.NewDecoder(input)
	enc := json.NewEncoder(output)
	enc.SetEscapeHTML(false)
	enc.SetIndent("", "  ")
	for {
		var doc interface{}
		if err := dec.Decode(&doc); err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
		if err := enc.Encode(doc); err != nil {
			return err
		}
	}
}

var reverseMode = len(os.Args) != 0 && filepath.Base(os.Args[0]) == "yaml2json"

func init() {
	flag.Usage = func() {
		if reverseMode {
			fmt.Fprintln(flag.CommandLine.Output(), "Usage: yaml2json < input.yaml > output.json\n"+
				"   or: yaml2json input.yaml > output.json")
		} else {
			fmt.Fprintln(flag.CommandLine.Output(), "Usage: json2yaml < input.json > output.yaml\n"+
				"   or: json2yaml input.json > output.yaml")
		}
	}
}
