# Microservices Monorepo (Starter)

This repo contains three simple Go microservices using Gin and a starter CI/CD pipeline.

## How to run locally

1. In each service folder run `go mod init` (already present in repo) and `go build`.
2. From `deployments/` run `docker compose up --build`.

## GitHub Actions

- The workflow builds each service image and pushes to GHCR.
- After build it SSHs into EC2 and runs `docker compose pull && docker compose up -d`.

## Points

- Containerized services with Docker and Docker Compose.
- CI/CD with GitHub Actions, image registry (GHCR), automated deploy via SSH.
- Health checks, graceful shutdown, separation of concerns across services.
