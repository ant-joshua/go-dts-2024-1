package helper

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// func TestFailSum(t *testing.T) {
// 	result := Sum(5, 5)

// 	if result != 11 {
// 		// t.FailNow()
// 		t.Fatalf("Sum was incorrect, got: %d, want: %d.", result, 11)

// 	}

// 	fmt.Println("This is the end of the test fail sum")
// }

func TestSum(t *testing.T) {
	result := Sum(5, 5)

	if result != 10 {
		t.FailNow()
		// t.Errorf("Sum was incorrect, got: %d, want: %d.", result, 10)
	}

	fmt.Println("This is the end of the test sum")
}

func TestFailSumRequire(t *testing.T) {
	result := Sum(5, 5)
	require.Equal(t, 11, result, "Sum was incorrect")

	fmt.Println("This is the end of the test fail sum require")
}

func TestSumAssert(t *testing.T) {
	result := Sum(5, 5)

	assert.Equal(t, 10, result, "Sum was incorrect")

	fmt.Println("This is the end of the test sum assert")
}

func TestTableSum(t *testing.T) {
	tests := []struct {
		requested    int
		expected     int
		errorMessage string
	}{
		{Sum(10, 5), 15, "Sum was incorrect"},
		{Sum(10, 2), 13, "Sum was incorrect"},
		{Sum(10, 7), 16, "Sum was incorrect"},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("Test %d", test.requested), func(t *testing.T) {
			require.Equal(t, test.expected, test.requested, test.errorMessage)
		})
	}
}
