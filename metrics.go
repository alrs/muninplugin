package muninplugin

import (
	"fmt"
	"reflect"
	"strings"
)

type Metrics map[string]*Metric

func NewMetrics() Metrics {
	return make(Metrics)
}

func (ms Metrics) Values() string {
	var result []string
	for k, v := range ms {
		result = append(result, fmt.Sprintf("%s.value %.2f\n", k, v.Val))
	}
	return strings.Join(result, "")
}

func (ms Metrics) Config() string {
	var result []string
	for k, v := range ms {
		result = append(result, fmt.Sprintf("%s.label %s\n", k, k))
		val := reflect.ValueOf(*v.Def)
		for i := 0; i < val.NumField(); i++ {
			value := val.Field(i)
			kind := value.Kind().String()
			fieldType := val.Type().Field(i)
			tags := fieldType.Tag
			muninTag := tags.Get("munin")
			switch kind {
			case "float32":
				result = append(result,
					fmt.Sprintf("%s%s %.2f\n", k, muninTag, value.Float()))
			case "string":
				if value.String() != "" {
					result = append(result,
						fmt.Sprintf("%s%s %s\n", k, muninTag, value.String()))
				}
			case "int":
				result = append(result,
					fmt.Sprintf("%s%s %d\n", k, muninTag, value.Int()))
			case "bool":
				result = append(result, fmt.Sprintf("%s%s %s\n", k, muninTag, toYN(value.Bool())))
			}
		}
	}
	return strings.Join(result, "")
}
