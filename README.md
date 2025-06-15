# GHOST
[![GitHub license][License img]][License src] [![GitHub release][Release img]][Release src] [![Github main status][Github main status badge]][Github main status src] [![Conventional Commits][Conventional commits badge]][Conventional commits src]

Go High-level Open Service Templater

GHOST creates in working dir fully functional microservice with basic functionality:

- Ready to use Makefile with help target
- Ready to use docker-compose with override template
- Ready to use security checks by trivy, semgrep and govulncheck
- Kubernetes livecycle compatible health checks
- Log level changing by HTTP-request
- Easy configurable logs (plain text|JSON)
- Graceful shutdown

## Usage

<details>
  <summary>Example</summary>

```shell
$ ghost -h
Usage: ghost [options...]

OPTIONS
  -d, --description        <string>    (default: Go microservice)                Project short description
  -g, --go-image           <string>    (default: nafigat0r/go:1.24.3)            Go docker image
  -c, --govulncheck-image  <string>    (default: nafigat0r/govulncheck:1.1.4)    Govulncheck docker image
  -h, --help                                                                     display this help message
  -l, --linter-image       <string>    (default: nafigat0r/golangci-lint:2.1.6)  Linter docker image
  -m, --module-name        <string>    (default: github.com/test/test)           Go module name
  -n, --name               <string>    (default: test)                           Project short name
  -t, --shutdown-timeout   <duration>  (default: 10s)                            Timeout for graceful shutdown
  -v, --version                                                                  display version
  -r, --with-rest          <bool>      (default: false)                          Add HTTP server with REST API functionality

ENVIRONMENT
  GHOST_DESCRIPTION        <string>    (default: Go microservice)                Project short description
  GHOST_GO_IMAGE           <string>    (default: nafigat0r/go:1.24.3)            Go docker image
  GHOST_GOVULNCHECK_IMAGE  <string>    (default: nafigat0r/govulncheck:1.1.4)    Govulncheck docker image
  GHOST_LINTER_IMAGE       <string>    (default: nafigat0r/golangci-lint:2.1.6)  Linter docker image
  GHOST_MODULE_NAME        <string>    (default: github.com/test/test)           Go module name
  GHOST_NAME               <string>    (default: test)                           Project short name
  GHOST_SHUTDOWN_TIMEOUT   <duration>  (default: 10s)                            Timeout for graceful shutdown
  GHOST_WITH_REST          <bool>      (default: false)                          Add HTTP server with REST API functionality
```
</details>

#### Docker

<details>
  <summary>Run via docker</summary>

```shell
docker run \
    -u $(id -u):$(id -g) --rm -ti \
    -v "$(pwd):/var/ghost" \
    nafigat0r/ghost --help
```
</details>

## Versioning
This software follows *"Semantic Versioning"* specifications. The signature of exported package functions is used
as a public API. Read more on [SemVer.org][semver src].


[License img]: https://img.shields.io/github/license/nafigator/ghost?color=teal
[License src]: https://www.tldrlegal.com/license/mit-license
[Release img]: https://img.shields.io/github/v/tag/nafigator/ghost?logo=github&color=teal&filter=!*/*
[Release src]: https://github.com/nafigator/ghost
[Github main status src]: https://github.com/nafigator/ghost/actions/workflows/go.yml?query=branch%3Amain
[Github main status badge]: https://github.com/nafigator/ghost/actions/workflows/go.yml/badge.svg?branch=main
[Conventional commits src]: https://conventionalcommits.org
[Conventional commits badge]: https://img.shields.io/badge/Conventional%20Commits-1.0.0-teal.svg
[semver src]: http://semver.org
