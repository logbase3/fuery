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

package main

const (
	// This constant is just the name of the file and should not be used for
	// opening the file. Instead the HISTORY_FILE variable should be used.
	HISTORY_FILE_NAME string = "/.fuery_history"

	// Current program version.
	VERSION string = "0.1"

	// Current year
	YEAR string = "2015"

	// Program header
	HEADER string = "    Fuery v" + VERSION + " Copyright (C) " + YEAR + " log₃()\n" +
		"    This program comes with ABSOLUTELY NO WARRANTY.\n" +
		"    This is free software, and you are welcome to redistribute it\n" +
		"    under certain conditions. Read the license file or visit:\n" +
		"    'http://www.gnu.org/copyleft/gpl.html' for more information.\n" +
		"\n    You can type 'help' or '\\h' to get a list of options. For more\n" +
		"    information visit https://github.com/logbase3/fuery\n"
)
