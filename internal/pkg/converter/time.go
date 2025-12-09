package converter

import (
	"fmt"
	"time"
)

func FormatDateRussian(t time.Time) string {
	months := []string{
		"января", "февраля", "марта", "апреля",
		"мая", "июня", "июля", "августа",
		"сентября", "октября", "ноября", "декабря",
	}
	return fmt.Sprintf("%d %s %d, %02d:%02d",
		t.Day(),
		months[t.Month()-1],
		t.Year(),
		t.Hour(),
		t.Minute())
}
