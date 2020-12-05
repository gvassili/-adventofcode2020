package day03

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

const simpleInput = `#.#
.#.`

const input = `..##.......
#...#...#..
.#....#..#.
..#.#...#.#
.#...##..#.
..#.##.....
.#.#.#....#
.#........#
#.##...#...
#...##....#
.#..#...#.#
`

func TestDay01_Prepare(t *testing.T) {
	var challenge Challenge
	err := challenge.Prepare(bytes.NewReader([]byte(simpleInput)))
	assert.NoError(t, err)
	assert.Equal(t, []byte("#.#.#."), challenge.slope)
	assert.Equal(t, 2, challenge.height)
	assert.Equal(t, 3, challenge.width)
}

func BenchmarkDay01_Prepare(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		buf := bytes.NewReader([]byte(input))
		b.StartTimer()
		var challenge Challenge
		challenge.Prepare(buf)
	}
}

func TestDay01_Part1(t *testing.T) {
	var challenge Challenge
	err := challenge.Prepare(bytes.NewReader([]byte(input)))
	assert.NoError(t, err)
	r, err := challenge.Part1()
	assert.NoError(t, err)
	assert.Equal(t, "7", r)
}

func BenchmarkDay01_Part1(b *testing.B) {
	buf := bytes.NewReader([]byte(input))
	var challenge Challenge
	challenge.Prepare(buf)
	for i := 0; i < b.N; i++ {
		challenge.Part1()
	}
}

func TestDay01_Part2(t *testing.T) {
	var challenge Challenge
	err := challenge.Prepare(bytes.NewReader([]byte(input)))
	assert.NoError(t, err)
	r, err := challenge.Part2()
	assert.NoError(t, err)
	assert.Equal(t, "336", r)
}

func BenchmarkDay01_Part2(b *testing.B) {
	buf := bytes.NewReader([]byte(input))
	var challenge Challenge
	challenge.Prepare(buf)
	for i := 0; i < b.N; i++ {
		challenge.Part2()
	}
}
