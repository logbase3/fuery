/*
	Fuery, is a small and simple tool for querying files using SQL.
	Copyright (C) 2013 log₃() <contact@logbase3.com>

	This program is free software: you can redistribute it and/or modify
	it under the terms of the GNU General Public License as published by
	the Free Software Foundation, either version 3 of the License, or
	(at your option) any later version.

	This program is distributed in the hope that it will be useful,
	but WITHOUT ANY WARRANTY; without even the implied warranty of
	MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
	GNU General Public License for more details.

	You should have received a copy of the GNU General Public License
	along with this program. If not, see <http://www.gnu.org/licenses/>.

	For more information visit https://github.com/logbase3/fuery/fuery
	or send an e-mail to contact@logbase3.com
*/

// Command fuery is a small and simple tool for querying files using SQL.
// You can embed blocks of code in your godoc, such as this:
//  fmt.Println("Hello")
// To do that, simply add an extra indent to your comment's text.
package main // import "logbase3.com/fuery/fuery"

import (
	"flag"
	"fmt"
	"github.com/logbase3/fuery/fuery/constants"
	"github.com/peterh/liner"
	"os"
	"os/user"
	"path/filepath"
	"strings"
)

// Global Variables
var (
	prompt      string = "> "
	home        string
	historyFile string
	verbose     bool
	filesList   []string
)

// Initialize global variables
func init() {
	// Get user's history file
	usr, err := user.Current()
	if err != nil {
		panic(err)
	}
	home = usr.HomeDir
	historyFile = home + filepath.FromSlash(constants.HistoryFileName)
}

// Manage flags
func init() {
	const (
		verbUsage = "When set to true, prints more information."
	)

	// Debug options
	flag.BoolVar(&verbose, "verbose", false, verbUsage)
	flag.BoolVar(&verbose, "v", false, verbUsage+" (shorthand)")
}

func main() {
	flag.Parse()
	filesList = flag.Args()

	fmt.Printf(constants.Header)

	line := liner.NewLiner()
	defer line.Close()

	// Set liner options
	line.SetCompleter(commandCompleter)
	//line.SetWordCompleter(lineCompleter)
	line.SetCtrlCAborts(false)
	line.SetTabCompletionStyle(liner.TabPrints)

	// Load history file
	if f, err := os.Open(historyFile); err == nil {
		line.ReadHistory(f)
		f.Close()
	}

	// prompt loop
	var input string
	var err error
	for {
		if input, err = line.Prompt(prompt); err != nil {
			fmt.Println("\nError reading line: ", err)
			break
		}
		input = strings.Trim(input, "\t\r\n ")

		// If empty, ignore line
		if input == "" {
			continue
		}

		if input[0] == '\\' { // If input is a system command
			switch strings.Split(input, " ")[0] {
			case "\\copyright":
				fmt.Printf("%s\n\n", constants.License)
			case "\\?":
				fmt.Printf("%s\n\n", constants.SystemHelp)
			case "\\h":
				fmt.Printf("%s\n\n", "SQL STATEMENTS HELP")
			case "\\q":
				// Bug(Roberto Lapuente): This does not saves history on it's way out
				os.Exit(0)
			case "\\watch":
				fmt.Printf("%s\n\n", "WATCHING")
			case "\\copy":
			case "\\echo":
			case "\\i":
			case "\\o":
			case "\\timing":
			default:
				fmt.Println("Got:", input)
			}
		} else { // If input is a statement

		}

		line.AppendHistory(input)
	}

	// Write history to file
	if f, err := os.Create(historyFile); err != nil {
		fmt.Printf("Error writing history file %s: %s", historyFile, err)
	} else {
		line.WriteHistory(f)
		f.Close()
	}
}

func commandCompleter(line string) (c []string) {
	var systemCommands = []string{"\\copyright", "\\?", "\\h"}
	var sqlStatements = []string{"SELECT", "ATUN", "CREMAdeELOTE"}

	for _, n := range systemCommands {
		if strings.HasPrefix(n, strings.ToLower(line)) {
			c = append(c, n+" ")
		}
	}

	for _, n := range sqlStatements {
		if strings.HasPrefix(n, strings.ToUpper(line)) {
			c = append(c, n+" ")
		}
	}
	return
}

func lineCompleter(line string, pos int) (head string, c []string, tail string) {
	var systemCommands = []string{"\\copyright", "\\?", "\\h"}
	var sqlStatements = []string{"SELECT"}

	for _, n := range systemCommands {
		if strings.HasPrefix(n, line[pos:]) {
			c = append(c, n)
		}
	}

	for _, n := range sqlStatements {
		if strings.HasPrefix(n, strings.ToUpper(line[pos:])) {
			c = append(c, n)
		}
	}
	head = line[:pos]
	tail = line[pos:]
	return
}
