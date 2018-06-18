package errors

import (
	"fmt"
)

// IndexError represents the situation when index is not found
type IndexError struct {
	Idx int
}

// Error is the required method for the error interface
func (err IndexError) Error() string {
	return fmt.Sprintf("index %v is out of range", err.Idx)
}

// NewIndexError returns new IndexError
func NewIndexError(index int) IndexError {
	return IndexError{
		Idx: index,
	}
}
