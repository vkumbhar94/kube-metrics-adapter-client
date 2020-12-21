package collector

import (
	"fmt"
	"time"

	autoscalingv2 "k8s.io/api/autoscaling/v2beta2"
	"k8s.io/metrics/pkg/apis/custom_metrics"
	"k8s.io/metrics/pkg/apis/external_metrics"
)

type CollectorPlugin interface {
	// NewCollector creates new collector
	NewCollector(hpa *autoscalingv2.HorizontalPodAutoscaler, config *MetricConfig, interval time.Duration) (Collector, error)
}
type MetricConfig struct {
	MetricTypeName
	CollectorType   string
	Config          map[string]string
	ObjectReference custom_metrics.ObjectReference
	PerReplica      bool
	Interval        time.Duration
	MetricSpec      autoscalingv2.MetricSpec
}
type Collector interface {
	GetMetrics() ([]CollectedMetric, error)
	Interval() time.Duration
}
type CollectedMetric struct {
	Type     autoscalingv2.MetricSourceType
	Custom   custom_metrics.MetricValue
	External external_metrics.ExternalMetricValue
}
type MetricTypeName struct {
	Type   autoscalingv2.MetricSourceType
	Metric autoscalingv2.MetricIdentifier
}

type PluginNotFoundError struct {
	MetricTypeName MetricTypeName
}

func (p *PluginNotFoundError) Error() string {
	return fmt.Sprintf("no plugin found for %s", p.MetricTypeName)
}
