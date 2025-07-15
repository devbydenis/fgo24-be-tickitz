FROM golang:alpine AS build
WORKDIR /build
COPY . .
RUN go build -o backend-app main.go

FROM alpine:3.22
WORKDIR /app
COPY --from=build /build/backend-app /app/backend
EXPOSE 8080
ENTRYPOINT [ "/app/backend" ]