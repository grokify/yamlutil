# YAML Utility for Go

[![Used By][used-by-svg]][used-by-url]
[![Build Status][build-status-svg]][build-status-url]
[![Go Report Card][goreport-svg]][goreport-url]
[![Docs][docs-godoc-svg]][docs-godoc-url]
[![License][license-svg]][license-url]

## Overview

The `yamlutil` package provides a collection of Go utilities for YAML.

1. The primary feature is the `GetNodeJsonSchemaPath()` function which returns a `*yaml.Node` for a pre-parsed JSON Schema pointer path. The primary use case is to get line numbers of linters using JSON Schema pointer to access a YAML file, e.g. OpenAPI specification files.
1. A secondary feature is the `ReadFileAsJson()` function which will read a YAML file and convert it to JSON using `github.com/ghodss/yaml`.

## Installation

```bash
$ go get github.com/grokify/yamlutil
```

## Contributing

Features, Issues, and Pull Requests are always welcome.

To contribute:

1. Fork it ( http://github.com/grokify/yamlutil/fork )
2. Create your feature branch (`git checkout -b my-new-feature`)
3. Commit your changes (`git commit -am 'Add some feature'`)
4. Push to the branch (`git push origin my-new-feature`)
5. Create new Pull Request

Please report issues and feature requests on [Github](https://github.com/grokify/yamlutil).

 [used-by-svg]: https://sourcegraph.com/github.com/grokify/yamlutil/-/badge.svg
 [used-by-url]: https://sourcegraph.com/github.com/grokify/yamlutil?badge
 [build-status-svg]: https://github.com/grokify/yamlutil/workflows/test/badge.svg?branch=master
 [build-status-url]: https://github.com/grokify/yamlutil/actions
 [goreport-svg]: https://goreportcard.com/badge/github.com/grokify/yamlutil
 [goreport-url]: https://goreportcard.com/report/github.com/grokify/yamlutil
 [docs-godoc-svg]: https://pkg.go.dev/badge/github.com/grokify/yamlutil
 [docs-godoc-url]: https://pkg.go.dev/github.com/grokify/yamlutil
 [license-svg]: https://img.shields.io/badge/license-MIT-blue.svg
 [license-url]: https://github.com/grokify/yamlutil/blob/master/LICENSE
