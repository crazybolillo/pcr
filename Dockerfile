FROM golang:1.23-alpine

WORKDIR /app

COPY . .

ENV CGO_ENABLED=0
RUN go build -o /bin/pcr ./cmd

FROM scratch

COPY --from=0 /bin/pcr /bin/pcr

EXPOSE 8100

ENTRYPOINT ["pcr"]