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

import "strconv"

type Type int

const (
	INT Type = iota
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

type DataType interface {
	String() string
}
