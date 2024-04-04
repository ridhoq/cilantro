FROM mcr.microsoft.com/oss/go/microsoft/golang:1.22.1-cbl-mariner2.0 as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

RUN go build -o cilantro ./cmd/cilantro

FROM mcr.microsoft.com/cbl-mariner/base/core:2.0

WORKDIR /root/

COPY --from=builder /app/cilantro .

EXPOSE 5050

CMD ["./cilantro"] 