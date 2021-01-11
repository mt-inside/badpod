FROM golang:1.15 as build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

# including the .git dir, for version stuff to work
COPY . .
RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go install -a -tags netgo -ldflags "-w -extldflags '-static' $(build/ldflags.sh)" cmd/badpod.go


FROM gcr.io/distroless/static-debian10:latest AS run

ARG PORT=8080

COPY --from=build /go/bin/badpod /
COPY data data

EXPOSE $PORT
ENTRYPOINT ["/badpod"]
