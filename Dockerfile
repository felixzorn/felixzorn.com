# Build
FROM golang:1.24-alpine AS build
WORKDIR /src
COPY . .
RUN CGO_ENABLED=0 go build -trimpath -ldflags="-s -w" -o /deinname-dev .

# Run: nur das Binary, kein Betriebssystem drumherum
FROM scratch
COPY --from=build /deinname-dev /deinname-dev
ENV ADDR=:8080
EXPOSE 8080
ENTRYPOINT ["/deinname-dev"]
