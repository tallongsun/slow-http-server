FROM golang:1.11
COPY . /src
WORKDIR /src
RUN go build -o web .

FROM alpine:3.8
RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2
COPY --from=0 /src/web /bin/web
CMD ["web"]
