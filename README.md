# jsed: A JSON Editor

jsed is a small command-line utility to add, remove, and search for data in a JSON structure.

Not to be confused with any other [jsed](https://github.com/search?q=jsed&type=Repositories).

## Examples

```shell
$ echo {} | jsed add key --path foo --value bar
{
  "foo": "bar"
}

$ echo {} | jsed add key --path foo --value bar | jsed add array --path bar.baz --value a --value b --value c
{
  "bar": {
    "baz": [
      "a",
      "b",
      "c"
    ]
  },
  "foo": "bar"
}

$ echo {} | jsed add key --path foo --value bar | jsed add array --path bar.baz --value a --value b --value c | jsed get --path bar.baz.1
b
```
