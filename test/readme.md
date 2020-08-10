- Clone Project in GOPATH directory,

`git clone https://github.com/mayadata-io/kubera-api-testing.git`

- get into code path `cd /kubera-api-testing/test`
- run below command
    - `go get github.com/yalp/jsonpath`
    - `go build`
- to Start server need User_Credentials  and HostIP of Kubera .
```
./test -username=<Kubera User_name> -password=<Kubera Password> -host=<KuberaAccessIP>
```