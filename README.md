# CC-CEDICT as a microservice

This project uses the open source Chinese-English dictionary to expose an API for translation of chinese characters. The script dictionary_bootstrapper.sh automatically downloads the latest dictionary and executes a go program to parse the dictionary and store it into a BoltDB file.

## Getting Started

### Building Dockerfile

```
docker build -t cndict .
```

### Runing Docker container

```
docker run -p 3141:3141 cndict
```

### Consuming the API

```
curl -s localhost:3141/lookupCharacters/电脑业者 | jq
```
```
{
  "Traditional": "電腦業者",
  "Pinyin": "dian4 nao3 ye4 zhe3",
  "Translations": [
    "software developer"
  ]
}

```


## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details