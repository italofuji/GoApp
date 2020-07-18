FROM scratch

WORKDIR /go/app

ADD sum .

CMD ["./sum"]