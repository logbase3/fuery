/*
	Fuery (File Query) Is a small and simple tool for querying files using SQL.
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

package main
import (
	"code.google.com/p/go-gnureadline"
	"fmt"
	"os"
)

const HISTORY_FILE string = "my.history" // TODO: Change the config files to user home

func main() {
	header := "    fuery (File Query) Copyright (C) 2013  log₃()\n" +
	"    This program comes with ABSOLUTELY NO WARRANTY; for details type `show w'.\n" +
	"    This is free software, and you are welcome to redistribute it\n" +
	"    under certain conditions; type `show c' for details.\n" +
	"\n    For more information visit https://github.com/logbase3/fuery\n"

	fmt.Printf("%s\n", header)

	var err error

	term := os.ExpandEnv("TERM")
	gnureadline.ReadHistory(HISTORY_FILE)
	gnureadline.StifleHistory(100)  // Maximum 10 history entries
	gnureadline.ReadInitFile(".inputrc")  // Read in a keybinding initialization file
	line := ""

	for i:=1; err == nil && line != "quit"; i++ {
			line, err = gnureadline.Readline(fmt.Sprintf("fuery> "), true)
			switch line {
				case "vi":
						gnureadline.Rl_editing_mode_set(gnureadline.Vi)
				case "emacs":
						gnureadline.Rl_editing_mode_set(gnureadline.Emacs)
				case "insert":
						gnureadline.Rl_insert_mode_set(true)
				case "overwrite":
						gnureadline.Rl_insert_mode_set(false)
				}
				fmt.Printf("You typed: %s\n", line)
	}
	fmt.Printf("History length %d\n",  gnureadline.HistoryLength())
	fmt.Printf("History max entries %d\n",  gnureadline.HistoryMaxEntries())
	gnureadline.WriteHistory(HISTORY_FILE)
	gnureadline.Rl_reset_terminal(term)
}
