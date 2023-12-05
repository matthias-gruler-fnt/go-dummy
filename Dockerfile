FROM scratch
WORKDIR /app
COPY go-dummy /app/
EXPOSE 8000
ENTRYPOINT ["/app/go-dummy"]