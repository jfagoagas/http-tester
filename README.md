# http-tester
Tool to check HTTP responses and Security Headers 

## Usage
```
./http-verbs-test -u <URL> -h <Bearer Token>
```
If you need an Authorization header with a Bearer JWT,
you must input in a `config.yml` like this:
```
token: <JWT>
```