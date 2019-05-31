**giuser** is a command line tool for searching users and repositories information from github.com. It's built with `Go` language.


## Dependencies
If you haven't `Go` installed on your machine, then you can install it by the following commands:

1. Download the `go package` from [here](https://golang.org/dl/). 
2. Download and extract it into `/usr/local`, creating a Go tree in `/usr/local/go`. For example:

```
tar -C /usr/local -xzf go$VERSION.$OS-$ARCH.tar.gz
```

Choose the archive file appropriate for your installation. For instance, if you are installing Go version 1.2.1 for 64-bit x86 on Linux, the archive you want is called `go1.2.1.linux-amd64.tar.gz`.

(Typically these commands must be run as root or through sudo.)

Add /usr/local/go/bin to the PATH environment variable. You can do this by adding these line to your `/etc/profile` (for a system-wide installation):

```
export PATH=$PATH:/usr/local/go/bin
export GOPATH=/home/your_username/go
export GOBIN=$GOPATH/bin
PATH=$PATH:$GOPATH:$GOBIN
export PATH
```
## Installation
Install this cli by the following command:
```
go get github.com/SamsadSajid/giuser
```

## Usage
For help, type `giuser` in your terminal and it will show all available commands and it's usage