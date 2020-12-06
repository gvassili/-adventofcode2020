package day02

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"testing"
)

const input = `1-3 a: abcde
1-3 b: cdefg
2-9 c: ccccccccc`

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

func TestDayChallenge_Prepare(t *testing.T) {
	var challenge Challenge
	err := challenge.Prepare(bytes.NewReader([]byte(input)))
	assert.NoError(t, err)
	assert.Equal(t, []password{
		{policy{1, 3, 'a'}, "abcde"},
		{policy{1, 3, 'b'}, "cdefg"},
		{policy{2, 9, 'c'}, "ccccccccc"},
	}, challenge.passwords)
}

func BenchmarkChallenge_Prepare(b *testing.B) {
	for i := 0; i < b.N; i++ {
		buf := bytes.NewReader(fullInput)
		var challenge Challenge
		challenge.Prepare(buf)
	}
}

func TestDay02_Part1(t *testing.T) {
	var challenge Challenge
	err := challenge.Prepare(bytes.NewReader([]byte(input)))
	assert.NoError(t, err)
	r, err := challenge.Part1()
	assert.NoError(t, err)
	assert.Equal(t, "2", r)
}

func BenchmarkChallenge_Part1(b *testing.B) {
	buf := bytes.NewReader(fullInput)
	var challenge Challenge
	challenge.Prepare(buf)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		challenge.Part1()
	}
}

func TestChallenge_Part2(t *testing.T) {
	var challenge Challenge
	err := challenge.Prepare(bytes.NewReader([]byte(input)))
	assert.NoError(t, err)
	r, err := challenge.Part2()
	assert.NoError(t, err)
	assert.Equal(t, "1", r)
}

func BenchmarkChallenge_Part2(b *testing.B) {
	buf := bytes.NewReader(fullInput)
	var challenge Challenge
	challenge.Prepare(buf)
	for i := 0; i < b.N; i++ {
		challenge.Part2()
	}
}
