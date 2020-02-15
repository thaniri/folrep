# Stage 1 includes all source code
FROM "golang:alpine" as stage-1
WORKDIR /go/src
COPY . /go/src
RUN cd /go/src && go build -o webApp cmd/main.go

# Stage 2 only executes the binary
FROM alpine
WORKDIR /go/bin
COPY --from=stage-1 /go/src/webApp /go/bin/
EXPOSE 8080
ENTRYPOINT "/go/bin/webApp"
