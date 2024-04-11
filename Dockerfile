FROM golang:1.19-alpine3.15 AS build

WORKDIR /app

COPY go.mod ./
RUN go mod download && go mod verify

COPY  . .

RUN go build -o /app/cicd-app

#  Runtime Steps
FROM alpine:3.15

COPY --from=build /app/cicd-app .

EXPOSE 8080

CMD [ "./cicd-app" ]

