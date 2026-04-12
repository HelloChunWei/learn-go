package headers

import (
	"bytes"
	"fmt"
	"regexp"
	"strings"
)

type Headers struct {
	headers map[string]string
}

var SEPARATOR = []byte("\r\n")

func NewHeaders() *Headers {
	return &Headers{
		headers: map[string]string{},
	}
}

func (h *Headers) Get(name string) string {
	return h.headers[strings.ToLower(name)]
}

func (h *Headers) set(name, value string) {
	lowerName := strings.ToLower(name)
	if v, ok := h.headers[lowerName]; ok {
		h.headers[strings.ToLower(name)] = v + "," + value
	} else {
		h.headers[strings.ToLower(name)] = value
	}
}

func (h *Headers) ForEach(cb func(n, v string)) {
	for n, v := range h.headers {
		cb(n, v)
	}
}

var ERROR_BAD_HEADER = fmt.Errorf("bad header")
var ERROR_BAD_KEY = fmt.Errorf("bad filed name")

func isValidateKey(str string) bool {
	// 只允許 A-Z, a-z, 0-9 和特定特殊字元
	pattern := `^[A-Za-z0-9!#$%&'*+\-.^_` + "`" + `|~]+$`
	matched, _ := regexp.MatchString(pattern, str)
	return matched
}

func parseHeader(data []byte) (string, string, error) {
	parts := bytes.SplitN(data, []byte(":"), 2)
	if len(parts) != 2 {
		return "", "", ERROR_BAD_HEADER
	}
	if bytes.HasSuffix(parts[0], []byte(" ")) {
		return "", "", ERROR_BAD_HEADER
	}

	strKey := string(parts[0])
	if !isValidateKey(strKey) {
		return "", "", ERROR_BAD_KEY
	}
	strValue := string(bytes.TrimSpace(parts[1]))
	return strings.ToLower(strKey), strValue, nil

}

func (h *Headers) Parse(data []byte) (int, bool, error) {
	read := 0
	done := false

	for {
		idx := bytes.Index(data[read:], SEPARATOR)
		if idx == -1 {
			break
		}
		// empty header
		if idx == 0 {
			done = true
			read += len(SEPARATOR)
			break
		}
		key, value, err := parseHeader(data[read : read+idx])
		if err != nil {
			return 0, false, err
		}
		read += idx + len(SEPARATOR)
		h.set(key, value)
	}
	return read, done, nil

}
