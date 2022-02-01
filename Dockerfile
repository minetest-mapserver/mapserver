FROM scratch
COPY mapserver /bin/mapserver
EXPOSE 8080
ENTRYPOINT ["/bin/mapserver"]
