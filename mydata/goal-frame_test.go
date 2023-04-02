package mydata

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGoalFrame(t *testing.T) {
	frame := GoalFrame(1, 2, GoalData{Start: 1, End: 2, Current: 3, Unit: "Unit"})

	assert.Equal(t, 1, frame.Icon)
	assert.Equal(t, 2, frame.Duration)
	assert.Equal(t, 1, frame.GoalData.Start)
	assert.Equal(t, 2, frame.GoalData.End)
	assert.Equal(t, 3, frame.GoalData.Current)
	assert.Equal(t, "Unit", frame.GoalData.Unit)
}

func TestGoalFrameJSONEncoding(t *testing.T) {
	frame := GoalFrame(42, 2, GoalData{Start: 1, End: 2, Current: 3, Unit: "Unit"})

	encoded, err := json.Marshal(frame)
	assert.NoError(t, err)

	assert.Equal(t, "{\"icon\":42,\"duration\":2,\"goalData\":{\"start\":1,\"current\":3,\"end\":2,\"unit\":\"Unit\"}}", string(encoded))
}
