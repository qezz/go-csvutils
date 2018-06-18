package csvparseutils

import (
	"strconv"
)

// Convert slice of strings to slice of float64s
func StringSliceToFloat64(ss []string) ([]float64, error) {
	result := make([]float64, len(ss))
	for i, item := range ss {
		value, err := strconv.ParseFloat(item, 64)
		if err != nil {
			return nil, err
		}
		result[i] = value
	}

	return result, nil
}

// Compare two maps, return true if they are equal, false otherwise
func AreMapsEqual(a, b map[string]([]float64)) bool {
	if len(a) != len(b) {
		return false
	}

	for k, v := range a {
		if w, ok := b[k]; !ok || !AreFloat64SlicesEqual(v, w) {
			return false
		}
	}

	return true
}

// AreStringSlicesEqual compares two slices, return true if they are equal, false otherwise
func AreStringSlicesEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	for i, item := range a {
		if b[i] != item {
			return false
		}
	}

	return true
}

// AreFloat64SlicesEqual compares two slices, return true if they are equal, false otherwise
func AreFloat64SlicesEqual(a, b []float64) bool {
	if len(a) != len(b) {
		return false
	}

	for i, item := range a {
		if b[i] != item {
			return false
		}
	}

	return true
}
