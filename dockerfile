FROM golang:1.18.0-alpine3.14 AS build

COPY . /build
WORKDIR /build

RUN apk add --no-cache git
RUN cd cmd && go build -o genealogy-tree

FROM alpine:latest AS final

ARG GENEALOGY_TREE_ADDR=0.0.0.0
ARG CODE_ENV=dev
ARG DB_ADDR=0.0.0.0
ARG DB_PWD=test

ENV GENEALOGY_TREE_ADDR=${GENEALOGY_TREE_ADDR}
ENV CODE_ENV=${CODE_ENV}
ENV DB_ADDR=${DB_ADDR}
ENV DB_PWD=${DB_PWD}

COPY --from=build /build/cmd/genealogy-tree ./

CMD ["./genealogy-tree"]