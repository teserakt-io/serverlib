![alt text](logo.png)

[![GoDoc][godoc-image]][godoc-url] ![Go](https://github.com/teserakt-io/serverlib/workflows/Go/badge.svg?branch=master)

# Serverlib

Reusable modules across Teserakt server applications

## Modules

`config` module holds a layer on top of [viper](https://github.com/spf13/viper) providing some helpers easing configuration definition and validation.

`path` module does provide a path resolver, helping building path to files from the configuration file location.

## Testing

```
./scripts/unittest.sh
```

[godoc-image]: https://godoc.org/github.com/teserakt-io/serverlib?status.svg
[godoc-url]: https://godoc.org/github.com/teserakt-io/serverlib
