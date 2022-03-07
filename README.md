# phrase_counter


** Docker

- to build:
$ docker build -t phrase_counter ./

- to run with stdin as input source:
$ cat file.txt | docker run -i  phrase_counter

- to run with file(s) as argument(s)
$ docker run  phrase_counter "file1.txt file2.txt"
