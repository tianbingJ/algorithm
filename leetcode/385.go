package leetcode

import (
	"fmt"
)

type NestedInteger struct {
}

// Return true if this NestedInteger holds a single integer, rather than a nested list.
func (n NestedInteger) IsInteger() bool {
	return false
}

// Return the single integer that this NestedInteger holds, if it holds a single integer
// The result is undefined if this NestedInteger holds a nested list
// So before calling this method, you should have a check
func (n NestedInteger) GetInteger() int {
	return 0
}

// Set this NestedInteger to hold a single integer.
func (n *NestedInteger) SetInteger(value int) {
}

// Set this NestedInteger to hold a nested list and adds a nested integer to it.
func (n *NestedInteger) Add(elem NestedInteger) {
}

// Return the nested list that this NestedInteger holds, if it holds a nested list
// The list length is zero if this NestedInteger holds a single integer
// You can access NestedInteger's List element directly if you want to modify it
func (n NestedInteger) GetList() []*NestedInteger {
	return nil
}

/**
类似编译原理当中的语法解析
 */
type Parser struct {
	input   string
	pos int
}

func (parser *Parser) NextByte() byte {
	return parser.input[parser.pos]
}

func (parser *Parser) ended() bool {
	return parser.pos >= len(parser.input)
}

func (parser *Parser) matchList() NestedInteger {
	result := NestedInteger{}
	parser.matchByte('[')
	for !parser.ended() {
		nextToken := parser.NextByte();
		if nextToken == ']' {
			break
		}
		switch nextToken {
		case '[':
			nexted := parser.matchList()
			result.Add(nexted)
		case ',':
			parser.matchByte(',')
			continue
		default:
			nexted := parser.matchNumber()
			result.Add(nexted)
		}
	}
	parser.matchByte(']')
	return result
}

func (parser * Parser) matchNumber() NestedInteger {
	result := 0
	neg := 1
	if parser.NextByte() == '-' {
		neg = -1
		parser.pos ++
	}
	for ; !parser.ended() && parser.NextByte() >= '0' && parser.NextByte() <= '9';  parser.pos ++{
		result = result * 10 + int(parser.NextByte()) - '0';
	}
	nested := NestedInteger{}
	nested.SetInteger(result * neg)
	//fmt.Print(result)
	return nested
}

func (parser *Parser) matchByte(b byte) {
	//fmt.Printf("%c",b)
	if parser.input[parser.pos] != b {
		panic(fmt.Sprintf("expected:%v found:%v ", b, parser.input[parser.pos]))
	}
	parser.pos ++
}

func (parser *Parser) parse() *NestedInteger {
	if parser.ended() {
		return nil
	}
	var result NestedInteger
	if parser.NextByte() == '[' {
		result = parser.matchList()
	} else {
		result = parser.matchNumber()
	}
	return &result
}

func deserialize(s string) *NestedInteger {
	parser := Parser{
		input:s,
		pos:0,
	}
	return parser.parse()
}
