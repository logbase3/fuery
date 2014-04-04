/*
	filequery is a small and simple tool for querying files using SQL.
	Copyright (C) 2013 log₃() <contact@logbase3.com>

	This program is free software: you can redistribute it and/or modify
	it under the terms of the GNU Lesser General Public License as published by
	the Free Software Foundation, either version 3 of the License, or
	(at your option) any later version.

	This program is distributed in the hope that it will be useful,
	but WITHOUT ANY WARRANTY; without even the implied warranty of
	MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
	GNU Lesser General Public License for more details.

	You should have received a copy of the GNU Lesser General Public License
	along with this program. If not, see <http://www.gnu.org/licenses/>.

	For more information visit https://github.com/logbase3/fuery
	or send an e-mail to contact@logbase3.com
*/

package filequery

// This file includes all the diferent structures needed to represent a query:
// Select, Update, etc.. and the necesary types needed for representing each of
// those structures.

// Enum-Like type Operator
type Operator uint8

const (
	OperatorEqual Operator = iota
	OperatorNotEqual
	OperatorGrater
	OperatorLess
	OperatorGraterOrEqual
	OperatorLessOrEqual
	OperatorBetween
	OperatorLike
	OperatorIn
	OperatorIs
	OperatorIsNot
)

type Condition struct {
	field	int
	operator Operator
	operand  string
}

// Enum-Like type AggFunction
type AggFunction uint8

const (
	AggSum AggFunction = iota
	AggAvg
)

type Select struct {
	fields		   []int
	file			 string
	conditions	   []Condition
	aggFunctions	 []AggFunction
	groupFields	  []int
	havingConditions []Condition
	orderFields	  []int
}
