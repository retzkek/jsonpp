# jsonpp
Simple JSON pretty-printer. Reads from stdin, writes to stdout.

## Install

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
