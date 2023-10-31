FROM golang:latest AS build

WORKDIR /app
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o app

# Final Stage
FROM scratch

WORKDIR /app
COPY --from=build /app/app .

EXPOSE 8080
CMD ["./app"]