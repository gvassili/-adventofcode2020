package day08

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"strconv"
)

type opCode int

const (
	opJmp opCode = iota
	opAcc
	opNop
)

type instruction struct {
	op    opCode
	param int
}

type Challenge struct {
	instructions []instruction
}

func (c *Challenge) Day() int {
	return 8
}

func (c *Challenge) Prepare(r io.Reader) error {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		var mnemonic string
		var param int
		i, err := fmt.Fscanf(bytes.NewReader(scanner.Bytes()), "%3s %d", &mnemonic, &param)
		if err != nil {
			return err
		} else if i != 2 {
			return errors.New("invalid number of token retrieve by scanf")
		}
		var ins instruction
		switch mnemonic {
		case "jmp":
			ins.op = opJmp
		case "acc":
			ins.op = opAcc
		default:
			ins.op = opNop
		}
		ins.param = param
		c.instructions = append(c.instructions, ins)
	}
	return scanner.Err()
}

func (c *Challenge) step(acc, pc int) (int, int) {

	return acc, pc
}

func (c *Challenge) run() (int, bool, bool) {
	acc, pc := 0, 0
	visited := make([]bool, len(c.instructions))
	for {
		if pc == len(c.instructions) {
			return acc, false, false
		} else if pc < 0 || pc > len(c.instructions) {
			return acc, false, true
		}
		if visited[pc] {
			return acc, true, false
		}
		visited[pc] = true
		ins := c.instructions[pc]
		switch ins.op {
		case opJmp:
			pc += ins.param
		case opAcc:
			acc += ins.param
			fallthrough
		default:
			pc++
		}
	}
}

func (c *Challenge) Part1() (string, error) {
	acc, looped, err := c.run()
	if !looped || err {
		return "", errors.New("program didn't loop as expected")
	}
	return strconv.Itoa(acc), nil
}

func (c *Challenge) Part2() (string, error) {
	for i := 0; i < len(c.instructions); i++ {
		tmpOp := c.instructions[i].op
		if tmpOp == opJmp {
			c.instructions[i].op = opNop
		} else if tmpOp == opNop {
			c.instructions[i].op = opJmp
		} else {
			continue
		}
		if r, looped, err := c.run(); !err && !looped {
			return strconv.Itoa(r), nil
		}
		c.instructions[i].op = tmpOp
	}
	return "", errors.New("could not find corrupted instruction")
}
