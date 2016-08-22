# [Statix](https://dovys.github.io/statix)

Statix is a simple web server for static site development that works with Linux, macOS & Windows.

Not intended to be run on any production environment.

[Download](https://dovys.github.io/statix)

## Usage
Serving current directory
```sh
cd /path/to/project
./statix

2016/08/22 11:36:03 Listening on :3003 and serving /path/to/project
```

**Additional flags**
```sh
./statix --port 8080 --host 127.0.0.1 --path ~/www

2016/08/22 22:26:45 Listening on 127.0.0.1:8080 and serving /Users/dovys/www
```

## Building from source
Install [golang](https://golang.org/dl/)
```sh
go get github.com/dovys/statix
go build github.com/dovys/statix
```
