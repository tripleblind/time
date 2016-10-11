package time

import tt "time"

// Source defines a source for the current time
type Source func() tt.Time
