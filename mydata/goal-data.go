package mydata

// GoalData is the data structure for the goal frame.
type GoalData struct {
	Start   int    `json:"start"`
	Current int    `json:"current"`
	End     int    `json:"end"`
	Unit    string `json:"unit"`
}
