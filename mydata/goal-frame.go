package mydata

// GoalFrame creates a frame with the given icon, duration and goal data.
func GoalFrame(icon int, duration int, goalData GoalData) Frame {
	return Frame{
		Icon:     icon,
		Duration: duration,
		GoalData: &goalData,
	}
}
