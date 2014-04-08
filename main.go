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
	"fmt"
	"github.com/kless/term/readline"
)

const HISTORY_FILE string = "my.history" // TODO: Change the config files to user home
const HEADER string = "    Fuery (File Query) Copyright (C) 2013  log₃()\n" +
	"    This program comes with ABSOLUTELY NO WARRANTY; for details type `show w'.\n" +
	"    This is free software, and you are welcome to redistribute it\n" +
	"    under certain conditions; type `show c' for details.\n" +
	"\n    For more information visit https://github.com/logbase3/fuery\n"

func main() {

	fmt.Printf("%s\n", HEADER)

	hist,_ := readline.NewHistory(HISTORY_FILE)
	hist.Load()

	ln, err := readline.NewDefaultLine(hist)
	if err != nil {
		fmt.Printf("%s", err)
		return
	}
	defer func() {
		fmt.Printf("%s\n", "Hasta luego")
		hist.Save()

		if err = ln.Restore(); err != nil {
			fmt.Printf("%s", err)
		}
	}()

	for {
		line, err := ln.Read()
		if err != nil {
			if err == readline.ErrCtrlD {
				err = nil
			} else {
				fmt.Printf("%s", err)
				return
			}
			break
			}

			fmt.Printf("%s%s\n\n", "Escribiste: ", line)

		}
}
