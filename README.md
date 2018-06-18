# Several CSV utils

## Types

### RawCsv

RawCsv type is just `[][]string` inside, but introduces some useful methods:

* `ColumnIndexByName` to get the column index by its header name
* `FilterByIndeces` to produce the filtered RawCsv by provided column indices 
* `FilterByNames` to produce the filtered RawCsv by provided headers (uses `ColumnIndexByName` and `FilterByIndeces` inside)
* `GetSummary` is more complicated, see below.

You can easily create RawCsv with:

```go
// similar to "encoding/csv", but read everything
rawcsvdata := csvutils.RawCsvFromReader(/*typical io Reader*/) 

// from [][]string
records:= [][]string{
	{"a", "b", "c"},
	{"1", "2", "3"},
	{"4", "5", "6"},
},
rawcsvdata := csvutils.RawCsvFromRows(records)
```

### Summary

`GetSummary` returns the summary (`CsvSummary`) by the provided `SummaryConfig` 

The key fields of the `SummaryConfig` are:

* `GroupBy string` is the column name to group by
* `SumBy []string` are the columns names to sum

For example, if you have CSV with the following data:

| "NameColumn" | "ValueColumn" |
| :--- | :--- |
| "B" | "1" |
| "D" | "2" |
| "C" | "3" |
| "D" | "4" |
| "B" | "10" |
| "B" | "20" |
| "C" | "30" |
| "A" | "40" |


It can be summarized with

```go
cfg := NewSummaryConfig("NameColumn", []string{"ValueColumn"})
rawcsvdata.GetSummary(cfg)
```

and it produces the following mapping (actually `map[string]([]float64)`):

| | |
| :--- | :--- |
| "A" | 40 |
| "B" | 31 |
| "C" | 33 |
| "D" | 6 |

See `summary_test.go` for more examples

## Future improvements

* Streaming API
* Custom error types

## License 

MIT or Beer-ware at your choice

## Author 

Sergey Mishin, <sergei.a.mishin@gmail.com>
