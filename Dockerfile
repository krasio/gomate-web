FROM scratch

ENV GOMATE_PORT 8080

EXPOSE $GOMATE_PORT

COPY gomate-web /
CMD ["/gomate-web"]
