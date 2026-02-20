package formatters

import "time"

func TimeRuFormat(t time.Time) string {
	return t.Format("02.01.2006 15:04:05")
}
