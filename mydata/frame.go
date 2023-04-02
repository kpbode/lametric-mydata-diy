package mydata

// Frame Wraps data according to https://help.lametric.com/support/solutions/articles/6000225467-my-data-diy
type Frame struct {
	Text      *string   `json:"text,omitempty"`
	Icon      int       `json:"icon,omitempty"`
	Duration  int       `json:"duration,omitempty"`
	GoalData  *GoalData `json:"goalData,omitempty"`
	ChartData []int     `json:"chartData,omitempty"`
}
