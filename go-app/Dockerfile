FROM ubuntu AS build
RUN apt-get update && apt-get install -y golang-go
ENV GO111MODULE=off
COPY . .
RUN CGO_ENABLE=0 go build -o /app .
#stage2
FROM scratch
COPY --from=build /app /app
ENTRYPOINT ["/app"]
