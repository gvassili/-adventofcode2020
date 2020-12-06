package day06

import (
	"bufio"
	"bytes"
	"io"
	"strconv"
)

type group struct {
	fOrAnswer  uint32
	fAndAnswer uint32
}

type Challenge struct {
	groups []group
}

func (c *Challenge) Day() int {
	return 6
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
		var group group
		userScanner := bufio.NewScanner(bytes.NewReader(scanner.Bytes()))
		for userScanner.Scan() {
			var fUAnswer uint32
			for _, a := range userScanner.Bytes() {
				fUAnswer |= uint32(1) << (uint32(a) - 'a')
			}
			if group.fOrAnswer == 0 {
				group.fAndAnswer = fUAnswer
			} else {
				group.fAndAnswer &= fUAnswer
			}
			group.fOrAnswer |= fUAnswer
		}
		if err := userScanner.Err(); err != nil {
			return err
		}
		c.groups = append(c.groups, group)
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}

func u32PopCount(i uint32) int {
	i = i - ((i >> 1) & 0x55555555)
	i = (i & 0x33333333) + ((i >> 2) & 0x33333333)
	return int((((i + (i >> 4)) & 0x0F0F0F0F) * 0x01010101) >> 24)
}

func (c *Challenge) Part1() (string, error) {
	answerSum := 0
	for _, group := range c.groups {
		answerSum += u32PopCount(group.fOrAnswer)
	}
	return strconv.Itoa(answerSum), nil
}

func (c *Challenge) Part2() (string, error) {
	answerSum := 0
	for _, group := range c.groups {
		answerSum += u32PopCount(group.fAndAnswer)
	}
	return strconv.Itoa(answerSum), nil
}
