/*
	Fuery, is a small and simple library for querying files using SQL.
	Copyright (C) 2016 log₃() <contact@logbase3.com>

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

	For more information visit https://logbase3.com/fuery
	or send an e-mail to contact@logbase3.com
*/

// Package Fuery, is a small and simple library for querying files using SQL.
package fuery // import "logbase3.com/fuery"

import (
	"fmt"
	"strings"
)

type DataType int

const (
	INT int = iota
	TEXT
)

// Constants for output configuration
const (
	separator       = " | "
	headerSeparator = "-+-"
	headerCharacter = "-"
	generalFormat   = "%%-%ds"
	numericFormat   = "%%%ds"
	columnTemplate  = "Column %d"
)

type Table struct {
	table [][]string
	types []DataType
}

func NewTable() *Table {
	table := make([][]string, 0, 3)
	table = append(table, []string{"5", "6", "8", "9"})
	table = append(table, []string{"Hola", "Atun", "Caca", "Ricas fresas"})
	table = append(table, []string{"Adios", "Con carne", "Para comer", "Con crema rica"})

	types := []DataType{0, 1, 1}

	return &Table{table, types}
}

func (t *Table) maxCellLength() (lengths []int) {
	var length int
	lengths = make([]int, 0, len(t.table))

	for column := range t.table {
		length = len(fmt.Sprintf(columnTemplate, column))
		for _, elem := range t.table[column] {
			if len(elem) > length {
				length = len(elem)
			}
		}
		lengths = append(lengths, length)
	}
	return
}

func (t *Table) columnMaxLength(column int) (length int) {
	for _, elem := range t.table[column] {
		if len(elem) > length {
			length = len(elem)
		}
	}
	return
}

func (t *Table) String() string {
	res := make([]string, 0, len(t.table[0]))

	// Construct format string with column sizes
	formatSlice := make([]string, 0, len(t.table))
	separatorFormatSlice := make([]string, 0, len(t.table))
	for column, length := range t.maxCellLength() {
		if t.types[column] == 0 {
			formatSlice = append(formatSlice, fmt.Sprintf(numericFormat, length))
		} else {
			formatSlice = append(formatSlice, fmt.Sprintf(generalFormat, length))
		}
		separatorFormatSlice = append(separatorFormatSlice, fmt.Sprintf(generalFormat, length))
	}
	separatorFormatString := fmt.Sprintf(" %s ", strings.Join(separatorFormatSlice, separator))
	headerFormatString := fmt.Sprintf("-%s-", strings.Join(formatSlice, headerSeparator))
	formatString := fmt.Sprintf(" %s ", strings.Join(formatSlice, separator))

	// Start appending results
	var row []interface{}

	// Build header
	for column, length := range t.maxCellLength() {
		diff := length - len(fmt.Sprintf(columnTemplate, column))
		headerFormatGap := fmt.Sprintf(generalFormat, diff/2)
		headerFormatGap = fmt.Sprintf(headerFormatGap, "")
		row = append(row, headerFormatGap+fmt.Sprintf(columnTemplate, column))
	}
	res = append(res, fmt.Sprintf(separatorFormatString, row...))

	// Build header/body separator
	var column string
	row = make([]interface{}, 0, len(t.table))
	for _, length := range t.maxCellLength() {
		column = ""
		for i := 0; i < length; i++ {
			column += headerCharacter
		}
		row = append(row, column)
	}
	res = append(res, fmt.Sprintf(headerFormatString, row...))

	// Build row
	for y := 0; y < len(t.table[0]); y++ {
		row = make([]interface{}, 0, len(t.table))
		for x := 0; x < len(t.table); x++ {
			row = append(row, t.table[x][y])
		}
		res = append(res, fmt.Sprintf(formatString, row...))
	}
	return strings.Join(res, "\n")
}
