package metrics

import (
	"fmt"
)

type metric struct {
	name   string
	labels map[string]interface{}
}

type Option func(*metric)

func Category(v string) Option {
	return func(m *metric) {
		m.labels["category"] = v
	}
}

func Latency(v int) Option {
	return func(m *metric) {
		m.labels["latency"] = v
	}
}

func Tags(v []string) Option {
	return func(m *metric) {
		m.labels["tags"] = v
	}
}

// func Context(v ctx) Option {
// 	return func(m *metric) {
// 		// extract
// 		// if v, ok := ctx.Value("category").(string); ok {
// 		// 	category = v
// 		// }
// 		//m.labels["tags"] = v
// 	}
// }

// Count a metric
func Count(name string, opts ...Option) {
	m := &metric{
		name:   name,
		labels: map[string]interface{}{},
	}

	for _, opt := range opts {
		opt(m)
	}

	m.print()
}

// print
func (m metric) print() {
	out := m.name + ":"
	for k, v := range m.labels {
		out += fmt.Sprintf(" %s=%v", k, v)
	}
	fmt.Println(out)
}
