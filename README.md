# pipe-editor

Using visual editor in the pipeline.

```sh
$ export EDITOR=vi    # set the editor (vi is default)
$ echo hello | pe     # stdin ->   vi
modified_in_editor    #   vi  -> stdout
```

## Installation

You need to set up Go and `$GOPATH`.

```sh
$ export PATH=$PATH:$GOPATH/bin
$ go get github.com/shrhdk/pipe-editor/cmd/pe
```
