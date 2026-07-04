package monitor

import (
	"testing"

	"github.com/yeying-community/router/common/config"
)

func TestMetricAutoRecoverAfterSecondsFallsBackToDefault(t *testing.T) {
	previous := config.MetricAutoRecoverAfterSeconds
	t.Cleanup(func() {
		config.MetricAutoRecoverAfterSeconds = previous
	})

	tests := []struct {
		name  string
		value int
		want  int
	}{
		{name: "zero", value: 0, want: 300},
		{name: "negative", value: -1, want: 300},
		{name: "positive", value: 60, want: 60},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			config.MetricAutoRecoverAfterSeconds = tt.value

			if got := metricAutoRecoverAfterSeconds(); got != tt.want {
				t.Fatalf("metricAutoRecoverAfterSeconds() = %d, want %d", got, tt.want)
			}
		})
	}
}
