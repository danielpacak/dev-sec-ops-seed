[![Release][release-img]][release]
[![Build Actions][build-action-img]][build-action]
[![Codecov][codecov-img]][codecov]

# dev-sec-ops-seed

A seed project for configuring DevSecOps on GitHub.

## Prerequisites

Note that GitHub Actions are not triggered on forked repositories. Hence you must set up your own project based on this
repository. Make sure to add the following secrets to the project's settings:

- `GITHUB_REGISTRY_USER` – your GitHub identifier
- `GITHUB_REGISTRY_TOKEN` – your personal GitHub access token with scope for publishing releases and packages
- `CODECOV_TOKEN` – Codecov repository token if you want to publish code coverage reports to [Codecov][codecov]

## Getting started

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

## Deployment

### Kubernetes

When deploying an application, whose container images reside in a private registry such as GitHub Docker Registry,
Kubernetes needs to know the credentials required to pull the image.

To run a deployment, which uses an image from the private repository, you need to do two things:

1. Create a secret holding the credentials for the registry
2. Reference the secret in the `imagePullSecrets` field of the Deployment manifest

The following command will create a secret holding GitHub Package Registry credentials.

```
$ kubectl create secret docker-registry github-docker-registry \
  --docker-server="https://docker.pkg.github.com/v2/" \
  --docker-username=$GITHUB_USER \
  --docker-password=$GITHUB_TOKEN
```

If you inspect the contents of the newly created Secret with `kubectl` describe, you'll see that it includes a single
entry `.dockerconfigjson`. This is equivalent to the `~/.docker/config.json` file in your home directory, which is
created by Docker when you run the docker login command.

To have Kubernetes use the Secret when pulling images from your private GitHub Package Repository, all you need to do
is specify the Secret's name in the Deployment spec, as shown in the following listing.

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: seed
spec:
  replicas: 3
  selector:
    matchLabels:
      app: seed
  template:
    metadata:
      labels:
        app: seed
    spec:
      imagePullSecrets:
        - name: github-docker-registry
      containers:
        - name: main
          image: docker.pkg.github.com/<github id>/dev-sec-ops-seed/seed:0.0.5
          imagePullPolicy: IfNotPresent
          env:
            - name: "SEED_HTTP_ADDR"
              value: ":8080"
          ports:
            - name: http-port
              containerPort: 8080
```

You create the seed Deployment with the following command.

```
$ kubectl create –f kube/seed.yaml --record
```

Notice the `--record` flag passed to `kubectl` create command which is used in case you have to check the history of
releases or rollback a release.

You could now perform a rolling update of the seed deployment from the current version 0.0.5 to 0.0.6 by updating its image.

```
$ kubect set image deployment seed \
  docker.pkg.github.com/<your github id>/dev-sec-ops-seeed/seed:0.0.6
```

To test that the deployment has been updated to version 0.0.6 you can run:

```
$ kubect port-forwar servcie/seed 8080:8080 &> /dev/null &
$ curl http://localhost:8080/api/info | jq .version
0.0.6
```

[release-img]: https://img.shields.io/github/release/danielpacak/dev-sec-ops-seed.svg
[release]: https://github.com/danielpacak/dev-sec-ops-seed/releases
[build-action-img]: https://github.com/danielpacak/dev-sec-ops-seed/workflows/build/badge.svg
[build-action]: https://github.com/danielpacak/dev-sec-ops-seed/actions
[codecov-img]: https://codecov.io/gh/danielpacak/dev-sec-ops-seed/branch/master/graph/badge.svg
[codecov]: https://codecov.io/gh/danielpacak/dev-sec-ops-seed
[twelve-factor]: https://12factor.net
[curl]: https://curl.haxx.se
[codecov]: https://codecov.io/