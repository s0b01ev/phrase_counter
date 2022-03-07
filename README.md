# phrase_counter


** Docker

- to build:
$ docker build -t phrase_counter ./

- to run with stdin as input source:
$ cat file.txt | docker run -i  phrase_counter

- to run with file(s) as argument(s)
$ docker run  phrase_counter "/phrase_counter/samples/file1.txt /phrase_counter/samples/fileN.txt"
