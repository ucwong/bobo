# Bobo

## Simplest web server (less than ```70``` lines coding of Golang) but available

Implement the basic features of a web backend with the simplest ways. The orignal only has ```bobo.go``` for simple, it is ok to change or split it as needed

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
```
time ./bench.sh

... ...

real	0m29.865s
user	0m13.490s
sys	0m3.074s
```

#### Storage
https://github.com/dgraph-io/badger

## Customized
### User
#### Register
To register user information in json format
##### Method
```
POST
```
##### URL
```
/user/0x2a2a0667f9cbf4055e48eaf0d5b40304b8822184?msg=aHellox&sig=0xee78eaa27526b412d0e970b85f47c96aa0aa67ed1c06f577ffe712a91284659a0a38529194a53891c84919369e09bf7e08d1655544cb044671461e210ddad1eb00
```
##### Params
```
msg: Plain message
sig: the signatue of msg above
```

##### DATA
```
{...}
```

#### Find
To find user information by address (0xabcd)
##### Method
```
GET
```
##### URL
```
/user/0x2a2a0667f9cbf4055e48eaf0d5b40304b8822184
```
##### Params
```
NULL
```

##### DATA
```
NULL
```
