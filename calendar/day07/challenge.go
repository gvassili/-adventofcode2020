package day07

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

type content struct {
	amount int
	name   string
}

type Challenge struct {
	containers map[string][]content
	recipes    map[string][]string
}

func (c *Challenge) Day() int {
	return 7
}

func sanitizeBagName(name string) string {
	toks := strings.Split(name, " ")
	return toks[0] + " " + toks[1]
}

func (c *Challenge) Prepare(r io.Reader) error {
	c.containers = make(map[string][]content)
	c.recipes = make(map[string][]string)
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		t := scanner.Text()
		si := strings.Index(t, " contain ")
		container := sanitizeBagName(t[:si])
		cont := t[si+len(" contains ")-1 : len(t)-1]
		if cont[:2] != "no" {
			items := strings.Split(cont, ", ")
			for _, item := range items {
				name := sanitizeBagName(item[2:])
				amount, err := strconv.Atoi(item[:1])
				if err != nil {
					return err
				}
				c.containers[container] = append(c.containers[container], content{amount, name})
				c.recipes[name] = append(c.recipes[name], container)
			}
		}
	}
	return scanner.Err()
}

func (c *Challenge) Part1() (string, error) {
	seenItem := make(map[string]struct{})
	var scanContent func(string, []string, int)
	scanContent = func(name string, containers []string, p int) {
		for _, container := range containers {
			seenItem[container] = struct{}{}
			scanContent(container, c.recipes[container], p+1)
		}
	}
	scanContent("shiny gold", c.recipes["shiny gold"], 0)
	return strconv.Itoa(len(seenItem)), nil
}

func (c *Challenge) Part2() (string, error) {
	var scanContent func(string, []content, int) int
	scanContent = func(name string, contents []content, p int) int {
		sum := 0
		for _, content := range contents {
			sum += scanContent(content.name, c.containers[content.name], p+1)*content.amount + content.amount
		}
		return sum
	}
	return strconv.Itoa(scanContent("shiny gold", c.containers["shiny gold"], 0)), nil
}
