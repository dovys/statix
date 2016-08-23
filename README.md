# [Statix](https://dovys.github.io/statix)

Statix is a simple web server for static site development that works with Linux, macOS & Windows.

Not intended to be run on any production environment.

## Install

#### macOS
```sh
wget https://github.com/dovys/statix/releases/download/1.0.0/statix-macOS-v1.0.0-darwin-10.6-386.tar.gz
tar -C /usr/local/bin -xzf statix-macOS-v1.0.0-darwin-10.6-386.tar.gz
```

#### Linux
```sh
wget https://github.com/dovys/statix/releases/download/1.0.0/statix-linux-v1.0.0-linux-386.tar.gz
tar -C /usr/local/bin -xzf statix-linux-v1.0.0-linux-386.tar.gz
```

#### Windows
[Download statix.exe](https://github.com/dovys/statix/releases/download/1.0.0/windows-v1.0.0-windows-4.0-386.zip)

## Usage
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
