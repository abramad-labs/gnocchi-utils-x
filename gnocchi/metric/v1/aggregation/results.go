package aggregation

import (
	"github.com/abramad-labs/gophercloud-utils-x/gnocchi/metric/v1/measures"
	"github.com/gophercloud/gophercloud"
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
