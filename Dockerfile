
FROM golang:nanoserver AS build

COPY . /code
WORKDIR /code

RUN go build web1.go

FROM microsoft/nanoserver

COPY --from=build /code/web1.exe /web1.exe

EXPOSE 8080

CMD ["\\http.exe"]
