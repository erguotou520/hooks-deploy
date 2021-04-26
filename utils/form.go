package utils

import (
	"net/http"
	"regexp"
	"strconv"
)

func ParseFormCollection(r *http.Request, typeName string) []map[string]string {
	var result []map[string]string
	r.ParseForm()
	for key, values := range r.Form {
		re := regexp.MustCompile(typeName + "\\[([0-9]+)\\]\\[([a-zA-Z]+)\\]")
		matches := re.FindStringSubmatch(key)

		if len(matches) >= 3 {

			index, _ := strconv.Atoi(matches[1])

			for index >= len(result) {
				result = append(result, map[string]string{})
			}

			result[index][matches[2]] = values[0]
		}
	}
	return result
}
