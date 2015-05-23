# peco-fuzzyfinder

fuzzy matcher for [peco](https://github.com/peco/peco).

## Instllation

```
go get github.com/Sixeight/go-fuzzaldrin
go build peco-fuzzyfinder.go
```

## Configuration

Add below into you `~/.peco/config.json`

```json
"CustomFilter": {
  "FuzzyFinder": {
    "Cmd": "/path/to/peco-fuzzyfinder",
    "Args": [ "$QUERY" ]
  }
}
```

## License

MIT

## Author

Tomohiro Nishimura (a.k.a Sixeight)
