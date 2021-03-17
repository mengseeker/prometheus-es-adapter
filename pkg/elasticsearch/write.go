package elasticsearch

import (
	"github.com/prometheus/common/model"
	"github.com/prometheus/prometheus/prompb"
)

type prometheusSample struct {
	Labels    model.Metric `json:"label"`
	Value     float64      `json:"value"`
	Timestamp int64        `json:"timestamp"`
}

// WriteService will proxy Prometheus write requests to Elasticsearch
type WriteService struct {
}

// NewWriteService creates and returns a new elasticsearch WriteService
func NewWriteService() (*WriteService, error) {
	svc := &WriteService{}
	return svc, nil
}

// Close will close the underlying elasticsearch BulkProcessor
func (svc *WriteService) Close() error {
	return nil
}

// Write will enqueue Prometheus sample data to be batch written to Elasticsearch
func (svc *WriteService) Write(req []*prompb.TimeSeries) {

}
