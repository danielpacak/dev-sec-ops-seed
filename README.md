[![Release][release-img]][release]
[![Build Actions][build-action-img]][build-action]
[![Codecov][codecov-img]][codecov]

# dev-sec-ops-seed

A seed project for configuring DevSecOps on GitHub.

The sample application is an HTTP service exposing a very simple RESTful API. Built with [the 12-factor app][twelve-factor]
principles in mind it's configurable with environment variables. It can be compiled to a statically linked binary or
built into a Docker container.

One of the endpoints implemented by the application accessible at `/api/info` returns build info such as Git commit hash,
tag, and build date. The same information is logged on application startup to the standard output. This is useful when
checking which version of the application is currently running in your cluster.

You can `make container-run` to compile the Golang code, build a Docker image, and run the container locally.
If everything goes well you should be able to query the API, for example, with [curl][curl].

```
$ curl localhost:8080/api/info
{"version":"dev","commit":"none","date":"unknown"}
```

Alternatively, you could run the released version of the sample application as follows.

```
$ docker run --rm --name seed -p 8080:8080 \
  docker.pkg.github.com/<github id>/dev-sec-ops-seed/seed:<release tag>
```

Similarly, you can use the API with curl.

```
$ curl localhost:8080/api/info
{"version":"0.0.5","commit":"2074f8fb1156f32c1a4adda6e297bb0ff3c2c08f","date":"2019-10-16T08:24:14Z"}
```

Note that whenever you run a released version of the app, the `/api/info` response returns the exact Git commit hash,
and Git reference as well as build datetime.

[release-img]: https://img.shields.io/github/release/danielpacak/dev-sec-ops-seed.svg
[release]: https://github.com/danielpacak/dev-sec-ops-seed/releases
[build-action-img]: https://github.com/danielpacak/dev-sec-ops-seed/workflows/build/badge.svg
[build-action]: https://github.com/danielpacak/dev-sec-ops-seed/actions
[codecov-img]: https://codecov.io/gh/danielpacak/dev-sec-ops-seed/branch/master/graph/badge.svg
[codecov]: https://codecov.io/gh/danielpacak/dev-sec-ops-seed
[twelve-factor]: https://12factor.net
[curl]: https://curl.haxx.se
