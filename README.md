# bobo

## Simplest webserver (less than 100 lines coding) but available

To implement all the funcs of a webserver production with the simplest ways, you can change the parts as needed. The orignal only has ```bobo.go``` for simple, it is ok to split it as needed

### (Feel free to fork it and add advanced logics and features)

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

#### Storage
https://github.com/dgraph-io/badger

or

https://github.com/syndtr/goleveldb
