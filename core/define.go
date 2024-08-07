package core

import (
	"context"
	"net/http"

	dto "github.com/prometheus/client_model/go"
)

type Fixture interface {
	Setup(ctx context.Context) error
	TearDown(ctx context.Context) error
}

type Exporter interface {
	Start(ctx context.Context) (string, error)
}

type MetricFamiliesChecker interface {
	String() string
	Check(metricFamily map[string]*dto.MetricFamily) (bool, string)
}

type HTTPClient interface {
	Get(string) (*http.Response, error)
}
