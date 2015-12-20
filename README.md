# jsed: A JSON Editor

jsed is a small command-line utility to add, remove, and search for data in a JSON structure.

Not to be confused with any other [jsed](https://github.com/search?q=jsed&type=Repositories).

## Examples

```shell
$ echo {} | jsed add object --path foo --value bar -r
{
  "foo": "bar"
}

$ echo {} | jsed add object --path foo --value bar | jsed add array --path bar.baz --value a --value b --value c -r
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

$ echo {} | jsed add object --path foo --value bar | jsed add array --path bar.baz --value a --value b --value c | jsed get --path bar.baz.1
b

$ echo {} | jsed add object --path service \
                            --key name --value redis_master \
                            --key address --value 127.0.0.1 \
                            --key port --value 8000 \
                            --key enableTagOverride --value false \
                            --key checks --value []  \
          | jsed add array --path service.tags --value master --value redis --value mysql \
          | jsed add object --path service.checks --key script --value /usr/local/bin/check_redis.py --key interval --value 10s \
          | jsed add object --path service.checks --key script --value /usr/local/bin/check_mysql.py --key interval --value 10s -r > service.json

$ cat service.json
{
  "service": {
    "address": "127.0.0.1",
    "checks": [
      {
        "interval": "10s",
        "script": "/usr/local/bin/check_redis.py"
      },
      {
        "interval": "10s",
        "script": "/usr/local/bin/check_mysql.py"
      }
    ],
    "enableTagOverride": false,
    "name": "redis_master",
    "port": 8000,
    "tags": [
      "master",
      "redis",
      "mysql"
    ]
  }
}

$ jsed get --file service.json --path service..checks..*..script=/usr/local/bin/check_redis.py --delimiter ..
/usr/local/bin/check_redis.py

$ jsed get --file service.json --path service.checks.* -r
[
  {
    "interval": "10s",
    "script": "/usr/local/bin/check_redis.py
  },
  {
    "interval": "10s",
    "script": "/usr/local/bin/check_mysql.py
  }
]
```
