FROM golang:alpine
LABEL maintainer="Amod Amatya"

RUN mkdir /src
COPY ./main /src

WORKDIR /src

RUN go build -o main .

EXPOSE 8080

CMD [ "/src/main" ]