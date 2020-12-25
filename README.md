# bobo
## Simplest webserver but available

#### To implement all the funcs of a webserver production with the simplest ways, you can change the parts as needed

#### Get, Set and Persistence

```
go run main.go
```
http://localhost:8080/get?k=1

Get some value with key=1

http://localhost:8080/set?k=1&v=2

Set invoke, set!

#### Use badger as the storage
https://github.com/dgraph-io/badger
