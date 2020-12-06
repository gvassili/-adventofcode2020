package day03

import (
	"bufio"
	"io"
	"strconv"
)

type Challenge struct {
	slope  []byte
	width  int
	height int
}

func (c *Challenge) Day() int {
	return 3
}

func (c *Challenge) Prepare(r io.Reader) error {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		row := scanner.Bytes()
		if c.width == 0 {
			c.width = len(row)
		}
		c.slope = append(c.slope, row...)
		c.height++
	}
	return scanner.Err()
}

func (c *Challenge) slide(dy, dx int) int {
	impact := 0
	for y, x := 0, 0; y < c.height; y, x = y+dy, x+dx {
		if c.slope[(x%c.width)+(y*c.width)] == '#' {
			impact++
		}
	}
	return impact
}

func (c *Challenge) Part1() (string, error) {
	return strconv.Itoa(c.slide(1, 3)), nil
}

func (c *Challenge) Part2() (string, error) {
	return strconv.Itoa(
		c.slide(1, 1) *
			c.slide(1, 3) *
			c.slide(1, 5) *
			c.slide(1, 7) *
			c.slide(2, 1),
	), nil
}
