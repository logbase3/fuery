package main

// This file includes all the diferent structures needed to represent a query:
// Select, Update, etc.. and the necesary types needed for representing each of
// those structures.

// Declaration of type Field. Just integers indicating the index of the field,
// but negative numbers are reseved for aggregate functions.
type Field uint8
const (
	FieldSum = iota
	FieldAvg
)

type Select struct {
	fields           []Field
	files            []string
	conditions       []Condition
	groupFields      []Field
	havingConditions []Condition
	orderFields      []Field
}

// Declaration of Enum-Like type Operator
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
	field    Field
	operator Operator
	operand  string
	operand  string
}
