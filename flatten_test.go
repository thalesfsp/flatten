package flatten

import (
	"fmt"
	"testing"
)

func Test_flatten(t *testing.T) {
	// Test case for compounded slice: [1, 2, [3, 4, [5, 6]]]
	s1 := []int{1, 2}
	s2 := []int{3, 4}
	s3 := []int{5, 6}
	sMatrix := []interface{}{
		[]interface{}{
			s1,
			[]interface{}{
				s2,
				[]interface{}{
					s3,
				},
			},
		},
	}

	sResult, err := Flatten(sMatrix)

	if err != nil {
		t.Errorf(
			"Fail. Expected '%s', got '%s', reason: %s", "[1 2 3 4 5 6]",
			sResult,
			err.Error(),
		)
	} else {
		fmt.Println("Slice result:", sResult)
	}

	// Test case for multidimensional array (2D):
	// | 1 2 |
	// | 3 4 |
	var twoD [2][2]int
	twoD[0][0] = 1
	twoD[0][1] = 2
	twoD[1][0] = 3
	twoD[1][1] = 4

	aResult, err := Flatten(twoD)

	if err != nil {
		t.Errorf(
			"Fail. Expected '%s', got '%s', reason: %s", "[1 2 3 4]",
			aResult,
			err.Error(),
		)
	} else {
		fmt.Println("Array result:", aResult)
	}

	// Test case for empty argument
	if _, err := Flatten(); err == nil {
		t.Errorf(
			"Fail. Expected error: %s",
			err.Error(),
		)
	}
}

func BenchmarkFlattenSlice(b *testing.B) {
	for n := 0; n < b.N; n++ {
		// Slice: [1, 2, [3, 4, [5, 6]]]
		s1 := []int{1, 2}
		s2 := []int{3, 4}
		s3 := []int{5, 6}
		sMatrix := []interface{}{
			[]interface{}{
				s1,
				[]interface{}{
					s2,
					[]interface{}{
						s3,
					},
				},
			},
		}

		Flatten(sMatrix)
	}
}

func BenchmarkFlattenArray(b *testing.B) {
	// Multidimensional array (2D):
	// | 1 2 |
	// | 3 4 |
	var twoD [2][2]int
	twoD[0][0] = 1
	twoD[0][1] = 2
	twoD[1][0] = 3
	twoD[1][1] = 4

	Flatten(twoD)
}

func BenchmarkFlattenEmpty(b *testing.B) {
	Flatten()
}

func Example() {
	// The following code demonstrates how to flatten a 4D array and assumes
	// that the flatten package has already been imported.

	// Multidimensional array (4D)
	const rolls = 4
	const columns = 4
	var fourD [rolls][columns]int

	// Generator/filler
	for roll := 0; roll < rolls; roll++ {
		for column := 0; column < columns; column++ {
			fourD[roll][column] = (roll + 1) + (column * 2)
		}
	}

	result, err := Flatten(fourD)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)

	// Output:
	// [1 3 5 7 2 4 6 8 3 5 7 9 4 6 8 10]
}
