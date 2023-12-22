FROM golang:1.21.4-alpine3.18 AS build-stage
WORKDIR /home/app/
COPY ./ /home/app/
RUN mkdir -p /home/b
RUN go mod download
RUN go build -v -o /home/build/api ./cmd/api
    
FROM gcr.io/distroless/static-debian11

COPY --from=build-stage /home/build/api /api
CMD ["/api"]