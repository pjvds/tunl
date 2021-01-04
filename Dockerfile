FROM gcr.io/distroless/base
COPY tunl tunl

ENTRYPOINT ["./tunl"]
