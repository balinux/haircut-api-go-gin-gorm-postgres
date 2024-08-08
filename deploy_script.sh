#!/bin/bash

# Pull the latest image
podman pull docker.io/balinux/haircut:postgres

# Stop the currently running container (if any)
podman stop myapp || true
podman rm myapp || true

# Run the new container
podman run -d --name myapp -p 8886:8888 docker.io/balinux/haircut:postgres
