FROM dhi.io/golang:1-dev AS build-stage

WORKDIR /payment-processing/

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY ./cmd/ ./cmd/
COPY ./config/ ./config/
COPY ./internal/ ./internal/
COPY ./pkg/ ./pkg/

RUN CGO_ENABLED=1 GOOS=linux go build -o ./build/ ./...

FROM dhi.io/golang:1 AS runtime-stage

WORKDIR /payment-processing/

COPY --from=build-stage /payment-processing/build/server ./server
COPY --from=build-stage /payment-processing/build/healthcheck ./healthcheck

CMD ["./server"]
