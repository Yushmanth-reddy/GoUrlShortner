FROM golang:alpline as builder

RUN mkdir /builder

ADD . /build/

WORKDIR /build

RUN go build -o main.

FROM alpline

RUN adduser -S -D -H /app appuser

USER appuser

COPY . /app

COPY --from=builder /build/main /app/

WORKDIR /app

EXPOSE 3000

CMD [ "./main" ]
