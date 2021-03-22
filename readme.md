## Process of starting a project in golang

```
go mod init github.com/{project-name}
```

### Run

```
sh ops/run.sh
```

### Build

```
sh ops/build.sh
```

### Run Build

```
export $(grep -v '^#' .env | xargs) && ./bin/server
```

```
* bin
* config
* core
    - {component}
        - repository
        - usecase
        - interactor
        - presenter
* entrypoint
    - graph
- api
* ops
    - cmd
* entity
* infra
    - middleware
    - datastore
* pkg
    - {pkg}
```

### Study reference

https://github.com/caohoangnam/go-clean-architecture => https://hackernoon.com/creating-clean-architecture-using-golang-9h5i3wgr
https://github.com/eminetto/clean-architecture-go-v2 => https://eltonminetto.dev/en/post/2020-07-06-clean-architecture-2years-later/
https://github.com/ridhoperdana/go-clean-arch => https://medium.com/easyread/graphql-delivery-on-golangs-clean-architecture-5c995a17b3a8
https://github.com/manakuro/golang-clean-architecture/tree/master
