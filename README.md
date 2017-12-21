# Gomate Web
Web app using [Gomate](https://github.com/krasio/gomate).

# Usage
Load your data using [gomate-cli](https://github.com/krasio/gomate-cli).

# Build
```
$ make build
```
You can provide RELEASE build option (defaults to 0.0.1).

# Run
You can provide GOMATE_PORT (defaults to 8080) and GOMATE_REDIS_URL (defaults to redis://localhost:9999/0).
```
$ ./gomate-web 
2017/12/21 14:02:40 [GOMATE] commit: 53834b2, build time: 2017-12-21_01:02:37, release: 0.0.1.
```

Open [http://localhost:8080/?kind=suburb&q=well](http://localhost:8080/?kind=suburb&q=well).

# License
MIT licence.
