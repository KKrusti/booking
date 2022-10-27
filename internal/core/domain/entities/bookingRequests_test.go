package entities

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {

	request := []Request{
		{
			Id: "A",
		},
		{
			Id: "B",
		},
		{
			Id: "C",
		},
	}

	allCombinations := combinations(request)

	assert.Equal(t, 7, len(allCombinations))

	a := []Request{{Id: "A"}}
	b := []Request{{Id: "B"}}
	c := []Request{{Id: "C"}}
	ab := []Request{{Id: "A"}, {Id: "B"}}
	bc := []Request{{Id: "B"}, {Id: "C"}}
	ac := []Request{{Id: "A"}, {Id: "C"}}
	abc := []Request{{Id: "A"}, {Id: "B"}, {Id: "C"}}
	assert.Contains(t, allCombinations, a)
	assert.Contains(t, allCombinations, b)
	assert.Contains(t, allCombinations, c)
	assert.Contains(t, allCombinations, ab)
	assert.Contains(t, allCombinations, ac)
	assert.Contains(t, allCombinations, bc)
	assert.Contains(t, allCombinations, abc)

}
