## Usage example
```
./pinger --target=mail.ru --target=8.8.8.8 --target=yandex.ru --listen=0.0.0.0:9123 --username=test --password=test
```

To view ping results
```
curl --user test:test http://localhost:9123 2>/dev/null | python -m json.tool
```

```json
{
    "8.8.8.8": {
        "best": 84.85044,
        "count": 114,
        "last": 84.95047,
        "loss10": 0,
        "loss30": 0,
        "loss300": 0,
        "mean": 84.929794,
        "median": 84.92787,
        "stddev": 0.031026525,
        "wrost": 85.044525
    },
    "mail.ru": {
        "best": 112.74954,
        "count": 114,
        "last": 115.91721,
        "loss10": 0,
        "loss30": 0,
        "loss300": 0,
        "mean": 115.32941,
        "median": 115.173996,
        "stddev": 0.5426216,
        "wrost": 116.57005
    },
    "yandex.ru": {
        "best": 70.343025,
        "count": 114,
        "last": 70.42257,
        "loss10": 0,
        "loss30": 0,
        "loss300": 0,
        "mean": 70.47223,
        "median": 70.42935,
        "stddev": 0.09575593,
        "wrost": 70.76251
    }
}
```

Options and defaults

```
Usage of ./pinger:
  -listen string
    	Bind internal http server to address and port. Default: 0.0.0.0:9123 (default "0.0.0.0:9123")
  -password string
    	Http basic auth password. Default empty
  -target value
    	List of hosts for ping
  -username string
    	Http basic auth username. Default empty
```

**WARNING**: If username is empty, authorization will be disabled

## Build
```
docker run --rm -v "$PWD":/app -w /app golang:1.16 go build -o pinger
```
