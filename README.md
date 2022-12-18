# k64dec

[![Go Report Card](https://goreportcard.com/badge/github.com/laghoule/k64dec)](https://goreportcard.com/report/github.com/laghoule/k64dec)
[![Coverage](https://sonarcloud.io/api/project_badges/measure?project=laghoule_k64dec&metric=coverage)](https://sonarcloud.io/summary/new_code?id=laghoule_k64dec)
[![Security Rating](https://sonarcloud.io/api/project_badges/measure?project=laghoule_k64dec&metric=security_rating)](https://sonarcloud.io/summary/new_code?id=laghoule_k64dec)
[![Vulnerabilities](https://sonarcloud.io/api/project_badges/measure?project=laghoule_k64dec&metric=vulnerabilities)](https://sonarcloud.io/summary/new_code?id=laghoule_k64dec)

This little tool print the decoded base64 of a Kubernetes secret:

## Cmdline

```text
Usage of k64dec
      --file string   kubernetes secret in json or yaml file
      --version       print version
```

## Usage

Piped from a `kubectl get secrets`:

```bash
kubectl get secrets -n kured -o yaml kured-secret-values | k64dec
secretdata
configuration:
  notifyUrl: slack://kured@feingeiXe9Ze/uuGo4Aedaiph/xae1beizaeQu
```

Via a saved file:

```bash
k64dec --file secret.yaml
secretdata
configuration:
  notifyUrl: slack://kured@feingeiXe9Ze/uuGo4Aedaiph/xae1beizaeQu
```

With docker:

```bash
kubectl get secrets -n kured -o yaml kured-secret-values | docker run -i laghoule/k64dec
```

## Notes

For running unittest, you need to set your timezone to UTC:

```bash
export TZ=UTC
go test ./...
```
