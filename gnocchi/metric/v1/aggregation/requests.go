package aggregation

import (
	"net/url"
	"time"

	"github.com/abramad-labs/gophercloud-utils-x/gnocchi"
	"github.com/gophercloud/gophercloud"
)

type AggregateQueryBuilder interface {
	ToAggregateQuery() (string, error)
}

type AggregateBodyBuilder interface {
	ToAggregateBody() (map[string]interface{}, error)
}

type AggregateQueryOpts struct {
	Granularity string `q:"granularity"`
	Start       *time.Time
	Stop        *time.Time
}

type AggregateBodyOpts struct {
	Operation    string `json:"operations"`
	Search       string `json:"search"`
	ResourceType string `json:"resource_type"`
}

func (opts AggregateBodyOpts) ToAggregateBody() (map[string]interface{}, error) {
	b, err := gophercloud.BuildRequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	return b, nil
}

func (opts AggregateQueryOpts) ToAggregateQuery() (string, error) {
	q, err := gophercloud.BuildQueryString(opts)
	params := q.Query()

	if opts.Start != nil {
		params.Add("start", opts.Start.Format(gnocchi.RFC3339NanoNoTimezone))
	}

	if opts.Stop != nil {
		params.Add("stop", opts.Stop.Format(gnocchi.RFC3339NanoNoTimezone))
	}

	q = &url.URL{RawQuery: params.Encode()}
	return q.String(), err
}

func Aggregate(c *gophercloud.ServiceClient, bodyOpts AggregateBodyBuilder, queryOpts AggregateQueryBuilder) (r GetResult) {
	url := aggregateURL(c)
	if queryOpts != nil {
		query, err := queryOpts.ToAggregateQuery()
		if err != nil {
			r.Err = err
			return
		}
		url += query
	}

	b, err := bodyOpts.ToAggregateBody()
	if err != nil {
		r.Err = err
		return
	}

	_, r.Err = c.Post(url, b, &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200},
		MoreHeaders: map[string]string{
			"Accept": "application/json, */*",
		},
	})
	return
}
