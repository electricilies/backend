FROM golang:1.24
WORKDIR /app
COPY . .
RUN go build -o /backend ./cmd/main.go

FROM scratch
COPY --from=0 /backend /backend
EXPOSE 8080
CMD ["/backend"]

