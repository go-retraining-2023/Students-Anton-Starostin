FROM golang:latest AS build
WORKDIR /src
COPY . .
RUN go build -o myapp ./cmd

FROM golang:latest
WORKDIR /app
COPY --from=build /src/myapp .
CMD ["/app/myapp"]
