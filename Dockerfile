FROM golang:1.13.9-alpine AS build
WORKDIR /src
COPY . .

RUN apk update && apk add git

RUN go get -d -v

# RUN go test -cover -v

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o bin/cdn .

FROM scratch

LABEL "maintainer"="XTRadio Ops <contact@xtradio.org"
LABEL "version"="0.1"
LABEL "description"="XTRadio CDN"

COPY --from=build /src/bin/xtradio-api /bin/cdn

EXPOSE 10000
EXPOSE 10001

CMD ["/bin/cdn"]
