package aggregation

import (
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/utils/gnocchi/metric/v1/measures"
)

type GetResult struct {
	commonResult
}

type commonResult struct {
	gophercloud.Result
}

type AggregatedResult struct {
	Measures struct {
		Aggregated [][]interface{} `json:"aggregated"`
	} `json:"measures"`
}

func (r commonResult) Extract() (*measures.Measure, error) {
	var s *measures.Measure
	err := r.ExtractInto(&s)
	return s, err
}
