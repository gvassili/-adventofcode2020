package day07

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"testing"
)

const input = `light red bags contain 1 bright white bag, 2 muted yellow bags.
dark orange bags contain 3 bright white bags, 4 muted yellow bags.
bright white bags contain 1 shiny gold bag.
muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.
shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.
dark olive bags contain 3 faded blue bags, 4 dotted black bags.
vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.
faded blue bags contain no other bags.
dotted black bags contain no other bags.`

var fullInput = func() []byte {
	r, err := os.Open("./input")
	if err != nil {
		panic(err)
	}
	b, err := ioutil.ReadAll(r)
	if err != nil {
		panic(err)
	}
	return b
}()

func TestChallenge_Prepare(t *testing.T) {
	var c Challenge
	err := c.Prepare(bytes.NewReader([]byte(input)))
	assert.NoError(t, err)
	assert.Equal(t, map[string][]content{
		"light red":    {{1, "bright white"}, {2, "muted yellow"}},
		"dark orange":  {{3, "bright white"}, {4, "muted yellow"}},
		"bright white": {{1, "shiny gold"}},
		"muted yellow": {{2, "shiny gold"}, {9, "faded blue"}},
		"shiny gold":   {{1, "dark olive"}, {2, "vibrant plum"}},
		"dark olive":   {{3, "faded blue"}, {4, "dotted black"}},
		"vibrant plum": {{5, "faded blue"}, {6, "dotted black"}},
	}, c.containers)
	assert.Equal(t, map[string][]string{
		"bright white": {"light red", "dark orange"},
		"muted yellow": {"light red", "dark orange"},
		"shiny gold":   {"bright white", "muted yellow"},
		"faded blue":   {"muted yellow", "dark olive", "vibrant plum"},
		"dark olive":   {"shiny gold"},
		"vibrant plum": {"shiny gold"},
		"dotted black": {"dark olive", "vibrant plum"},
	}, c.recipes)
}

func BenchmarkChallenge_Prepare(b *testing.B) {
	for i := 0; i < b.N; i++ {
		buf := bytes.NewReader(fullInput)
		var c Challenge
		c.Prepare(buf)
	}
}

func TestChallenge_Part1(t *testing.T) {
	var c Challenge
	err := c.Prepare(bytes.NewReader([]byte(input)))
	assert.NoError(t, err)
	r, err := c.Part1()
	assert.NoError(t, err)
	assert.Equal(t, "4", r)
}

func BenchmarkChallenge_Part1(b *testing.B) {
	buf := bytes.NewReader(fullInput)
	var c Challenge
	c.Prepare(buf)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		c.Part1()
	}
}

func TestChallenge_Part2(t *testing.T) {
	var c Challenge
	err := c.Prepare(bytes.NewReader([]byte(input)))
	assert.NoError(t, err)
	r, err := c.Part2()
	assert.NoError(t, err)
	assert.Equal(t, "32", r)
}

func BenchmarkChallenge_Part2(b *testing.B) {
	buf := bytes.NewReader([]byte(input))
	var c Challenge
	c.Prepare(buf)
	for i := 0; i < b.N; i++ {
		c.Part2()
	}
}
