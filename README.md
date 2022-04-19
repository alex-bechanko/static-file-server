# static-file-server
A simple static file server that can be configured to return a specific file from provided route.
Built so that I could primarily use golang as my toolchain without pulling in other language or OS dependencies.
This package is meant for extremely simple development work and should absolutely not be used for anything production-facing.

## Build

```
go build -o static-file-server main.go
```

## Usage
To see help information:

```
./static-file-server --help
```