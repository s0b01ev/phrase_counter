# phrase_counter

##  How to run:

- from sources: 

```
$ go run parse.go  path/to/file1 /path/to/file2
```
  or 
```
$ cat path/to/file | go run parse.go
```

- binary 
```
$ ./parse-{{ arch }}  path/to/file1 /path/to/file2
```
  or
```
$ cat path/to/file parse-{{ arch }}
```

 {{ arch }} - 'amd64-linux' and 'arm64' (Mac M1)

 amd64-linux binary was compiled on arm64 system  with Go cross-compiling option:
 $ OOS=linux GOARCH=amd64 go build -o parse-amd64-linux parse.go
 


## Docker

- to build:

```
$ docker build -t phrase_counter ./
```

- to run with stdin as input source:

```
$ cat file.txt | docker run -i  phrase_counter
```

- to run with file(s) as argument(s)

```
$ docker run  phrase_counter "/phrase_counter/samples/file1.txt" "/phrase_counter/samples/fileN.txt"
```



## Performance

2 algorithms were considered:
- consequent: parse.go - datasources processed consequently
- concurent: parse-conc.go  - datasources processed concurently, then results merged

Serial algorithm is more efficient on large volumes ( ~1000 Moby Dicks). Intermediate results merging takes more than a half of total run time.


## TODO (what to improve)

- more tests (spent too much time trying to find most effective algorithm)
- make it more user friendly: "--help", arguments check etc


## Known bugs

- Unicode support on Linux. Development was done on arm64 and code perfectly works with unicode. There are 3 formatting errors in processing unicode-test.txt file
- NOT BUG: all words with punct. symbols inside  without surrounding space(s) are not splitted ie H.M.S, a:b, a-b, don't etc...




