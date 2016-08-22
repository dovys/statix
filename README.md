# [Statix](https://dovys.github.io/statix)

Statix is a simple web server for static site development that works with Linux, macOS & Windows.

Not intended to be run on any production environment.

[Download](https://dovys.github.io/statix)

## Usage
Serving current directory
```sh
cd /path/to/project
./statix

2016/08/22 11:36:03 Listening on :3003 and serving /Users/dovys/dev/dovyscom
```

**Additional flags**
```sh
  --host # Host.
  --port # Port to listen on. (default 3003)
  --path # Path to serve files from. Default: current directory.
```

## Building from source
Install [golang](https://golang.org/dl/)
```sh
go get github.com/dovys/statix
go build github.com/dovys/statix
```
