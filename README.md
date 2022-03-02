# Redis Clone
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT) [![Go](https://github.com/ferdigokdemir/in-memory-database/actions/workflows/github-actions-build.yml/badge.svg)](https://github.com/ferdigokdemir/in-memory-database/actions/workflows/github-actions-build.yml)



This project is an In-memory Database. You can save data to memory, and you can get it faster. Written using Golang and i use only standart libs.


## Whick packages used?

| Package              |      Description          |
| :-------------------- | :-----------------------: |
| [log](https://pkg.go.dev/log)        |      Used for server logs          |
| [os](https://pkg.go.dev/os)  |     Used for OS environment variables and io operations.        |
| [encoding/json](https://pkg.go.dev/encoding/json)      | Used for http json responses    |
| [fmt](https://pkg.go.dev/fmt)    |   Used for prints      |
| [net/http](https://pkg.go.dev/net/http)        |   Used for http server        |
| [time](https://pkg.go.dev/time)      |   Used for time operations        |
| [strconv](https://pkg.go.dev/strconv)        |      Used for string conversations          |
| [testing](https://pkg.go.dev/testing)        | Used for unit tests     |


## Installation

```bash
go install
go run .
```

### How to production build

```bash
# Linux, MacOS
go build
./in_memory_database
```

### How to run with Docker

```bash
docker-compose up
```

### How to generate Api Documentation
```bash
npm install apidoc -g
apidoc -i ./ -o apidoc
```


## Usage

### Get Item - <sub><sup>This method for getting item from memory</sub></sup>

```json
// POST /api/v1/getItem
{
    "key": "username"
}
```

### Set Item - <sub><sup>This method for save item to database and memory</sub></sup>

```json
// POST /api/v1/setItem
{
    "key": "username",
    "value": "ferdigokdemir"
}
```

### Flush Items - <sub><sup>This method for remove items from memory and database</sub></sup>

```json
// POST /api/v1/flushItems
{

}
```

## Heroku

Example Live URL: https://in-memory-database.herokuapp.com

[![Deploy](https://www.herokucdn.com/deploy/button.svg)](https://heroku.com/deploy?template=https://github.com/ferdigokdemir/in-memory-database)


## Credits

- [Ferdi Gökdemir](https://github.com/ferdigokdemir)
- [All Contributors](https://github.com/ferdigokdemir/in-memory-database/graphs/contributors)

## License

The MIT License (MIT).
Please see [License File](https://github.com/ferdigokdemir/in-memory-database/blob/main/LICENSE.md) for more information.
