package elasticsearch

import (
	"github.com/prometheus/client_golang/prometheus"
)

const (
	namespace = "es_adapter"
)

var (
	flushedDesc = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "flushed_total"),
		"Number of times the flush interval has been invoked",
		nil,
		nil,
	)
	committedDesc = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "committed_total"),
		"Number of times workers committed bulk requests",
		nil,
		nil,
	)
	indexedDesc = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "indexed_total"),
		"Number of requests indexed",
		nil,
		nil,
	)
	createdDesc = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "created_total"),
		"Number of requests that ES reported as creates (201)",
		nil,
		nil,
	)
	updatedDesc = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "updated_total"),
		"Number of requests that ES reported as updates",
		nil,
		nil,
	)
	deletedDesc = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "deleted_total"),
		"Number of requests that ES reported as deletes",
		nil,
		nil,
	)
	succeededDesc = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "succeeded_total"),
		"Number of requests that ES reported as successful",
		nil,
		nil,
	)
	failedDesc = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "failed_total"),
		"Number of requests that ES reported as failed",
		nil,
		nil,
	)
	queuedDesc = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "queued_total"),
		"Number of requests queued per worker",
		[]string{"worker"},
		nil,
	)
	durationDesc = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "duration"),
		"Duration of last commit per worker",
		[]string{"worker"},
		nil,
	)
)

// Describe describes all the metrics exported by the memcached exporter. It
// implements prometheus.Collector.
func (svc *WriteService) Describe(ch chan<- *prometheus.Desc) {
	ch <- flushedDesc
	ch <- committedDesc
	ch <- indexedDesc
	ch <- createdDesc
	ch <- updatedDesc
	ch <- deletedDesc
	ch <- succeededDesc
	ch <- failedDesc
	ch <- queuedDesc
	ch <- durationDesc
}

// Collect fetches the statistics from the elasticsearch bulk processor, and
// delivers them as Prometheus metrics. It implements prometheus.Collector.
func (svc *WriteService) Collect(ch chan<- prometheus.Metric) {

}
