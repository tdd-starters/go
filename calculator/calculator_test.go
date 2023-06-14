package calculator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var expressionTests = []struct {
	name     string
	input    string
	expected int64
}{
	{"Single digit", "0", 0},
	// {"Add", "1 + 2", 3},
	// ... more tests here
	// {"Complex", "(1 + 2) * 5 / 2 + -3 * 7", -14},
}

func TestExpressions(t *testing.T) {
	for _, data := range expressionTests {
		t.Run(data.name, func(t *testing.T) {
			// Given an input expression
			expression := data.input

			// when the string is evaluated
			result := Evaluate(expression)

			// it should be an empty string
			assert.Equal(t, data.expected, result)
		})
	}
}

/*
var errorTests = []struct {
	name     string
	input    string
}{
	{"string", "foo"},
	// ... more tests here
}
func TestError(t *testing.T) {
	for _, data := range expressionTests {
		t.Run(data.name, func(t *testing.T) {
			// Given an invalid input expression
			expression := data.input

			// when the string is evaluated
			_, error := Evaluate(expression)

			// then the error should not be nil
			assert.NotEqual(t, nil, error)
		})
	}
}
*/
