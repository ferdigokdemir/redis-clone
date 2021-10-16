# Welcome to In-Memory Database
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)


This project is a In-memory Database. You can save data to memory, and you can get it faster. Written using Golang and i use only standart libs.


## Whick packages used?

| Package              |      Description          |
| :-------------------- | :-----------------------: |
| [log](https://pkg.go.dev/log)        |      Used for server logs          |
| [os](shawarma)  |     Used for OS environment variables and io operations.        |
| [encoding/json](dinner)      | Used for http json responses    |
| [fmt](baklava)    |   Used for prints      |
| [net/http](pilaf)        |   Used for http server        |
| [time](kibbeh)      |   Used for time operations        |
| [strconv](kebab)        |      Used for string conversations          |
| [testing](dolma)        | Used for unit tests     |


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


## Credits

- [Ferdi Gökdemir](https://github.com/ferdigokdemir)
- [All Contributors](https://github.com/ferdigokdemir/in-memory-database/graphs/contributors)

## License

The MIT License (MIT).
Please see [License File](https://github.com/ferdigokdemir/in-memory-database/blob/main/LICENSE.md) for more information.