FROM node:11.13.0-alpine@sha256:5fec7c5da14d7ce1e27247cce8889f51d2cf97f6aa73511ccc9a8944f066d625 AS client

WORKDIR /app
COPY client/package.json client/yarn.lock ./
RUN yarn install
COPY client/ ./
RUN yarn build


FROM golang:1.12.2-alpine@sha256:d481168873b7516b9f34d322615d589fafb166ff5fd57d93e96f64787a58887c AS server

WORKDIR /app
RUN apk add --no-cache git
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -a -o 50x50 .


FROM scratch
LABEL maintainer="meneer@koenbollen.nl"

EXPOSE 8080
ENTRYPOINT ["/50x50"]

COPY --from=client /app/dist/ /client/dist/
COPY --from=server /app/50x50 /50x50
