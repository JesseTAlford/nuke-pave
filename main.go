package main

import (
	"bufio"
	"fmt"
	"github.com/cloudfoundry/cli/plugin"
	"os"
	"strings"
)

type NukePlugin struct{}

func (c *NukePlugin) Run(cliConnection plugin.CliConnection, args []string) {
	if args[0] == "nuke-pave" {
		fmt.Println("\ncircumstances are cyclical\n")
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("are you certain you wish to evoke creative destruction?\n\n> ")
		input, _ := reader.ReadString('\n')
		if input == "y\n" || input ==  "yes\n" || input == "Y\n" || input == "Yes\n" {
			fmt.Println("\ndestroying and rebuilding the space you inhabit...")
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
