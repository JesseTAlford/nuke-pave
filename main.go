package main

import (
	"fmt"
	"github.com/cloudfoundry/cli/plugin"
	"strings"
)

type NukePlugin struct{}

func (c *NukePlugin) Run(cliConnection plugin.CliConnection, args []string) {
	if args[0] == "nuke-pave" {
		fmt.Println("\ncircumstances are cyclical\n\ndestroying and rebuilding the space you inhabit...")
		target, err := cliConnection.CliCommandWithoutTerminalOutput("target")
		if err != nil {
			fmt.Println("PLUGIN ERROR: Error from CliCommand: ", err)
		}
		spaceName := strings.TrimSpace(strings.TrimPrefix(target[4], "Space:"))
		cliConnection.CliCommandWithoutTerminalOutput("delete-space", "-f", spaceName)
		cliConnection.CliCommandWithoutTerminalOutput("create-space", spaceName)
		
		cliConnection.CliCommand("target", "-s", spaceName)
		fmt.Println("\nwhat was once is now again")
	}
}

func (c *NukePlugin) GetMetadata() plugin.PluginMetadata {
	return plugin.PluginMetadata{
		Name: "Nuke-and-Pave",
		Commands: []plugin.Command{
			plugin.Command{
				Name:     "nuke-pave",
				HelpText: "Deletes, recreates, and retargets your space.",
			},
		},
	}
}

func main() {
	plugin.Start(new(NukePlugin))
}
