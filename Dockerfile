FROM dhi.io/golang:1-dev AS build-stage

WORKDIR /api-gateway/

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY ./cmd/ ./cmd/
COPY ./config/ ./config/
COPY ./internal/ ./internal/
COPY ./migrations/ ./migrations/
COPY ./pkg/ ./pkg/

RUN CGO_ENABLED=1 GOOS=linux go build -o ./build/ ./...

FROM dhi.io/golang:1 AS runtime-stage

WORKDIR /api-gateway/

COPY --from=build-stage /api-gateway/build/server ./server
COPY --from=build-stage /api-gateway/build/healthcheck ./healthcheck

CMD ["./server"]
