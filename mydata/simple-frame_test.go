package mydata

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSimpleFrame(t *testing.T) {
	frame := SimpleFrame("Hello World", 1, 2)

	assert.Equal(t, "Hello World", *frame.Text)
	assert.Equal(t, 1, frame.Icon)
	assert.Equal(t, 2, frame.Duration)
}

func TestSimpleFrameJSONEncoding(t *testing.T) {
	frame := SimpleFrame("Hello World", 42, 2)

	encoded, err := json.Marshal(frame)
	assert.NoError(t, err)

	assert.Equal(t, "{\"text\":\"Hello World\",\"icon\":42,\"duration\":2}", string(encoded))
}

func TestSimpleFrameJSONEncodingNoIcon(t *testing.T) {
	frame := SimpleFrame("Hello World", 0, 2)

	encoded, err := json.Marshal(frame)
	assert.NoError(t, err)

	assert.Equal(t, "{\"text\":\"Hello World\",\"duration\":2}", string(encoded))
}

func TestSimpleFrameJSONEncodingNoDuration(t *testing.T) {
	frame := SimpleFrame("Hello World", 42, 0)

	encoded, err := json.Marshal(frame)
	assert.NoError(t, err)

	assert.Equal(t, "{\"text\":\"Hello World\",\"icon\":42}", string(encoded))
}
