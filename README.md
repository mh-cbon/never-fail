# never-fail

package never-fail is a cli command wrapper to run a command and ignore exit code


On linux it is the same as `whatever || echo "ignore exit code of previous command"`.

Using that binary this capability is available for any OS.

# TOC
- [Install](#install)
  - [Go](#go)
- [Cli](#cli)

# Install

#### Go
```sh
go get github.com/mh-cbon/never-fail
```

# Cli

###### $ never-fail -h
```sh
never-fail - head

USAGE
 never-fail <bin> <args>
 never-fail --help|-h
```
