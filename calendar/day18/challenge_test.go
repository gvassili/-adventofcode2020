package day18

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"testing"
)

const inputSimple = `1 + 2 * 3 + 4 * 5 + 6`
const input = `2 * 3 + (4 * 5)
5 + (8 * 3 + 9 + 3 * 4 * 3)
5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))
((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2`

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
	err := c.Prepare(bytes.NewReader([]byte(inputSimple)))
	assert.NoError(t, err)
	assert.Equal(t, []expression{{1, byte('+'), 2, byte('*'), 3, byte('+'), 4, byte('*'), 5, byte('+'), 6}}, c.expressions)
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
	assert.Equal(t, "26335", r)
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
	assert.Equal(t, "693891", r)
}

func BenchmarkChallenge_Part2(b *testing.B) {
	buf := bytes.NewReader(fullInput)
	var c Challenge
	c.Prepare(buf)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		c.Part2()
	}
}
