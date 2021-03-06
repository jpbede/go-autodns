# InternetX JSON API client package for Go
[![PkgGoDev](https://pkg.go.dev/badge/go.bnck.me/autodns)](https://pkg.go.dev/go.bnck.me/autodns)
[![Codacy Badge](https://app.codacy.com/project/badge/Grade/d8f56b33ce9c4653983e81e0ab8a3b8c)](https://www.codacy.com/gh/jpbede/go-autodns/dashboard)
[![codecov](https://codecov.io/gh/jpbede/go-autodns/branch/main/graph/badge.svg?token=ACJ41YHXN1)](https://codecov.io/gh/jpbede/go-autodns)
![test](https://github.com/jpbede/go-autodns/workflows/test/badge.svg)

This repository contains a Go package for accessing the [InternetX JSON API](https://help.internetx.com/display/APIXMLEN/JSON+API+Basics).

## Installation
Install using go get:
```shell
go get go.bnck.me/autodns
```

## Usage
Import the lib as usual
```go
import "go.bnck.me/autodns"
```

Create a new client without options:
```go
autodnsClient, err := autodns.New("username", "password", 1) // change 1 with your context number
```
The client now offers more specific sub-clients, for example for managing contacts or domains.
For more information have a look at the [documentation](https://pkg.go.dev/github.com/jpbede/go-autodns).
