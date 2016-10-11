package time

import tt "time"

// UTC is a Source that returns the current time in UTC
var UTC Source = func() tt.Time {
	return tt.Now().UTC()
}
