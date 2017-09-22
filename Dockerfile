FROM alpine
ADD simpleservice /simpleservice
RUN chmod +x /simpleservice
ENTRYPOINT ["/simpleservice"]