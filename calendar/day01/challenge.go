package day01

import (
	"bufio"
	"errors"
	"io"
	"strconv"
)

type Challenge struct {
	input    []int
	inputMap []int
}

func (d *Challenge) Day() int {
	return 1
}

func (c *Challenge) Prepare(r io.Reader) error {
	c.input = make([]int, 0, 2048)
	c.inputMap = make([]int, 2048)
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		n, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return err
		}
		c.input = append(c.input, n)
		c.inputMap[n] = n
	}
	return scanner.Err()
}

func (d *Challenge) Part1() (string, error) {
	for _, n := range d.input {
		t := 2020 - n
		if t >= 0 && t == d.inputMap[t] {
			return strconv.Itoa((2020 - n) * n), nil
		}
	}
	return "", errors.New("could not find result")
}

func (d *Challenge) Part2() (string, error) {
	for i, n := range d.input {
		for _, m := range d.input[i:] {
			t := 2020 - n - m
			if t > 0 && t == d.inputMap[t] {
				return strconv.Itoa(t * n * m), nil
			}
		}
	}
	return "", errors.New("could not find result")
}
