FROM scratch
COPY mapserver /bin/mapserver
ENV MT_CONFIG_PATH "mapserver.json"
ENV MT_LOGLEVEL "INFO"
ENV MT_READONLY "false"
EXPOSE 8080
ENTRYPOINT ["/bin/mapserver"]