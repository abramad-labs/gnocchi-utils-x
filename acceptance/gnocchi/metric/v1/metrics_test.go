//go:build acceptance || metric || metrics
// +build acceptance metric metrics

package v1

import (
	"testing"

	"github.com/abramad-labs/gophercloud-utils-x/acceptance/clients"
	"github.com/abramad-labs/gophercloud-utils-x/gnocchi/metric/v1/metrics"
	"github.com/gophercloud/gophercloud/acceptance/tools"
)

func TestMetricsCRUD(t *testing.T) {
	client, err := clients.NewGnocchiV1Client()
	if err != nil {
		t.Fatalf("Unable to create a Gnocchi client: %v", err)
	}

	metric, err := CreateMetric(t, client)
	if err != nil {
		t.Fatalf("Unable to create a Gnocchi metric: %v", err)
	}
	defer DeleteMetric(t, client, metric.ID)

	tools.PrintResource(t, metric)
}

func TestMetricsList(t *testing.T) {
	client, err := clients.NewGnocchiV1Client()
	if err != nil {
		t.Fatalf("Unable to create a Gnocchi client: %v", err)
	}

	metric1, err := CreateMetric(t, client)
	if err != nil {
		t.Fatalf("Unable to create a Gnocchi metric: %v", err)
	}
	defer DeleteMetric(t, client, metric1.ID)

	metric2, err := CreateMetric(t, client)
	if err != nil {
		t.Fatalf("Unable to create a Gnocchi metric: %v", err)
	}
	defer DeleteMetric(t, client, metric2.ID)

	metric3, err := CreateMetric(t, client)
	if err != nil {
		t.Fatalf("Unable to create a Gnocchi metric: %v", err)
	}
	defer DeleteMetric(t, client, metric3.ID)

	listOpts := metrics.ListOpts{}
	allPages, err := metrics.List(client, listOpts).AllPages()
	if err != nil {
		t.Fatalf("Unable to list metrics: %v", err)
	}

	allMetrics, err := metrics.ExtractMetrics(allPages)
	if err != nil {
		t.Fatalf("Unable to extract metrics: %v", err)
	}

	for _, metric := range allMetrics {
		tools.PrintResource(t, metric)
	}
}
