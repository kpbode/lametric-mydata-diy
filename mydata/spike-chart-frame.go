package mydata

// SpikeChartFrame creates a frame with the given icon, duration and chart data.
func SpikeChartFrame(icon int, duration int, chartData []int) Frame {
	return Frame{
		Icon:      icon,
		Duration:  duration,
		ChartData: chartData,
	}
}
