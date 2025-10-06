package slug

import (
	"regexp"
	"strings"

	"github.com/rainycape/unidecode"
)

// New создание нового URL-дружественный slug а с использованием unidecode
func New(fullName string) string {
	// Транслитерация с unidecode (конвертирует кириллицу и другие символы в латиницу)
	transliterated := unidecode.Unidecode(fullName)

	// Заменяем все не-буквенно-цифровые символы на дефисы
	reg := regexp.MustCompile(`[^a-zA-Z0-9]+`)
	slug := reg.ReplaceAllString(transliterated, "-")

	// Приводим к нижнему регистру и обрезаем дефисы по краям
	slug = strings.ToLower(slug)
	slug = strings.Trim(slug, "-")

	return slug
}
