# syntax=docker/dockerfile:1.7-labs

# --------------- STAGE 1 ----------------
FROM golang:1.22.3 AS build

WORKDIR /var/www/Snipetz

# RUN git clone https://github.com/f7ed0/Snipetz.git

COPY backend .
RUN go mod tidy && go build -o api ./api_gateway

RUN ls -lah api

# --------------- STAGE 2 ----------------
FROM debian:latest as prod

WORKDIR /var/www

COPY --from=build /var/www/Snipetz/api /var/www

RUN ls -lah /var/www/api

ENV GIN_MODE=release

EXPOSE 80

CMD ["/var/www/api"]