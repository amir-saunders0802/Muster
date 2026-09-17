# Use Go only while building the application. This stage is not included in
# the final image, which keeps the production image small.
FROM golang:1.27 AS builder

WORKDIR /src

# Copy the module files first so Docker can reuse the dependency layer when
# application source files change.
COPY go.mod go.sum* ./
RUN go mod download

COPY . .

# Compile Muster into a small Linux executable file that can run without the Go compiler.
# Stripping symbols also reduces the binary size.
RUN CGO_ENABLED=0 GOOS=linux go build -trimpath -ldflags="-s -w" -o /muster-api ./cmd/muster-api


# Start a small Distroless image where the compiled Linux executable will be stored and run.
# Distroless contains the application runtime without a shell or package
# manager, reducing the final image size and attack surface.
FROM gcr.io/distroless/static-debian12:nonroot

WORKDIR /app

# The API loads this file using the relative path config/services.yaml.
COPY --from=builder /muster-api /app/muster-api
COPY --from=builder /src/config/services.yaml /app/config/services.yaml

# Document the port used by the Go HTTP server.
EXPOSE 8080

# Run without root privileges in the runtime container.
USER nonroot:nonroot

# Start the compiled API binary. Distroless provides no shell for commands.
ENTRYPOINT ["/app/muster-api"]
