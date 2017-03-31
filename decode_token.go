package main

import (
	"code.cloudfoundry.org/cli/plugin"
	"encoding/base64"
	"fmt"
	"github.com/hokaccha/go-prettyjson"
	"strings"
)

type DecodeTokenPlugin struct{}

func decode(target string) (r []byte, e error) {
	if m := len(target) % 4; m != 0 {
		target += strings.Repeat("=", 4-m)
	}
	return base64.URLEncoding.DecodeString(target)
}

func prettyPrint(target string) {
	f := prettyjson.NewFormatter()
	p, _ := f.Format([]byte(target))
	fmt.Println(string(p))
}

func (c *DecodeTokenPlugin) Run(cliConnection plugin.CliConnection, args []string) {
	if args[0] == "decode-token" || args[0] == "dt" {

		at, _ := cliConnection.AccessToken()
		t := strings.Split(at, " ")[1]
		eh := strings.Split(t, ".")[0]
		ep := strings.Split(t, ".")[1]
		es := strings.Split(t, ".")[2]

		h, _ := decode(eh)
		p, _ := decode(ep)

		fmt.Println("HEADER =================================")
		prettyPrint(string(h))
		fmt.Println("PAYLOAD ================================")
		prettyPrint(string(p))
		fmt.Println("SIGNATURE ==============================")
		fmt.Println(es)

	}
}

func (c *DecodeTokenPlugin) GetMetadata() plugin.PluginMetadata {
	return plugin.PluginMetadata{
		Name: "DecodeTokenPlugin",
		Version: plugin.VersionType{
			Major: 1,
			Minor: 0,
			Build: 0,
		},
		MinCliVersion: plugin.VersionType{
			Major: 6,
			Minor: 7,
			Build: 0,
		},
		Commands: []plugin.Command{
			{
				Name:     "decode-token",
				Alias:    "dt",
				HelpText: "Decode JWT AccessToken",

				UsageDetails: plugin.Usage{
					Usage: "decode-token,dt\n   cf decode-token",
				},
			},
		},
	}
}

func main() {
	plugin.Start(new(DecodeTokenPlugin))
}
