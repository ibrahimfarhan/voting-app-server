FROM golang:1.19-buster
WORKDIR /
COPY . .
EXPOSE 8080
CMD [ "go", "run", "main/main.go" ]