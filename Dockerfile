FROM gcr.io/distroless/static-debian12
COPY dist/server /app/server
ENTRYPOINT ["/app/server"]