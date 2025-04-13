package utils

import (
	"net/url"
	"strings"
)

func ExtractDomain(raw string) string {
	u, err := url.Parse(raw)
	if err != nil {
		return ""
	}
	host := u.Hostname()
	return strings.TrimPrefix(host, "www.")
}
