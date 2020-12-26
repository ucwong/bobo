# Bobo

## Simplest web server (less than 80 lines coding) but available

Implement the basic features of a web backend with the simplest ways. The orignal only has ```bobo.go``` for simple, it is ok to changed or split it as needed

Feel free to fork it and add advanced logics and features

## How to run it ?
```
go run bobo.go
```
or 
```
go build bobo.go
./bobo
```
## How to use it ?
http://localhost:8080/get?k=1

or
```
curl http://localhost:8080/get?k=1
```
Get some value with key=1

http://localhost:8080/set?k=1&v=2

or 
```
curl http://localhost:8080/set?k=1\&v=2
```

Set invoke, set!

## How to test it ?
```
time ./bench.sh

... ...

real	0m29.865s
user	0m13.490s
sys	0m3.074s
```

#### Storage
https://github.com/dgraph-io/badger
