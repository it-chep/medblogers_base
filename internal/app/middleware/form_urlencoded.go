package middleware

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// FormURLEncodedMiddleware универсальная middleware для преобразования form-urlencoded в JSON
func FormURLEncodedMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		contentType := r.Header.Get("Content-Type")

		// Проверяем, является ли запрос form-urlencoded
		if strings.Contains(contentType, "application/x-www-form-urlencoded") {
			log.Printf("[FormMiddleware] Processing form-urlencoded request to %s", r.URL.Path)

			// Парсим форму
			if err := r.ParseForm(); err != nil {
				log.Printf("[FormMiddleware] Error parsing form: %v", err)
				http.Error(w, "Bad form data", http.StatusBadRequest)
				return
			}

			// Логируем данные формы для отладки
			if len(r.Form) > 0 {
				log.Printf("[FormMiddleware] Form data for %s: %v", r.URL.Path, r.Form)
			}

			// Преобразуем form данные в JSON
			jsonData, err := convertFormToUniversalJSON(r.Form)
			if err != nil {
				log.Printf("[FormMiddleware] Error converting form to JSON: %v", err)
				http.Error(w, "Internal server error", http.StatusInternalServerError)
				return
			}

			log.Printf("[FormMiddleware] Converted JSON for %s: %s", r.URL.Path, jsonData)

			// Подменяем тело запроса на JSON
			r.Body = &customReader{
				data:   []byte(jsonData),
				offset: 0,
			}
			r.ContentLength = int64(len(jsonData))
			r.Header.Set("Content-Type", "application/json")
		}

		next.ServeHTTP(w, r)
	})
}

// universalFormConverter преобразует form данные в универсальный JSON
func convertFormToUniversalJSON(form url.Values) (string, error) {
	result := make(map[string]interface{})

	for key, values := range form {
		// Преобразуем ключи из camelCase в snake_case
		normalizedKey := normalizeKey(key)

		// Обрабатываем массивы (поля вида field[], field[0], field[1])
		if isArrayField(key) {
			arrayKey := getArrayKeyName(key)
			if _, exists := result[arrayKey]; !exists {
				result[arrayKey] = []interface{}{}
			}

			// Добавляем все значения в массив
			for _, value := range values {
				parsedValue := parseValue(value)
				result[arrayKey] = append(result[arrayKey].([]interface{}), parsedValue)
			}
		} else {
			// Обычное поле (не массив)
			if len(values) == 1 {
				// Одно значение
				result[normalizedKey] = parseValue(values[0])
			} else {
				// Несколько значений (преобразуем в массив)
				parsedValues := make([]interface{}, len(values))
				for i, value := range values {
					parsedValues[i] = parseValue(value)
				}
				result[normalizedKey] = parsedValues
			}
		}
	}

	// Конвертируем в JSON
	jsonData, err := json.Marshal(result)
	if err != nil {
		return "", err
	}

	return string(jsonData), nil
}

// normalizeKey преобразует ключи из разных форматов в snake_case
func normalizeKey(key string) string {
	// Удаляем индексы массивов: positions[0] -> positions
	if idx := strings.Index(key, "["); idx != -1 {
		key = key[:idx]
	}

	// Удаляем суффикс []: positions[] -> positions
	key = strings.TrimSuffix(key, "[]")

	// Преобразуем camelCase в snake_case
	var result []rune
	for i, r := range key {
		if i > 0 && 'A' <= r && r <= 'Z' {
			result = append(result, '_')
		}
		result = append(result, r)
	}

	return strings.ToLower(string(result))
}

// isArrayField проверяет, является ли поле массивом
func isArrayField(key string) bool {
	return strings.Contains(key, "[]") || strings.Contains(key, "[")
}

// getArrayKeyName извлекает имя ключа массива
func getArrayKeyName(key string) string {
	if strings.Contains(key, "[]") {
		return strings.TrimSuffix(key, "[]")
	}

	if idx := strings.Index(key, "["); idx != -1 {
		return key[:idx]
	}

	return key
}

// parseValue пытается определить тип значения
func parseValue(value string) interface{} {
	// Пробуем как integer
	if intVal, err := strconv.ParseInt(value, 10, 64); err == nil {
		return intVal
	}

	// Пробуем как float
	if floatVal, err := strconv.ParseFloat(value, 64); err == nil {
		return floatVal
	}

	// Пробуем как boolean
	if boolVal, err := strconv.ParseBool(value); err == nil {
		return boolVal
	}

	// Если не удалось преобразовать, возвращаем как строку
	return value
}

// customReader для подмены тела запроса
type customReader struct {
	data   []byte
	offset int64
}

func (r *customReader) Read(p []byte) (n int, err error) {
	if r.offset >= int64(len(r.data)) {
		return 0, io.EOF
	}

	n = copy(p, r.data[r.offset:])
	r.offset += int64(n)
	return n, nil
}

func (r *customReader) Close() error {
	r.offset = 0
	return nil
}
