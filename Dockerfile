FROM golang:1.19-buster AS build
WORKDIR /
COPY . .
RUN sh build.sh
FROM scratch AS bin
WORKDIR /
EXPOSE 8080
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /build/bin/voting-app /app/voting-app
CMD [ "/app/voting-app" ]