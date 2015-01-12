package normalizeurl

import (
    "testing"
)

func TestNormalize(t *testing.T) {
    tests := []struct{
        value string
        expected string
    }{
        {"www.sindresorhus.com", "http://sindresorhus.com"},
        {"sindresorhus.com", "http://sindresorhus.com"},
        {"HTTP://sindresorhus.com", "http://sindresorhus.com"},
        {"//sindresorhus.com", "http://sindresorhus.com"},
        {"http://sindresorhus.com", "http://sindresorhus.com"},
        {"http://sindresorhus.com:80", "http://sindresorhus.com"},
        {"https://sindresorhus.com:443", "https://sindresorhus.com"},
        {"ftp://sindresorhus.com:21", "ftp://sindresorhus.com"},
        {"http://www.sindresorhus.com", "http://sindresorhus.com"},
        {"www.sindresorhus.com", "http://sindresorhus.com"},
        {"http://sindresorhus.com/foo/", "http://sindresorhus.com/foo"},
        {"sindresorhus.com/?foo=bar%20baz", "http://sindresorhus.com/?foo=bar baz"},
        {"http://sindresorhus.com/?", "http://sindresorhus.com"},
        {"http://xn--xample-hva.com", "http://êxample.com"},
        {"http://sindresorhus.com/?b=bar&a=foo", "http://sindresorhus.com/?a=foo&b=bar"},
    }

	for i := range tests {
		actual, err := normalize(tests[i].value)
		if err != nil {
			t.Error("normalize failed", err)
		}
		if actual != tests[i].expected {
			t.Error("normalize did not match expected value", i, actual, tests[i].expected)
		}
	}
}
