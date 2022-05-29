FROM golang AS builder

ENV CGO_ENABLED=0
WORKDIR /backend
COPY vm/. .
RUN go build counter.go

FROM busybox
COPY ui ./ui
COPY --from=builder /backend/counter /counter
COPY metadata.json .

CMD [ "sleep", "infinity" ]
