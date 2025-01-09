# Build Stage
FROM golang:1.23.4 AS build
WORKDIR /app
COPY ./go.mod .
COPY ./go.sum .
RUN go mod tidy
RUN go mod vendor
COPY . .
RUN go build -o app.exe ./cmd

# Production Stage
FROM golang:1.23.4 AS production
WORKDIR /app
COPY --from=build /app/.env .
COPY --from=build /app/app.exe .
EXPOSE 3000
CMD [ "./app.exe" ]
