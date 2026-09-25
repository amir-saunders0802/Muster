# Muster

Muster is a small Go HTTP API that loads a service catalog from YAML and makes
the catalog available through a few simple endpoints.

The application separates catalog loading from the HTTP layer and uses Go's
standard `net/http` package for routing and structured `log/slog` logging.

## Requirements

- Go 1.27 or later
- Docker, if running the container

## Run locally

From the project root, run:

```bash
go run ./cmd/muster-api
```

The API listens on port `8080` and expects the catalog at
`config/services.yaml`.

## Endpoints

### Health check

```bash
curl http://localhost:8080/healthz
```

Returns:

```json
{ "status": "ok" }
```

### List all services

```bash
curl http://localhost:8080/services
```

### Find a service by name

```bash
curl http://localhost:8080/services/payments-api
```

If the service does not exist, the API returns a `404` response.

## Run with Docker

Build the image:

```bash
docker build -t muster:local .
```

Run the container locally:

```bash
docker run --rm -p 8080:8080 muster:local
```

For Kind, load the local image before applying the manifests:

```bash
kind load docker-image muster:local --name <your-cluster-name>
```

The application is built using a multi-stage Docker build. The final runtime
image uses Distroless, contains no shell or Go build tooling, and runs as a
non-root user.

## Kubernetes note

<!-- subPath mounts do not refresh in place, and the app loads the catalog once at startup. If the ConfigMap is changed, the pod keeps serving the old catalog until it is restarted. -->

When the catalog ConfigMap changes, restart the deployment so the new catalog is loaded:

```bash
kubectl rollout restart deployment/muster-api
```

## Tests

Run all tests with:

```bash
go test ./...
```
