# Express GO

![Golang EXpress Banner](https://www.educative.io/v2api/editorpage/5143839726108672/image/5395289693749248)

Fast, unopinionated, minimalist web framework for GO.


## Install

With a [correctly configured](https://golang.org/doc/install#testing) Go toolchain:

```sh
go get -u github.com/WatchJani/Express
```


## Examples

Let's start registering a couple of URL paths and handlers:

```go
func main() {
    app := express.New()

    app.Route("/").GET(GetData).POST(PostData)

    app.Listen("5000")
}
```

## Features

  * Robust routing
  * Focus on high performance
  * Executable for generating applications quickly
