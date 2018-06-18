package csvutils

import (
	"strings"
	"testing"
)

func TestRawCsvFromReader(t *testing.T) {
	rawcsv := &RawCsv{
		records: [][]string{
			{"a", "b", "c"},
			{"1", "2", "3"},
			{"4", "5", "6"},
		},
	}

	csvString := `a,b,c
1,2,3
4,5,6
`

	strReader := strings.NewReader(csvString)
	csv, err := RawCsvFromReader(strReader)
	if err != nil {
		t.Fatalf("Creating RawCsv failed:, %v\n", err)
	}

	if !equalRawCsvs(*csv, *rawcsv) {
		t.Fatal("FromReader produced RawCsv that differs:", csv, "!=", rawcsv)
	}
}

func TestColumnIndexByName(t *testing.T) {
	colsNames := []struct {
		name string
		in   string
		out  int // TODO: Convert out type to { idx, error }
	}{
		{
			name: "Get index by column name 'a'",
			in:   "a",
			out:  0,
		},
		{
			name: "Get index by column name 'b'",
			in:   "b",
			out:  1,
		},
	}

	rawcsv := &RawCsv{
		records: [][]string{{"a", "b"}, {"1", "2"}},
	}

	for _, cn := range colsNames {
		t.Run(cn.name, func(t *testing.T) {
			idx, err := rawcsv.ColumnIndexByName(cn.in)
			if err != nil {
				t.Fatal("Error encountered: ", err)
			}
			if idx != cn.out {
				t.Fatal("Found index isn't correct: ", idx, "!=", cn.out)
			}
		})
	}
}

func equalRawCsvs(a, b RawCsv) bool {
	if len(a.Rows()) != len(b.Rows()) {
		return false
	}
	for i, row := range a.Rows() {
		for j, value := range row {
			if value != b.Rows()[i][j] {
				return false
			}
		}
	}

	return true
}

func TestFilterByIndices(t *testing.T) {
	testcases := []struct {
		name string
		in   []int
		out  *RawCsv
	}{
		{
			name: "Filter by column with index 0",
			in:   []int{0},
			out: &RawCsv{
				records: [][]string{{"a"}, {"1"}},
			},
		},
		{
			name: "Filter by column with index 1",
			in:   []int{1},
			out: &RawCsv{
				records: [][]string{{"b"}, {"2"}},
			},
		},
		{
			name: "Filter by columns with indices 0 and 1",
			in:   []int{0, 1},
			out: &RawCsv{
				records: [][]string{{"a", "b"}, {"1", "2"}},
			},
		},
	}

	rawcsv := &RawCsv{
		records: [][]string{{"a", "b"}, {"1", "2"}},
	}

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			csv, err := rawcsv.FilterByIndices(testcase.in)
			if err != nil {
				t.Fatal("Error encountered: ", err)
			}
			if !equalRawCsvs(*csv, *testcase.out) {
				t.Fatal("Filter isn't correct: ", testcase, "!=", testcase.out)
			}
		})
	}
}

func TestFilterByNames(t *testing.T) {
	testcases := []struct {
		name string
		in   []string
		out  *RawCsv
	}{
		{
			name: "Filter by column with name 'a'",
			in:   []string{"a"},
			out: &RawCsv{
				records: [][]string{{"a"}, {"1"}},
			},
		},
		{
			name: "Filter by column with name 'b'",
			in:   []string{"b"},
			out: &RawCsv{
				records: [][]string{{"b"}, {"2"}},
			},
		},
		{
			name: "Filter by columns 'a' and 'b'",
			in:   []string{"a", "b"},
			out: &RawCsv{
				records: [][]string{{"a", "b"}, {"1", "2"}},
			},
		},
	}

	rawcsv := &RawCsv{
		records: [][]string{{"a", "b"}, {"1", "2"}},
	}

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			csv, err := rawcsv.FilterByNames(testcase.in)
			if err != nil {
				t.Fatal("Error encountered: ", err)
			}
			if !equalRawCsvs(*csv, *testcase.out) {
				t.Fatal("Filter isn't correct: ", testcase, "!=", testcase.out)
			}
		})
	}
}
