# Gomate Web
Web app using [Gomate](https://github.com/krasio/gomate).

# Usage
Load your data using [gomate-cli](https://github.com/krasio/gomate-cli).

$ Run
```
$ go build
$ ./gomate-web
```

Open [http://localhost:8080/?kind=suburb&q=well](http://localhost:8080/?kind=suburb&q=well).

# Run with Docker
```
$ make run
```

Open [http://localhost:8080/?kind=suburb&q=well](http://localhost:8080/?kind=suburb&q=well).

# Run with Kubernetes (Minikube)
```
$ make minikube
$ echo "$(minikube ip) gomate.test" | sudo tee -a /etc/hosts
```

Open [http://gomate.test/?kind=suburb&q=well](http://gomate.test/?kind=suburb&q=well).
See [https://blog.gopheracademy.com/advent-2017/kubernetes-ready-service/](https://blog.gopheracademy.com/advent-2017/kubernetes-ready-service/) for more.

# License
MIT licence.
