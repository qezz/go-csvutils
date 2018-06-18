package csvutils

import (
	"testing"

	"github.com/qezz/csvutils/csvparseutils"
)

func TestGetSummary(t *testing.T) {
	testcases := []struct {
		name string
		in   *RawCsv
		cfg  *SummaryConfig
		out  *CsvSummary
	}{
		{
			name: "GroupBy 'a' and SumBy 'b'",
			in: &RawCsv{
				records: [][]string{
					{"a", "b", "c"},
					{"z", "2", "3"},
					{"z", "2", "3"},
					{"y", "5", "6"},
					{"y", "5", "6"},
				},
			},
			cfg: NewSummaryConfig("a", []string{"b"}),
			out: &CsvSummary{
				headers: []string{"a", "b"},
				inner: &map[string]([]float64){
					"z": []float64{4.0},
					"y": []float64{10.0},
				},
			},
		},
		{
			name: "GroupBy 'Time Interval' and SumBy 'Count'",
			in: &RawCsv{
				records: [][]string{
					{"Time Interval", "Count"},
					{"2018-06-01T00:00:00Z/2018-06-01T01:00:00Z", "1"},
					{"2018-06-01T00:00:00Z/2018-06-01T01:00:00Z", "2"},
					{"2018-06-01T00:00:00Z/2018-06-01T01:00:00Z", "3"},
					{"2018-06-01T00:00:00Z/2018-06-01T01:00:00Z", "4"},
					{"2018-06-01T01:00:00Z/2018-06-01T02:00:00Z", "10"},
					{"2018-06-01T01:00:00Z/2018-06-01T02:00:00Z", "20"},
					{"2018-06-01T01:00:00Z/2018-06-01T02:00:00Z", "30"},
					{"2018-06-01T01:00:00Z/2018-06-01T02:00:00Z", "40"},
				},
			},
			cfg: NewSummaryConfig("Time Interval", []string{"Count"}),
			out: &CsvSummary{
				headers: []string{"Time Interval", "Count"},
				inner: &map[string]([]float64){
					"2018-06-01T00:00:00Z/2018-06-01T01:00:00Z": []float64{10},
					"2018-06-01T01:00:00Z/2018-06-01T02:00:00Z": []float64{100},
				},
			},
		},
		{
			name: "GroupBy 'Product' and SumBy 'Count'",
			in: &RawCsv{
				records: [][]string{
					{"Product", "Count"},
					{"Apple", "1"},
					{"Banana", "2"},
					{"Cherry", "3"},
					{"Banana", "4"},
					{"Apple", "10"},
					{"Apple", "20"},
					{"Cherry", "30"},
					{"Peaches", "40"},
				},
			},
			cfg: NewSummaryConfig("Product", []string{"Count"}),
			out: &CsvSummary{
				headers: []string{"Product", "Count"},
				inner: &map[string]([]float64){
					"Peaches": []float64{40},
					"Apple":   []float64{31},
					"Cherry":  []float64{33},
					"Banana":  []float64{6},
				},
			},
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			scfg := testcase.cfg
			summary, err := testcase.in.GetSummary(scfg)
			if err != nil {
				t.Fatal(err)
			}

			if !csvparseutils.AreStringSlicesEqual(summary.Headers(), testcase.out.Headers()) {
				t.Fatal("Header differs:", summary.Headers(), "!=", testcase.out.Headers())
			}
			if !csvparseutils.AreMapsEqual(*summary.Inner(), *testcase.out.Inner()) {
				t.Fatal("Summary differs:", *summary.Inner(), "!=", *testcase.out.Inner())
			}
		})
	}
}
