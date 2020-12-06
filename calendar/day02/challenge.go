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

func (c *Challenge) Day() int {
	return 2
}

func (c *Challenge) Prepare(r io.Reader) error {
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
		c.passwords = append(c.passwords, password)
	}
	return scanner.Err()
}

func (c *Challenge) Part1() (string, error) {
	validCount := 0
	for _, password := range c.passwords {
		n := strings.Count(password.str, string(password.policy.char))
		if n <= password.policy.max && n >= password.policy.min {
			validCount++
		}
	}
	return strconv.Itoa(validCount), nil
}

func (c *Challenge) Part2() (string, error) {
	validCount := 0
	for _, password := range c.passwords {
		if (password.str[password.policy.min-1] == password.policy.char) != // xor
			(password.str[password.policy.max-1] == password.policy.char) {
			validCount++
		}
	}
	return strconv.Itoa(validCount), nil
}
