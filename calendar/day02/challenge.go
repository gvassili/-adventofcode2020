package day02

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type policy struct {
	min  int
	max  int
	char byte
}

type password struct {
	policy policy
	str    string
}

type Challenge struct {
	passwords []password
}

func (d *Challenge) Day() int {
	return 2
}

func (d *Challenge) Prepare(r io.Reader) error {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		var password password
		_, err := fmt.Sscanf(scanner.Text(), "%d-%d %c: %s",
			&password.policy.min,
			&password.policy.max,
			&password.policy.char,
			&password.str)
		if err != nil {
			return err
		}
		d.passwords = append(d.passwords, password)
	}
	return scanner.Err()
}

func (d *Challenge) Part1() (string, error) {
	validCount := 0
	for _, password := range d.passwords {
		n := strings.Count(password.str, string(password.policy.char))
		if n <= password.policy.max && n >= password.policy.min {
			validCount++
		}
	}
	return strconv.Itoa(validCount), nil
}

func (d *Challenge) Part2() (string, error) {
	validCount := 0
	for _, password := range d.passwords {
		if (password.str[password.policy.min-1] == password.policy.char) != // xor
			(password.str[password.policy.max-1] == password.policy.char) {
			validCount++
		}
	}
	return strconv.Itoa(validCount), nil
}
