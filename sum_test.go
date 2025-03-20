package test_main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	t.Run("2 plus 5, return 7", func(t *testing.T){
		var result int = Sum(2, 5)
		assert.Equal(t, 4, result)
	})

	t.Run("3 plus 2, return 5", func(t *testing.T){
		var result int = Sum(3, 2)
		assert.Equal(t, 5, result)
	})

}

func TestSumWithTableDriven(t *testing.T) {
	var testCases []struct{
		name string
		a int
		b int
		expected int
	} = []struct{
		name string
		a int
		b int
		expected int
	}{
		{
			name: "3 and 5, return 8",
			a: 3,
			b: 5,
			expected: 8,
		},
		{
			name: "1 and 5, return 6",
			a: 1,
			b: 5,
			expected: 6,
		},
	}

	for _, testCase := range testCases{
		t.Run(testCase.name, func(t *testing.T){
			actual := Sum(testCase.a, testCase.b)
			assert.Equal(t, testCase.expected, actual)
		})
	}
}