package aggregation

import (
	"github.com/gophercloud/gophercloud"
)

const aggregatePath = "aggregates"

func aggregateURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL(aggregatePath)
}
