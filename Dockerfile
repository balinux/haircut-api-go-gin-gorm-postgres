FROM docker.io/library/golang:1.22-alpine AS build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o haircut-app .

# add stage for scratch image
FROM scratch
COPY --from=build /app/haircut-app /haircut-app
EXPOSE 8888
CMD [ "/haircut-app" ]
