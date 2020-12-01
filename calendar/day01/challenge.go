package day01

import (
	"bufio"
	"errors"
	"io"
	"sort"
	"strconv"
)

type Day01 struct {
	input []int
}

func (d *Day01) Day() int {
	return 1
}

func (d *Day01) Prepare(r io.Reader) error {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		n, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return err
		}
		d.input = append(d.input, n)
	}
	sort.Ints(d.input)
	return scanner.Err()
}

// did some test with binary search or map, sample is too small of optimisation, just sorting the array take two time longer, however it's worth it for part 2
func (d *Day01) Part1() (string, error) {
	for i, n := range d.input {
		subInput := d.input[i:]
		j := sort.SearchInts(subInput, 2020-n)
		if j != len(subInput) && subInput[j] == 2020-n {
			return strconv.Itoa(n * subInput[j]), nil
		}
	}
	return "", errors.New("could not find result")
}

func (d *Day01) Part2() (string, error) {
	for i, n := range d.input {
		for j, m := range d.input[i:] {
			r := 2020 - n - m
			if r < 0 {
				continue
			}
			subInput := d.input[i+j:]
			w := sort.SearchInts(subInput, r)
			if w != len(subInput) && subInput[w] == r {
				return strconv.Itoa(n * m * subInput[w]), nil
			}
		}
	}
	return "", errors.New("could not find result")
}
