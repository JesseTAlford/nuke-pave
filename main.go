package main

import (
	"fmt"

	"github.com/cloudfoundry/cli/plugin"
)

type BasicPlugin struct{}

func (c *BasicPlugin) Run(cliConnection plugin.CliConnection, args []string) {
	if args[0] == "nuke-pave" {
		fmt.Println("circumstances are cyclical\n
		destroying and rebuilding the space you inhabit")
	}
}

func (c *BasicPlugin) GetMetadata() plugin.PluginMetadata {
	return plugin.PluginMetadata{
		Name: "Nuke and Pave",
		Commands: []plugin.Command{
			plugin.Command{
				Name: 	  "nuke-pave",
				HelpText: "Deletes, recreates, and retargets your space.",
			},
		},
	}
}

func main() {
	plugin.Start(new(BasicPlugin))
}