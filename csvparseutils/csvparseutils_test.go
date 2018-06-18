package csvparseutils

import (
	"testing"
)

func TestStringSliceToFloat64(t *testing.T) {
	testcases := []struct {
		name string
		in   []string
		out  []float64
	}{
		{
			name: "positives",
			in:   []string{"0", "0.5", "1.0", "1.5"},
			out:  []float64{0.0, 0.5, 1.0, 1.5},
		},
		{
			name: "negatives",
			in:   []string{"-0", "-0.5", "-1.0", "-1.5"},
			out:  []float64{-0.0, -0.5, -1.0, -1.5},
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			floats, err := StringSliceToFloat64(testcase.in)
			if err != nil {
				t.Fatal("Error encountered: ", err)
			}

			for i, item := range floats {
				if item != testcase.out[i] {
					t.Fatal("Slices arn't the same: ", testcase, "!=", testcase.out)
				}
			}
		})
	}
}
