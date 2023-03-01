package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/Gartenschlaeger/strcli/pkg/colors"
	"github.com/Gartenschlaeger/strcli/pkg/commands"
)

type CommandHandler func(string, []string) (string, error)

var cmdsMap map[string]CommandHandler

func printUsage() {
	fmt.Printf("%vstr <command> [flags]%v\n", colors.Cyan, colors.Reset)
	os.Exit(1)
}

func printUnknownCommand(commandName string) {
	fmt.Printf("%vThe command '%s' is not a known command.%v\n", colors.Red, commandName, colors.Reset)
	os.Exit(1)
}

func printNoInput() {
	fmt.Printf("%vNo input%v\n", colors.Red, colors.Red)
	os.Exit(1)
}

func init() {
	cmdsMap = make(map[string]CommandHandler)
	cmdsMap["field"] = commands.FieldCommand
}

func getStandardInputString() string {
	fi, _ := os.Stdin.Stat()
	if (fi.Mode() & os.ModeCharDevice) == 0 {
		bytes, _ := ioutil.ReadAll(os.Stdin)
		str := string(bytes)
		str = strings.Trim(str, "\n ")

		return str
	}

	return ""
}

func main() {
	if len(os.Args) < 2 {
		printUsage()
	}

	cmdName := os.Args[1]
	if h, f := cmdsMap[cmdName]; f {
		input := getStandardInputString()
		if len(input) > 0 {
			result, err := h(input, os.Args[2:])
			if err != nil {
				fmt.Printf("%v%v%v\n", colors.Red, err, colors.Reset)
			}

			fmt.Print(result)
		} else {
			printNoInput()
		}
	} else {
		printUnknownCommand(cmdName)
	}
}
