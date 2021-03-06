![](okra-logo.png)
=======================

An OKR app to manage and facilitate your company's objectives and key results between team members.

##Dependencies
- [gin](github.com/gin-gonic/gin)
- [mgo](gopkg.in/mgo.v2)
- [mgo/bson](gopkg.in/mgo.v2/bson)
- [yaml](gopkg.in/yaml.v1)
- [cli](github.com/codegangsta/cli)

## Data Stores
- MongoDB

## Installation/Set up

1. Install MongoDB.

2. In your project directory, run

``` go
go get github.com/dmonay/okra
```


## Tests and Benchmarks


*NOTE*: No tests as of yet. To be added. 

To run tests, run 

```go
go test
```



To run benchmarks, run

```go 
go test -check.b
```

## Documentation

1. Start the server: 

```go 
go run server.go server
```

### Create your organization

    POST /create/organization
    

**Sample Request Body**

*NOTE*: Must pass in a parameter for the organization.

```json
{
"organization":"DopeStartup"
}
```


**Sample Response Body**

```json
{
   "You have successfully created an organization"
}
```


### Add an OKR tree to the organization

    POST /create/tree/:organization
    

**Sample Request Body**

*NOTE*: Must pass in a parameter for the timeframe. One of "annual" or "monthly".

```json
{
"timeframe":"weekly"
}
```


**Sample Response Body**

```json
{
   "You have successfully created a tree with the annual timeframe for the DopeStartup organization"
}
```

### Add or update mission in your tree

    POST /update/mission/:organization
    

**Sample Request Body**

*NOTE*: Must pass in a parameter for the mission.

```json
{
"mission":"To thrive amidst enemies"
}
```


**Sample Response Body**

```json
{
   "You have successfully added a mission"
}
```
 	