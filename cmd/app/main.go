package main

import (
	"fmt"
	"os"

	"github.com/Gartenschlaeger/strcli/pkg/colors"
	"github.com/Gartenschlaeger/strcli/pkg/commands"
)

type CommandHandler func([]string) error

var cmdsMap map[string]CommandHandler

func printUsage() {
	fmt.Printf("%vstr <command> [flags]%v\n", colors.Cyan, colors.Reset)
	os.Exit(1)
}

func printUnknownCommand(commandName string) {
	fmt.Printf("%vThe command '%s' is not a known command.%v\n", colors.Red, commandName, colors.Reset)
	os.Exit(1)
}

func init() {
	cmdsMap = make(map[string]CommandHandler)
	cmdsMap["field"] = commands.FieldCommand
}

func main() {
	if len(os.Args) < 2 {
		printUsage()
	}

	cmdName := os.Args[1]
	if h, f := cmdsMap[cmdName]; f {
		err := h(os.Args[2:])
		if err != nil {
			fmt.Printf("%v%v%v\n", colors.Red, err, colors.Reset)
		}
	} else {
		printUnknownCommand(cmdName)
	}
}
