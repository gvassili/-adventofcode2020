package day04

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/go-playground/validator/v10"
	"io"
	"regexp"
	"strconv"
	"strings"
)

type passport struct {
	Byr int    `p1:"required" p2:"required,min=1920,max=2002"`
	Iyr int    `p1:"required" p2:"required,min=2010,max=2020"`
	Eyr int    `p1:"required" p2:"required,min=2020,max=2030"`
	Hgt string `p1:"required" p2:"required,is-height"`
	Hcl string `p1:"required" p2:"required,is-hair-color"`
	Ecl string `p1:"required" p2:"required,oneof=amb blu brn gry grn hzl oth"`
	Pid string `p1:"required" p2:"required,numeric,len=9"`
}

type Challenge struct {
	passports []passport
}

func (c *Challenge) Day() int {
	return 4
}

func (c *Challenge) Prepare(r io.Reader) error {
	scanner := bufio.NewScanner(r)
	scanner.Split(func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		if atEOF && len(data) == 0 {
			return 0, nil, nil
		}
		if i := bytes.Index(data, []byte{'\n', '\n'}); i >= 0 {
			return i + 1, data[0:i], nil
		}
		if atEOF {
			return len(data), data, nil
		}
		return 0, nil, nil
	})

	for scanner.Scan() {
		var passport passport
		passportScanner := bufio.NewScanner(bytes.NewReader(scanner.Bytes()))
		passportScanner.Split(func(data []byte, atEOF bool) (advance int, token []byte, err error) {
			if atEOF && len(data) == 0 {
				return 0, nil, nil
			}
			if i := bytes.IndexAny(data, "\n \t"); i >= 0 {
				return i + 1, data[0:i], nil
			}
			if atEOF {
				return len(data), data, nil
			}
			return 0, nil, nil
		})

		for passportScanner.Scan() {
			field := strings.SplitN(passportScanner.Text(), ":", 2)
			if len(field) != 2 {
				continue
			}
			name, value := field[0], field[1]
			switch name {
			case "byr":
				passport.Byr, _ = strconv.Atoi(value)
			case "iyr":
				passport.Iyr, _ = strconv.Atoi(value)
			case "eyr":
				passport.Eyr, _ = strconv.Atoi(value)
			case "hgt":
				passport.Hgt = value
			case "hcl":
				passport.Hcl = value
			case "ecl":
				passport.Ecl = value
			case "pid":
				passport.Pid = value
			}
		}
		if err := passportScanner.Err(); err != nil {
			return err
		}
		c.passports = append(c.passports, passport)
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}

func (c *Challenge) Part1() (string, error) {
	validate := validator.New()
	validate.SetTagName("p1")
	validCount := 0
	for _, passport := range c.passports {
		if err := validate.Struct(&passport); err == nil {
			validCount++
		}
	}
	return strconv.Itoa(validCount), nil
}

func (c *Challenge) Part2() (string, error) {
	validate := validator.New()
	validate.SetTagName("p2")
	hairColorRegex := regexp.MustCompile(`#[0-9a-f]{6}`)
	validate.RegisterValidation("is-hair-color", func(fl validator.FieldLevel) bool {
		return hairColorRegex.MatchString(fl.Field().String())
	})
	validate.RegisterValidation("is-height", func(fl validator.FieldLevel) bool {
		var h int
		var u string
		if n, err := fmt.Sscanf(fl.Field().String(), "%d%s", &h, &u); n != 2 || err != nil {
			return false
		}
		switch u {
		case "cm":
			return h >= 150 && h <= 193
		case "in":
			return h >= 59 && h <= 76
		default:
			return false
		}
	})
	validCount := 0
	for _, passport := range c.passports {
		if err := validate.Struct(&passport); err == nil {
			validCount++
		}
	}
	return strconv.Itoa(validCount), nil
}
