FROM golang:1.18

WORKDIR /go/src
ENV PATH="/go/bin:${PATH}"
COPY . /go/src/

CMD ["tail", "-f", "/dev/null"]