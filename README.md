<img src="./logo.svg" style="width: 100%;">

# Repray

A simple reverse proxy tool

## Features

- HTTPS
- h2c/gRPC
- Rejecting TLS handshakes when receiving unknown hostnames

## Usage

```shellsession
$ repray <config_file>
```

## Config

### Type

```typescript
interface Config {
  listen: string
  upstream: string
  tls: {
    certFile: string
    keyFile: string
  } | null
}
```

### Example

```json
[
  {
    "listen": ":8443",
    "upstream": "h2c://127.0.0.1:3000",
    "tls": {
      "certFile": "foo.com.pem",
      "keyFile": "foo.com.key"
    }
  },
  {
    "listen": "127.0.0.1:8080",
    "upstream": "https://example.com"
  }
]
```
