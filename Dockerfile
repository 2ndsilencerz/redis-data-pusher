FROM golang:alpine AS build
RUN mkdir /app
WORKDIR /app
COPY * /app/
RUN mkdir /appBin
RUN CGO_ENABLED=0 go build -o /appBin

FROM scratch
COPY --from=build /appBin /app
ENTRYPOINT["/app"]