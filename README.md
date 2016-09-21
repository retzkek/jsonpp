# jsonpp
Simple JSON pretty-printer. Reads from stdin, writes to stdout.

## Install

Requires go 1.7 or newer https://golang.org/doc/install

```
$ go get github.com/retzkek/jsonpp
```

## Use

```
$ echo '{"foo":{"bar":"baz"}}' | jsonpp
{
    "foo": {
        "bar": "baz"
    }
}
```
