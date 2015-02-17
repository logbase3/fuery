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

	For more information visit https://github.com/logbase3/fuery
	or send an e-mail to contact@logbase3.com
*/

// Fuery, is a small and simple tool for querying files using SQL.
package main

import (
	"flag"
	"fmt"
	"github.com/peterh/liner"
	"os"
	"os/user"
	"path/filepath"
	"strings"
)

// Global Variables
var (
	PROMPT       string = "> "
	HOME         string
	HISTORY_FILE string
	VERBOSE      bool
	FILES_LIST   []string
)

// Initialize global variables
func init() {
	// Get user's history file
	usr, err := user.Current()
	if err != nil {
		panic(err)
	}
	HOME = usr.HomeDir
	HISTORY_FILE = HOME + filepath.FromSlash(HISTORY_FILE_NAME)
}

// Manage flags
func init() {
	const (
		verbUsage = "When set to true, prints more information."
	)

	// Debug options
	flag.BoolVar(&VERBOSE, "verbose", false, verbUsage)
	flag.BoolVar(&VERBOSE, "v", false, verbUsage+" (shorthand)")
}

func main() {
	flag.Parse()
	FILES_LIST = flag.Args()

	fmt.Printf("%s\n", HEADER)

	line := liner.NewLiner()
	defer line.Close()

	// Set liner options
	line.SetCompleter(commandCompleter)
	line.SetCtrlCAborts(false)

	// Load history file
	if f, err := os.Open(HISTORY_FILE); err == nil {
		line.ReadHistory(f)
		f.Close()
	}

	// Prompt loop
	for {
		if name, err := line.Prompt(PROMPT); err != nil {
			fmt.Println("\nError reading line: ", err)
			break
		} else {
			fmt.Println("Got:", name)
			line.AppendHistory(name)
		}
	}

	// Write history to file
	if f, err := os.Create(HISTORY_FILE); err != nil {
		fmt.Printf("Error writing history file %s: %s", HISTORY_FILE, err)
	} else {
		line.WriteHistory(f)
		f.Close()
	}
}

func commandCompleter(line string) (c []string) {
	var names = []string{"john", "james", "mary", "nancy"}
	for _, n := range names {
		if strings.HasPrefix(n, strings.ToLower(line)) {
			c = append(c, n)
		}
	}
	return
}
