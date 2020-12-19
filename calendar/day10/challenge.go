package day10

import (
	"bufio"
	"errors"
	"io"
	"sort"
	"strconv"
)

type Challenge struct {
	adapters []int
}

func (c *Challenge) Day() int {
	return 10
}

func (c *Challenge) Prepare(r io.Reader) error {
	scanner := bufio.NewScanner(r)
	//	c.adapters = append(c.adapters, 0)
	for scanner.Scan() {
		n, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return err
		}
		c.adapters = append(c.adapters, n)
	}
	sort.Ints(c.adapters)
	//c.adapters = append(c.adapters, c.adapters[len(c.adapters)-1]+3)
	return scanner.Err()
}

func (c *Challenge) Part1() (string, error) {
	jolt1Count, jolt3Count := 1, 1
	for i := 1; i < len(c.adapters); i++ {
		diff := c.adapters[i] - c.adapters[i-1]
		if diff == 1 {
			jolt1Count++
		} else if diff == 3 {
			jolt3Count++
		} else {
			return "", errors.New("invalid jolt difference")
		}
	}
	return strconv.Itoa(jolt1Count * jolt3Count), nil
}

func (c *Challenge) Part2() (string, error) {
	routeLengths := make(map[int]int, len(c.adapters))
	routeLengths[0] = 1
	for _, joltage := range c.adapters {
		var totalRoutes int
		for _, n := range []int{1, 2, 3} {
			totalRoutes += routeLengths[joltage-n]
		}
		routeLengths[joltage] = totalRoutes
	}
	return strconv.Itoa(routeLengths[c.adapters[len(c.adapters)-1]]), nil
}
