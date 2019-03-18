FROM scratch
EXPOSE 8080
ENTRYPOINT ["/golang-http-cdpr"]
COPY ./bin/ /