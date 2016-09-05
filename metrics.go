package muninplugin

import (
	"fmt"
	"reflect"
	"strings"
)

type Metrics map[string]*Metric

// NewMetrics returns a non-nil Metrics
func NewMetrics() Metrics {
	return make(Metrics)
}

// Values returns the output for use by a regular Munin data
// collection run.
func (ms Metrics) Values() string {
	var result []string
	for k, v := range ms {
		switch sw := v.Val.(type) {
		case float32, float64:
			result = append(result, fmt.Sprintf("%s.value %.2f\n", k, sw))
		case int8, uint8, int32, uint32, int64, uint64, int, uint:
			result = append(result, fmt.Sprintf("%s.value %d\n", k, sw))
		default:
			result = append(result, fmt.Sprintf("%s.value U", k))
		}
	}
	return strings.Join(result, "")
}

// Config returns the output for use by an inital Munin run
// to collect configuration data used by subsequent value runs.
func (ms Metrics) Config() string {
	var result []string
	for k, v := range ms {
		// The load label is the string used as a key in the Metrics map.
		result = append(result, fmt.Sprintf("%s.label %s\n", k, k))

		val := reflect.ValueOf(*v.Def)
		for i := 0; i < val.NumField(); i++ {
			value := val.Field(i)
			kind := value.Kind()
			fieldType := val.Type().Field(i)
			tags := fieldType.Tag
			muninTag := tags.Get("munin")
			switch kind {
			case reflect.Float64, reflect.Float32:
				result = append(result,
					fmt.Sprintf("%s%s %.2f\n", k, muninTag, value.Float()))
			case reflect.String:
				if value.String() != "" {
					result = append(result,
						fmt.Sprintf("%s%s %s\n", k, muninTag, value.String()))
				}
			case reflect.Int, reflect.Int64, reflect.Int32, reflect.Int8:
				result = append(result,
					fmt.Sprintf("%s%s %d\n", k, muninTag, value.Int()))
			case reflect.Bool:
				result = append(result, fmt.Sprintf("%s%s %s\n", k, muninTag, toYN(value.Bool())))
			case reflect.Interface:
				switch value.Interface().(type) {
				case int, int8, int32, int64:
					result = append(result,
						fmt.Sprintf("%s%s %d\n", k, muninTag, value.Interface()))
				case float32, float64:
					result = append(result,
						fmt.Sprintf("%s%s %.2f\n", k, muninTag, value.Interface()))
				}

			}
		}
	}
	return strings.Join(result, "")
}
