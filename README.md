[![logo](./logos/mobingi-205x119.png)](https://mobingi.com/)

[![CircleCI](https://circleci.com/gh/mobingi/mobingi/tree/master.svg?style=svg)](https://circleci.com/gh/mobingi/mobingi/tree/master)
[![Build status](https://ci.appveyor.com/api/projects/status/98feyjxtf91xhrdf/branch/master?svg=true)](https://ci.appveyor.com/project/flowerinthenight/mobingi/branch/master)

# Overview

`mobingi` is the official command line interface for [Mobingi](https://mobingi.com/) API and services. 

See the documentation on https://docs.mobingi.com/v/cli-en/.

# Getting started

### Getting the cli

The easiest way to get `mobingi` is to use one of the [pre-built release binaries](https://github.com/mobingi/mobingi/releases) which are available for OSX, Linux, and Windows.

If you want to try the latest version, you can build `mobingi` from the master branch. You need to have [Go](https://golang.org/) installed (version 1.8+ is required). Note that the master branch may be in an unstable or even broken state during development.

### Building the cli

```bash
$ git clone https://github.com/mobingi/mobingi
$ cd mobingi
$ go build -v
$ ./mobingi
```

You can also install the binary to your `$GOPATH/bin` folder (`$GOPATH/bin` should be added to your `$PATH` environment variable). 

```bash
$ go get -u -v github.com/mobingi/mobingi
$ mobingi
```
