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

func (d *Challenge) Day() int {
	return 3
}

func (d *Challenge) Prepare(r io.Reader) error {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		row := scanner.Bytes()
		if d.width == 0 {
			d.width = len(row)
		}
		d.slope = append(d.slope, row...)
		d.height++
	}
	return scanner.Err()
}

func (d *Challenge) slide(dy, dx int) int {
	impact := 0
	for y, x := 0, 0; y < d.height; y, x = y+dy, x+dx {
		if d.slope[(x%d.width)+(y*d.width)] == '#' {
			impact++
		}
	}
	return impact
}

func (d *Challenge) Part1() (string, error) {
	return strconv.Itoa(d.slide(1, 3)), nil
}

func (d *Challenge) Part2() (string, error) {
	return strconv.Itoa(
		d.slide(1, 1) *
			d.slide(1, 3) *
			d.slide(1, 5) *
			d.slide(1, 7) *
			d.slide(2, 1),
	), nil
}
