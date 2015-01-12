package normalizeurl

import (
	"fmt"
	"golang.org/x/net/idna"
	"net/url"
	"strings"
)

var (
	DefaultPorts = map[string]int{
		"http":  80,
		"https": 443,
		"ftp":   21,
	}
)

// Normalize url strings
// http://en.wikipedia.org/wiki/URL_normalization
func Normalize(s string) (string, error) {
	s = strings.TrimSpace(s)
	if strings.HasPrefix(s, "//") {
		s = "http:" + s
	}

	u, err := url.Parse(s)
	if err != nil {
		return s, err
	}

	if u.Scheme == "" {
		// Ugh...
		u, err = url.Parse("http://" + s)
		if err != nil {
			return s, err
		}
	}

	p, ok := DefaultPorts[u.Scheme]
	if ok {
		u.Host = strings.TrimSuffix(u.Host, fmt.Sprintf(":%d", p))
	}

	got, err := idna.ToUnicode(u.Host)
	if err != nil {
		return got, err
	} else {
		u.Host = got
	}

	u.Host = strings.TrimPrefix(u.Host, "www.")

	v := u.Query()
	u.RawQuery = v.Encode()
	u.RawQuery, _ = url.QueryUnescape(u.RawQuery)

	h := u.String()
	h = strings.TrimSuffix(h, "/")

	return h, nil
}
