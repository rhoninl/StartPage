FROM --platform=linux/amd64 golang:1.18-alpine as builder

WORKDIR /startpage
COPY . .

ARG TARGETOS
ARG TARGETARCH

RUN go mod download

RUN CGO_ENABLED=0 GOOS=$TARGETOS GOARCH=$TARGETARCH go build -a -o /output/startpage main.go 

FROM gcr.io/distroless/static-debian11

COPY --from=builder /output/startpage  /startpage

ENTRYPOINT [ "/startpage" ]