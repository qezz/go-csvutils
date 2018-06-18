package csvutils

// SummaryConfig represents the configuration to
type SummaryConfig struct {
	// Header to group by
	GroupBy string

	// Header to sum the values
	// Values should be convertable to float64
	SumBy []string
}

func NewSummaryConfig(groupBy string, sumBy []string) *SummaryConfig {
	return &SummaryConfig{
		GroupBy: groupBy,
		SumBy:   sumBy,
	}
}

// CsvSummary represents the summary
type CsvSummary struct {
	headers []string
	inner   *map[string]([]float64)
}

// NewSummaryConfig returns the new CsvSummary
func NewCsvSummary(h []string, m *map[string]([]float64)) *CsvSummary {
	return &CsvSummary{
		headers: h,
		inner:   m,
	}
}

// Inner returns the key-value map of the CsvSummary
func (summary CsvSummary) Inner() *map[string]([]float64) {
	return summary.inner
}

// Inner returns the key-value map of the CsvSummary
func (summary CsvSummary) Headers() []string {
	return summary.headers
}
