package day01

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

const input = `1721
979
366
299
675
1456`

func TestChallenge_Prepare(t *testing.T) {
	var challenge Challenge
	err := challenge.Prepare(bytes.NewReader([]byte(input)))
	assert.NoError(t, err)
	assert.Equal(t, []int{1721, 979, 366, 299, 675, 1456}, challenge.input)
}

func BenchmarkChallenge_Prepare(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		buf := bytes.NewReader([]byte(input))
		b.StartTimer()
		var challenge Challenge
		challenge.Prepare(buf)
	}
}

func TestChallenge_Part1(t *testing.T) {
	var challenge Challenge
	err := challenge.Prepare(bytes.NewReader([]byte(input)))
	assert.NoError(t, err)
	r, err := challenge.Part1()
	assert.NoError(t, err)
	assert.Equal(t, "514579", r)
}

func BenchmarkChallenge_Part1(b *testing.B) {
	buf := bytes.NewReader([]byte(input))
	var challenge Challenge
	challenge.Prepare(buf)
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
	assert.Equal(t, "241861950", r)
}

func BenchmarkChallenge_Part2(b *testing.B) {
	buf := bytes.NewReader([]byte(input))
	var challenge Challenge
	challenge.Prepare(buf)
	for i := 0; i < b.N; i++ {
		challenge.Part2()
	}
}
