# Mission CTRL Service

[![Go](https://img.shields.io/badge/go-00ADD8.svg?style=for-the-badge&logo=go&logoColor=white)](https://go.dev/)

Order management microservice.

## Development

To get started, clone this repository and follow theese steps.

Run Docker Containers:

```bash
docker-compose up -d
```

Build Image:

```bash
docker build --progress=plain --no-cache -t spaceops-mission-ctrl:latest -f deployments/local/Dockerfile .
```
