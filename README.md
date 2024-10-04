# Well Known
The easiest way to host your well known files in kubernetes.

## Config
Place your config file at `/etc/well-known/config.yaml`


```yaml
files:
  apple-app-site-association:
    content: |
      {
        "key": "value"
        ...
      }
    headers:
      Content-Type: application/json
  apple-developer-merchantid-domain-association:
    content: "..."
    headers:
      Content-Type: text/plain
  assetlinks.json:
    content:  |
      {
        "key": "value"
      }
    headers:
      Content-Type: application/json
      X-Custom-Header: value
```