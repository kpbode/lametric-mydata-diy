package mydata

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSpikeChartFrame(t *testing.T) {
	frame := SpikeChartFrame(1, 2, []int{3, 4, 5})

	assert.Equal(t, 1, frame.Icon)
	assert.Equal(t, 2, frame.Duration)
	assert.Equal(t, []int{3, 4, 5}, frame.ChartData)
}

func TestSpikeChartFrameJSONEncoding(t *testing.T) {
	frame := SpikeChartFrame(42, 2, []int{3, 4, 5})

	encoded, err := json.Marshal(frame)
	assert.NoError(t, err)

	assert.Equal(t, "{\"icon\":42,\"duration\":2,\"chartData\":[3,4,5]}", string(encoded))
}
