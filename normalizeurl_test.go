package normalizeurl

import (
	"testing"
)

func TestNormalize(t *testing.T) {
	tests := []struct {
		value    string
		expected string
	}{
		{"www.sekimura.org", "http://sekimura.org"},
		{"sekimura.org", "http://sekimura.org"},
		{"HTTP://sekimura.org", "http://sekimura.org"},
		{"//sekimura.org", "http://sekimura.org"},
		{"http://sekimura.org", "http://sekimura.org"},
		{"http://sekimura.org:80", "http://sekimura.org"},
		{"https://sekimura.org:443", "https://sekimura.org"},
		{"ftp://sekimura.org:21", "ftp://sekimura.org"},
		{"http://www.sekimura.org", "http://sekimura.org"},
		{"www.sekimura.org", "http://sekimura.org"},
		{"http://sekimura.org/foo/", "http://sekimura.org/foo"},
		{"sekimura.org/?foo=bar%20baz", "http://sekimura.org/?foo=bar baz"},
		{"http://sekimura.org/?", "http://sekimura.org"},
		{"http://xn--xample-hva.com", "http://êxample.com"},
		{"http://sekimura.org/?b=bar&a=foo", "http://sekimura.org/?a=foo&b=bar"},
	}

	for i := range tests {
		actual, err := Normalize(tests[i].value)
		if err != nil {
			t.Error("normalize failed", err)
		}
		if actual != tests[i].expected {
			t.Error("normalize did not match expected value", i, actual, tests[i].expected)
		}
	}
}
