package day18

import (
	"bufio"
	"io"
	"strconv"
)

type expression []interface{}

type Challenge struct {
	expressions []expression
}

func (c *Challenge) Day() int {
	return 18
}

func (c *Challenge) Prepare(r io.Reader) error {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		var expression expression
		nBeg := -1
		for i, c := range scanner.Bytes() {
			switch {
			case c >= '0' && c <= '9':
				if nBeg == -1 {
					nBeg = i
				}
			case nBeg != -1:
				n, err := strconv.Atoi(string(scanner.Bytes()[nBeg:i]))
				if err != nil {
					return err
				}
				expression = append(expression, n)
				nBeg = -1
				if c != ' ' {
					expression = append(expression, c)
				}
			case c == ' ':
			default:
				expression = append(expression, c)
			}
		}
		if nBeg != -1 {
			n, err := strconv.Atoi(string(scanner.Bytes()[nBeg:]))
			if err != nil {
				return err
			}
			expression = append(expression, n)
		}
		c.expressions = append(c.expressions, expression)
	}

	return scanner.Err()
}

func (c *Challenge) Part1() (string, error) {
	sum := 0
	for _, ex := range c.expressions {
		rpn := make(expression, 0, len(c.expressions))
		opStack := make([]byte, 0)

		for _, tok := range ex {
			switch t := tok.(type) {
			case int:
				rpn = append(rpn, t)
			case byte:
				switch t {
				case '+':
					for len(opStack) > 0 && opStack[len(opStack)-1] != '(' {
						rpn = append(rpn, opStack[len(opStack)-1])
						opStack = opStack[:len(opStack)-1]
					}
					opStack = append(opStack, t)
				case '*':
					for len(opStack) > 0 && opStack[len(opStack)-1] != '(' {
						rpn = append(rpn, opStack[len(opStack)-1])
						opStack = opStack[:len(opStack)-1]
					}
					opStack = append(opStack, t)
				case ')':
					for opStack[len(opStack)-1] != '(' {
						rpn = append(rpn, opStack[len(opStack)-1])
						opStack = opStack[:len(opStack)-1]
					}
					opStack = opStack[:len(opStack)-1]
				default:
					opStack = append(opStack, t)
				}
			}
		}
		for len(opStack) > 0 {
			if opStack[len(opStack)-1] != '(' {
				rpn = append(rpn, opStack[len(opStack)-1])
			}
			opStack = opStack[:len(opStack)-1]
		}
		numberStack := make([]int, 0)
		for len(rpn) > 0 {
			switch t := rpn[0].(type) {
			case int:
				numberStack = append(numberStack, t)
			case byte:
				if t == '+' {
					numberStack[len(numberStack)-2] = numberStack[len(numberStack)-1] + numberStack[len(numberStack)-2]
				} else if t == '*' {
					numberStack[len(numberStack)-2] = numberStack[len(numberStack)-1] * numberStack[len(numberStack)-2]
				}
				numberStack = numberStack[:len(numberStack)-1]
			}
			rpn = rpn[1:]
		}
		sum += numberStack[0]
	}
	return strconv.Itoa(sum), nil
}

func (c *Challenge) Part2() (string, error) {
	sum := 0
	for _, ex := range c.expressions {
		rpn := make(expression, 0, len(c.expressions))
		opStack := make([]byte, 0)

		for _, tok := range ex {
			switch t := tok.(type) {
			case int:
				rpn = append(rpn, t)
			case byte:
				switch t {
				case '+':
					for len(opStack) > 0 && opStack[len(opStack)-1] == '+' {
						rpn = append(rpn, opStack[len(opStack)-1])
						opStack = opStack[:len(opStack)-1]
					}
					opStack = append(opStack, t)
				case '*':
					for len(opStack) > 0 && opStack[len(opStack)-1] != '(' {
						rpn = append(rpn, opStack[len(opStack)-1])
						opStack = opStack[:len(opStack)-1]
					}
					opStack = append(opStack, t)
				case ')':
					for opStack[len(opStack)-1] != '(' {
						rpn = append(rpn, opStack[len(opStack)-1])
						opStack = opStack[:len(opStack)-1]
					}
					opStack = opStack[:len(opStack)-1]
				default:
					opStack = append(opStack, t)
				}
			}
		}
		for len(opStack) > 0 {
			if opStack[len(opStack)-1] != '(' {
				rpn = append(rpn, opStack[len(opStack)-1])
			}
			opStack = opStack[:len(opStack)-1]
		}
		numberStack := make([]int, 0)
		for len(rpn) > 0 {
			switch t := rpn[0].(type) {
			case int:
				numberStack = append(numberStack, t)
			case byte:
				if t == '+' {
					numberStack[len(numberStack)-2] = numberStack[len(numberStack)-1] + numberStack[len(numberStack)-2]
				} else if t == '*' {
					numberStack[len(numberStack)-2] = numberStack[len(numberStack)-1] * numberStack[len(numberStack)-2]
				}
				numberStack = numberStack[:len(numberStack)-1]
			}
			rpn = rpn[1:]
		}
		sum += numberStack[0]
	}
	return strconv.Itoa(sum), nil
}
