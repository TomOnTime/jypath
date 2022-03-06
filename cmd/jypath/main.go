package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/TomOnTime/jypath/datapath"
)

/*


jypath file.json >data.path
jypath -json file.json >data.path
jypath -yaml file.json >data.path
jypath -r -yaml file.path >data.yaml
jypath -r -json file.path >data.json
jypath -jq -json file.json >data.path

*/

var flagJSON bool
var flagYAML bool
var flagReverse bool

func init() {
	flag.BoolVar(&flagJSON, "json", false, "JSON format")
	flag.BoolVar(&flagJSON, "j", false, "alias for --json")
	flag.BoolVar(&flagYAML, "yaml", false, "YAML format (default)")
	flag.BoolVar(&flagYAML, "y", false, "alias for --yaml")
	flag.BoolVar(&flagReverse, "reverse", false, "reverse (input is PATH)")
	flag.BoolVar(&flagReverse, "r", false, "alias for --revers")

}

func main() {
	flag.Parse()
	w := flag.CommandLine.Output()
	if flagJSON && flagYAML {
		fmt.Fprintf(w, "Error: --json and --yaml may not both be set at the same time.")
		os.Exit(1)
	}
	fmt.Printf("Flags:\n\tJSON: %v\n\tYAML: %v\n\tReverse: %v\n", flagJSON, flagYAML, flagReverse)
	if flag.NArg() > 1 {
		fmt.Fprintf(w, "Error: Zero or one filenames expected. Found %v.", flag.NArg())
		os.Exit(1)
	}

	var idata []byte
	var data interface{}
	var err error

	if flag.NArg() == 1 {
		idata, err = ioutil.ReadFile(flag.Arg(0))
	} else {
		idata, err = ioutil.ReadAll(os.Stdin)
	}
	if err != nil {
		fmt.Fprintf(w, "Could not read input: %s", err)
	}

	if flagReverse {
		data = datapath.FromPaths(idata)
		var odata []byte
		if flagJSON {
			odata, err = datapath.ToJSON(data)
		} else {
			odata, err = datapath.ToYAML(data)
		}
		if err != nil {
			fmt.Fprintf(w, "Could not generate output: %s", err)
		}
		fmt.Fprintln(w, string(odata))
	} else {
		if flagJSON {
			data, err = datapath.FromJSON(idata)
		} else {
			data, err = datapath.FromYAML(idata)
		}
		if err != nil {
			fmt.Fprintf(w, "Could not parse input: %s", err)
		}
		fmt.Fprintln(w, datapath.ToPathsString(data))
	}
}
