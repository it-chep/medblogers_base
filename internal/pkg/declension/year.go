package declension

import "strconv"

func decline(n int, nominative, genitiveSingular, genitivePlural string) string {
	if n < 0 {
		n = -n
	}

	lastTwo := n % 100
	lastOne := n % 10

	if lastTwo >= 11 && lastTwo <= 19 {
		return strconv.Itoa(n) + " " + genitivePlural
	}

	if lastOne == 1 {
		return strconv.Itoa(n) + " " + nominative
	}

	if lastOne >= 2 && lastOne <= 4 {
		return strconv.Itoa(n) + " " + genitiveSingular
	}

	return strconv.Itoa(n) + " " + genitivePlural
}

// YearsInString для лет
func YearsInString(n int) string {
	return decline(n, "год", "года", "лет")
}

// MonthsInString - для месяцев
func MonthsInString(n int) string {
	return decline(n, "месяц", "месяца", "месяцев")
}

// DaysInString - для дней
func DaysInString(n int) string {
	return decline(n, "день", "дня", "дней")
}
