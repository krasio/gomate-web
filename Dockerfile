FROM scratch

ENV GOMATE_PORT 8080
ENV GOMATE_REDIS_URL redis://gomate-redis:6379/0

EXPOSE $GOMATE_PORT

COPY gomate-web /
CMD ["/gomate-web"]
