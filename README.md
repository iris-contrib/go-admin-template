# GoAdmin Template

A fork of [GoAdminGroup/example](https://github.com/GoAdminGroup/example) show how to run go-admin with [Iris](https://github.com/kataras/iris). Click [here](http://www.go-admin.cn/en) to learn more.

![Go-Admin UI](https://iris-go.com/images/about/go-admin.png)

## Get started

If you are Windows user, [go-sqlite-dirver](https://github.com/mattn/go-sqlite3) require to download the gcc to make it work.

Install & run the template with [Iris CLI](https://github.com/kataras/iris-cli):

```sh
$ iris-cli run go-admin
```

<details><summary>details...</summary>

Install & run manually:

```sh
$ git clone https://iris-contrib/go-admin-template.git
$ go run .
```

### Using Docker

```sh
$ git clone https://github.com/iris-contrib/go-admin-template.git
$ cd go-admin-template
$ docker build -t go-admin-template .
$ docker attach $(docker run -p 8080:8080 -it -d go-admin-template /bin/bash -c "cd /go/src/app && GOPROXY=http://goproxy.cn go run .")
```

</details>

<br/>

Navigate to [localhost:8080/admin](http://localhost:8080/admin). You should see your app running.

Username: `admin`
Password: `admin`
