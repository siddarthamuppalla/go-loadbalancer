# Go-LoadBalancer

Basic load balancer in Golang.

A load balancer is an application that helps route traffic to various backend servers and manages load on those servers.

Currently, this application follows Round robin strategy to distribute load equally across the available backend servers in the `config.yml` file.

```yaml
# provide the list of srever URLs 
servers:
  - "https://www.google.com"
  - "https://www.bing.com"
  - "https://siddartha-muppalla.netlify.app"

```

To use the application, execute the following in the root directory:
```bash
go run .
```