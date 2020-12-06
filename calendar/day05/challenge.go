package day05

import (
	"bufio"
	"errors"
	"io"
	"sort"
	"strconv"
)

type Challenge struct {
	seatIds []int
}

func (c *Challenge) Day() int {
	return 5
}

func (c *Challenge) Prepare(r io.Reader) error {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		var id int
		code := scanner.Bytes()
		for i := len(code) - 1; i >= 0; i-- {
			if code[i] == 'B' || code[i] == 'R' {
				id |= 1 << (len(code) - i - 1)
			}
		}
		c.seatIds = append(c.seatIds, id)
	}
	return scanner.Err()
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func (c *Challenge) Part1() (string, error) {
	var maxId int
	for _, id := range c.seatIds {
		maxId = max(maxId, id)
	}
	return strconv.Itoa(maxId), nil
}

func (c *Challenge) Part2() (string, error) {
	sort.Ints(c.seatIds)
	for i, id := range c.seatIds {
		if i > 0 && id != c.seatIds[i-1]+1 {
			return strconv.Itoa(id - 1), nil
		}
	}
	return "", errors.New("not found")
}
