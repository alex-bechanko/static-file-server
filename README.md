# static-file-server
A simple static file server that can be configured to return a specific file from provided route.
Built so that I could primarily use golang as my toolchain without pulling in other language or OS dependencies.

## Warning
This package is meant for extremely simple development work and should absolutely not be used for anything production-facing.
For the ease of development, this server set's it's CORS completely open.
This is a huge security vulnerability if used outside of local development, which is why you should not use it for external facing applications/users.

## Build

```
go build -o static-file-server main.go
```

## Usage
To see help information:

```
./static-file-server --help
```