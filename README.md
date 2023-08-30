# product
product service


## contetnts
* [General info](#general-info)
* [Technologies](#technologies)
* [Architecture](#architecture)
* [How to filter](#how-to-filter)
* [Setup](#setup)

## General info
this is a service to create product entity.

## Technologies
* Gin Framework
* MongoDB

## Architecture
* Hexagonal

## How to filter
To use filter use this json array of object in query string
```
[{"field":"field name", "value":"value","operator":"mongoDB operator"}]
```
Also pagination is required in query string
```
{"skip":0,"limit":1}
```

### Operator you allow to use
* $eq
* $ne
* $gt
* $gte
* $lt
* $lte
* $in
* $nin
* $exists
* $type

## Setup
To run this project:
```
$ go mod download
$ go mod tidy
$ go run .
```
