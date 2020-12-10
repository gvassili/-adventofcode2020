package day09

import (
	"bufio"
	"errors"
	"io"
	"math"
	"strconv"
)

type Challenge struct {
	numbers []int
}

func (c *Challenge) Day() int {
	return 9
}

func (c *Challenge) Prepare(r io.Reader) error {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		n, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return err
		}
		c.numbers = append(c.numbers, n)
	}
	return scanner.Err()
}

func preambleContainSum(preamble []int, x int) bool {
	for i, n := range preamble {
		if n > x {
			continue
		}
		for _, m := range preamble[i:] {
			if n+m == x {
				return true
			}
		}
	}
	return false
}

func (c *Challenge) Part1() (string, error) {
loop:
	for i, n := range c.numbers[25:] {
		preamble := c.numbers[i : 25+i]
		for i, m := range preamble {
			if m > n {
				continue
			}
			for _, w := range preamble[i:] {
				if m+w == n {
					continue loop
				}
			}
		}
		return strconv.Itoa(n), nil
	}
	return "", errors.New("could not find not sum number")
}

func minInt(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func maxInt(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func (c *Challenge) Part2() (string, error) {
	for i := range c.numbers {
		sum, min, max := 0, math.MaxInt64, 0
		for y, m := range c.numbers[i:] {
			sum += m
			min = minInt(min, m)
			max = maxInt(max, m)
			if sum > 57195069 {
				break
			} else if y != 0 && sum == 57195069 {
				return strconv.Itoa(min + max), nil
			}
		}
	}
	return "", errors.New("could not find not sum number")
}
