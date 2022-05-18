# go-clean-architecture

# Requirement

* Go 1.18.*
* gin 1.7.* (web framework)
* goose (db migration)
* gorm v2 1.22.* (orm)

# Usage

## development

# Architecture

# Code Structure

<details>
    <summary>ディレクトリ構成</summary>

```
.
├── db
├── image
├── src
│   ├── config
│   ├── domain
│   ├── gateway
│   ├── infrastructure
│   ├── interfaces
│   ├── middleware
│   ├── registry
│   ├── testutil
│   ├── usecase
│   └── util
├── Dockerfile     //productuin用
├── Dockerfile.dev //開発用
├── READNE.md
├── entrypoint.sh  //Dockerfileの起点
├── go.mod
├── go.sum
└── main.go          //起点
```

</details>

# Author

# License