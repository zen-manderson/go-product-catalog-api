# Installing Go

Go is easily installed, either by downloading a Mac `pkg` file from [go.dev/dl](https://go.dev/dl/) or using
Homebrew's `brew install go`. If you download from [go.dev/dl](https://go.dev/dl/), make sure you pick the
right option for your Mac's CPU chip, `go1.20.3.darwin-arm64.pkg` is easily confused with `go1.20.3.darwin-amd64.pkg`!

The [go.dev/dl](https://go.dev/dl/) download approach can make life a little easier when it comes to establishing and
maintaining the environment variables that Go requires to function correctly. Why? Here's the Go section from Mike
Broadway's `.zshrc` configuration file:

```shell
# Enable Golang use
GOPATH=/opt/homebrew/Cellar/go/1.20.3
GOBIN=$HOME/go/bin
PATH=$PATH:$GOPATH/bin
PATH=$PATH:$GOBIN
GOPRIVATE="github.com/zenbusiness/*"
GONOSUMDB="github.com/zenbusiness/*"
```

That Homebrew `/opt/homebrew/Cellar/go/1.20.3` nonsense in the `GOPATH` can be hard to figure out. And every time
you run `brew upgrade` it might change where Go is installed from `.../Cellar/go/1.20.3` to `.../Cellar/go/1.20.4`
without you realizing it.

## Grpcurl
Install `grpcurl` for sending server requests locally
```shell
brew install grpcurl
```

## The `GOPRIVATE` and `GONOSUMDB` environment variables

Go's equivalent of an npm Registry or Python package registry is the source code of the module stored in GitHub,
GitLab, BitBucket, etc. To reference Go library modules in private Git servers, i.e., [github/zenbusiness](https://github.com/zenbusiness),  
the `GOPRIVATE` environment variable must be set:

```shell
export GOPRIVATE="github.com/zenbusiness/*"
```

Don't forget to set it in your IDE too!

This prevents the Go toolset from trying to verify cache the source code of our private repos in the `https://proxy.golang.org`
module mirror and, in theory, not track the checksum values in the https://sum.golang.org/ checksum index database.
index, and checksum service. Together, `https://proxy.golang.org` and https://sum.golang.org/ act as a kind of master
index and cache of public Go modules hosted on GitHub, GitLab, BitBucket, and elsewhere. If `GOPRIVATE` is not set,
these services will try to verify the module checksums, not be able to connect to the GitHub repository, and fail any
build or reference checks your Go tools are attempting.

For exhaustive reference documentation on modules and these environment variables,
see [Go Modules Reference](https://go.dev/ref/mod).

### Extreme measures

If you are getting multi-line error messages with this nugget in the middle:

```text
     could not read Username for 'https://github.com': terminal prompts disabled
```

It is probably either that you need to set up a personal access token (see below) or that `GOPRIVATE` or `GONUSUMDB`
need to be set. If you have done all that and are still getting the error, there is a sledge hammer left to try.

First check if the displays of `go env GOPRIVATE` and `go env GONUSUMDB` match your environment variable settings
for `GOPRIVATE` and `GONUSUMDB`. If they don't, then something somewhere has set master override values using
`go env -w`. Before you throw your keyboard through you monitor, try this:

```shell
go env -w GOPRIVATE="github.com/zenbusiness/*"
go env -w GONOSUMDB="github.com/zenbusiness/*"
```

That will set the values globally, for all terminal sessions, and override any environment variable settings
you may have.
