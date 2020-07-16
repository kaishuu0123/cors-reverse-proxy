# CORS Reverse Proxy

Simple reverse proxy for CORS issue.

```plain
... has been blocked by CORS policy: No 'Access-Control-Allow-Origin' header is present on the requested resource. If an opaque response serves your needs, set the request's mode to 'no-cors' to fetch the resource with CORS disabled.
```

## Usage

```shell
docker pull kaishuu0123/cors-reverse-proxy

docker run -it -d \
  -e CORS_REVERSE_PROXY_TARGET_URL=http://example.com \
  -e CORS_REVERSE_PROXY_HOST=0.0.0.0 \
  -p 8181:8081
  --name cors-reverse-proxy \
  kaishuu0123/cors-reverse-proxy
```

or

```shell
go build

./cors-reverse-proxy \
  --target-url http://example.com/ \
  --host 0.0.0.0
  --port 8888
```

## Command Line Options & Environment Variables

| CLI Option(Long) | Shorthand | Environment Variables         | Example               | Default   |
| ---------------- | --------- | ----------------------------- | --------------------- | --------- |
| --target-url     | -t        | CORS_REVERSE_PROXY_TARGET_URL | `http://example.com/`   |           |
| --host           | -h        | CORS_REVERSE_PROXY_HOST       | `0.0.0.0`             | localhost |
| --port           | -p        | CORS_REVERSE_PROXY_PORT       | `8888`                | 8081      |
| --origin         | -o        | CORS_REVERSE_PROXY_ORIGIN     | `http://example.com/` | `*`       |

###### Inspired by: https://github.com/LordotU/local-cors-proxy-go/