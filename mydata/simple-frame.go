package mydata

// SimpleFrame creates a frame with the given text, icon and duration.
func SimpleFrame(text string, icon int, duration int) Frame {
	return Frame{
		Text:     &text,
		Icon:     icon,
		Duration: duration,
	}
}
