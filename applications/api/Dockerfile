FROM golang:1.25-alpine

ARG user=gopher
ARG uid=1000

RUN apk add --no-cache git

RUN go install github.com/air-verse/air@latest && \
    mv /go/bin/air /usr/bin/air

RUN adduser -D -u $uid -h /home/$user $user

WORKDIR /app

RUN chown -R $user:$user /app
RUN mkdir -p /go && chown -R $user:$user /go

USER $user

CMD ["air", "-c", ".air.toml"]