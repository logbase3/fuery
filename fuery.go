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
	"strconv"
	"strings"
	"unicode/utf8"
)

type DataType int

const (
	INT DataType = iota
	TEXT
)

type Int int64

func (t Int) String() string {
	return strconv.Itoa(int(t))
}

type Text string

func (t Text) String() string {
	return string(t)
}

type Record struct {
	ParentTable *Table
	Cells       []fmt.Stringer
}

type Table struct {
	Types   []DataType
	Records *list.List
}

func (t Table) Insert(values ...fmt.Stringer) {
	if len(values) > 0 {
		cells := make([]fmt.Stringer, 0, len(t.Types))
		cells = append(cells, values...)
		t.InsertRecords(Record{&t, cells})
	}
}

func (t Table) InsertRecords(records ...Record) {
	for _, record := range records {
		t.Records.PushBack(record)
	}
}

func NewTable(dataTypes ...DataType) *Table {
	table := &Table{dataTypes, list.New()}
	return table
}

// Constants for output configuration
const (
	separator       = " | "
	headerSeparator = "-+-"
	headerCharacter = "-"
	generalFormat   = "%%-%ds"
	numericFormat   = "%%%ds"
	columnTemplate  = "Column %d"
)

func (t Table) maxCellLength() []int {
	// Bug(Roberto Lapuente): Should use column names instead of numbers where available
	// Initialize slice with the lenght of the column names
	lengths := make([]int, 0, len(t.Types))
	for colNumber := range t.Types {
		lengths = append(lengths, len(fmt.Sprintf(columnTemplate, colNumber)))
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
	formatSlice := make([]string, 0, len(t.Types))
	separatorFormatSlice := make([]string, 0, len(t.Types))
	for column, length := range t.maxCellLength() {
		if t.Types[column] == INT {
			formatSlice = append(formatSlice, fmt.Sprintf(numericFormat, length))
		} else if t.Types[column] == TEXT {
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
	buff.Write([]byte(fmt.Sprintf(separatorFormatString, row...)))
	buff.Write([]byte("\n"))

	// Build header/body separator
	var column string
	row = make([]interface{}, 0, len(t.Types))
	for _, length := range t.maxCellLength() {
		column = ""
		for i := 0; i < length; i++ {
			column += headerCharacter
		}
		row = append(row, column)
	}
	buff.Write([]byte(fmt.Sprintf(headerFormatString, row...)))
	buff.Write([]byte("\n"))

	// Build rows
	for e := t.Records.Front(); e != nil; e = e.Next() {
		row = make([]interface{}, 0, len(t.Types))
		for _, cell := range e.Value.(Record).Cells {
			//s := cell.(fmt.Stringer)
			row = append(row, cell)
		}
		//buff.Write([]byte(fmt.Sprintf(formatString, Unpack(record.Cells...))))
		buff.Write([]byte(fmt.Sprintf(formatString, row...)))
		buff.Write([]byte("\n"))
	}
}

func Unpack(stuff ...fmt.Stringer) []interface{} {
	var a []interface{}
	for s := range stuff {
		a = append(a, s)
	}
	return a
}
