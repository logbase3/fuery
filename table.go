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

// Package fuery is a small and simple library for querying files using SQL.
package fuery // import "logbase3.com/fuery"

import (
	"bytes"
	"container/list"
	"fmt"
	"io"
	"strings"
	"unicode/utf8"
)

// Constants for output configuration
const (
	separator                 = " | "
	lineSeparator             = "-+-"
	lineCharacter             = '-'
	defaultCellFormat         = "%%-%ds"
	numericCellFormat         = "%%%ds"
	defaultColumnNameTemplate = "Column %d"
)

type Record struct {
	ParentTable *Table
	Cells       []DataType
}

type Table struct {
	names   []string
	types   []Type
	Records *list.List
}

func NewTable(dataTypes ...Type) *Table {
	names := make([]string, len(dataTypes))
	for i := range dataTypes {
		names[i] = fmt.Sprintf(defaultColumnNameTemplate, i)
	}
	table := &Table{names, dataTypes, list.New()}
	return table
}

func (t Table) Names() []string {
	return t.names
}

func (t Table) Types() []Type {
	return t.types
}

func (t *Table) SetNames(names []string) {
	if len(names) != len(t.types) {
		panic(fmt.Sprintf("Invalid name assignment: Names must be a slice with length %d", len(t.types)))
	}
	t.names = names
}

func (t Table) Insert(values ...DataType) {
	if len(values) > 0 {
		cells := make([]DataType, 0, len(t.types))
		cells = append(cells, values...)
		t.InsertRecords(Record{&t, cells})
	}
}

func (t Table) InsertRecords(records ...Record) {
	for _, record := range records {
		t.Records.PushBack(record)
	}
}

func (t Table) maxCellLength() []int {

	// Bug(Roberto Lapuente): Should use column names instead of numbers where available
	// Initialize slice with the lenght of the column names
	lengths := make([]int, 0, len(t.types))
	for colNumber := range t.types {
		lengths = append(lengths, len(t.Names()[colNumber]))
	}
	for e := t.Records.Front(); e != nil; e = e.Next() {
		for i, cell := range e.Value.(Record).Cells {
			colLength := utf8.RuneCountInString(cell.String())
			if colLength > lengths[i] {
				lengths[i] = colLength
			}
		}
	}
	return lengths
}

func (t Table) String() string {
	var buff bytes.Buffer
	t.Write(&buff)
	return buff.String()
}

func (t Table) Write(buff io.Writer) {
	// Construct format string with column sizes
	formatSlice := make([]string, 0, len(t.types))
	separatorFormatSlice := make([]string, 0, len(t.types))
	for colName, length := range t.maxCellLength() {
		if t.types[colName] == INT {
			formatSlice = append(formatSlice, fmt.Sprintf(numericCellFormat, length))
		} else if t.types[colName] == TEXT {
			formatSlice = append(formatSlice, fmt.Sprintf(defaultCellFormat, length))
		}
		separatorFormatSlice = append(separatorFormatSlice, fmt.Sprintf(defaultCellFormat, length))
	}
	separatorFormatString := fmt.Sprintf(" %s ", strings.Join(separatorFormatSlice, separator))
	headerFormatString := fmt.Sprintf("-%s-", strings.Join(formatSlice, lineSeparator))
	formatString := fmt.Sprintf(" %s ", strings.Join(formatSlice, separator))

	// Start appending results
	var row []interface{}

	// Build header
	for colName, length := range t.maxCellLength() {
		diff := length - len(t.Names()[colName])
		headerFormatGap := fmt.Sprintf(defaultCellFormat, diff/2)
		headerFormatGap = fmt.Sprintf(headerFormatGap, "")
		row = append(row, fmt.Sprintf("%s%s", headerFormatGap, t.Names()[colName]))
	}
	buff.Write([]byte(fmt.Sprintf(separatorFormatString, row...)))
	buff.Write([]byte("\n"))

	// Build header/body separator
	row = make([]interface{}, 0, len(t.types))
	for _, length := range t.maxCellLength() {
		column := make([]byte, length)
		for i := 0; i < length; i++ {
			column[i] += lineCharacter
		}
		row = append(row, string(column))
	}
	buff.Write([]byte(fmt.Sprintf(headerFormatString, row...)))
	buff.Write([]byte("\n"))

	// Build rows
	for e := t.Records.Front(); e != nil; e = e.Next() {
		row = make([]interface{}, 0, len(t.types))
		for _, cell := range e.Value.(Record).Cells {
			row = append(row, cell)
		}
		buff.Write([]byte(fmt.Sprintf(formatString, row...)))
		buff.Write([]byte("\n"))
	}
}
