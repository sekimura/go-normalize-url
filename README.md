# go-normalize-url [![Build Status](https://travis-ci.org/sekimura/go-normalize-url.svg?branch=master)](https://travis-ci.org/sekimura/go-normalize-url)

> [URL normalization](http://en.wikipedia.org/wiki/URL_normalization) tool for Go

Useful when you need to display, store, deduplicate, sort, compare, etc, URLs.


## Install

```
$ go get github.com/sekimura/go-nomalize-url
```


## Usage

```go
package main

import (
    "github.com/sekimura/go-normalize-url"
)

func main () {
    s, _ := normalizeurl.Normalize("sekimura.org")
    println(s)
    //=> http://sekimura.org

    s, _ = normalizeurl.Normalize("HTTP://xn--xample-hva.com:80/?b=bar&a=foo")
    println(s)
    //=> http://êxample.com/?a=foo&b=bar
}
```

## License

MIT © [Masayoshi Sekimura](http://sekimura.org)

