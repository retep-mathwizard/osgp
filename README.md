# Go at SkilStak

This repo contains the Go libraries, commands, and services used
at SkilStak. This includes all of our core IT infrastructure, which
we release to public domain for others to use and improve upon. All
internal development is done in Go or as native web apps. Many of
the main SkilStak commands and libraries are included in this repo
such as `bux`, `sks`, `wishes`, `colors` and more.

SkilStak admins and users will likely want to look at
[usr-share-skilstak][] and [home-config][] as well, which contain
compiled binaries of many of these commands for Linux AMD64 by
default.

## Development

Those doing SkilStak Go development will want to clone this repo
locally and either commit their changes or make pull requests. 

## Deployment

Any changes will need to be built for Linux AMD64 and the binaries
added to the [usr-share-skilstak][] and [home-config][] repos.  To
do this developers can setup a read-only go user setup (as for any student)
in the `admin` account:

```sh
sudo su - admin
mkdir go
mkdir go/bin
mkdir go/src
mkdir repos
cd repos
golink go skilstak
```

The `admin` account should use the `youraccount_rsa` key pair to be able to
read/clone.

## Install

It's best to do specific `go build -o /usr/share/skilstak/bin/<name>`
commands from the target system.

Oh, if you want to clean up the code automatically with `go fmt`:

```sh
ln -sf ../../pre-commit .git/hooks/pre-commit
```

## Cross Compiling

Although most of this code runs on Ubuntu Linux servers at SkilStak
it is authored often on Mac (Darwin) systems. We realize that
`user.Current()` from `os/user` and others do not cross-compile and
have accepted the extra workflow task of compiling on the target
Linux machines instead of cross-compiling locally.

[usr-share-skilstak]: http://github.com/skilstak/usr-share-skilstak
[home-config]: http://github.com/skilstak/home-config
