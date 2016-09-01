package muninplugin

type Metric struct {
	Def *MetricDefinition
	Val interface{} `munin:".value"`
}

// NewMetric returns a pointer to a Metric
func NewMetric() *Metric {
	return &Metric{
		Def: NewMetricDefinition(),
	}
}
