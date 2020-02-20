# MANGOPAY Go SDK

[![GoDoc](http://img.shields.io/badge/godoc-reference-blue.svg)](http://godoc.org/github.com/58-facettes/mangopay-go-sdk)
[![Build Status](https://travis-ci.org/58-facettes/mangopay-go-sdk.svg?branch=master)](https://travis-ci.org/58-facettes/mangopay-go-sdk)
[![codecov](https://codecov.io/gh/58-facettes/mangopay-go-sdk/branch/master/graph/badge.svg)](https://codecov.io/gh/58-facettes/mangopay-go-sdk)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
![Go: version](https://img.shields.io/github/go-mod/go-version/58-facettes/mangopay-go-sdk)

## Don't use this SDK in Production mode, this current SDK is not fully tested.

MangopaySDK is a Go client library to work with [Mangopay REST API](http://docs.mangopay.com/api-references/).

## Compatibility Notes

* Since v2.1 of this SDK, you must be using at least v2.01 of the API
([more information about the changes required](https://docs.mangopay.com/api-v2-01-overview/))

## Requirements

To use this SDK, you will need (as a minimum):

* Golang v1.13.5 (at least 1.13 is recommended)

## Installation

You can use Mangopay SDK library as a dependency in your project with:

```ssh
go get -u github.com/58-facettes/mangopay-go-sdk
```

## License

MangopaySDK is distributed under MIT license,
see the [LICENSE file](https://github.com/58-facettes/mangopay-go-sdk/blob/master/LICENSE).

## Unit Tests

This project is not yet fully test.
To run test just go into the root directory and run the following command.

```sh
make test
```

## Contacts & Bugs

Report bugs or suggest features using
[issue tracker on GitHub](https://github.com/58-facettes/mangopay-go-sdk/issues).


## Account creation

You can get yourself a [free sandbox account](https://www.mangopay.com/signup/create-sandbox/)
or sign up for a [production account](https://www.mangopay.com/signup/production-account/)
(note that validation of your production account can take a few days,
so think about doing it in advance of when you actually want to go live).

## Configuration

Using the credential info from the signup process above, you should then set `client-id`
to your Mangopay ClientId and `client-password` to your Mangopay APIKey.

`BaseUrl` is set to sandbox environment by default. To enable production
environment, set `mangopay.Config.Mode` to `mangopay.Production`.

```go
package main

import(
    "github.com/58-facettes/magopay-go-sdk"
)

func main() {
    // mangopay.Config.Mode = mangopay.Production // uncomment this to use the production environment
    // mangopay.Config.HTTPClient.Timeout = 3 * time.Second // The cURL response timeout in seconds (its 30 by default)
    api := mangopay.NewWithBasicAuth("your-client-id","your-client-password")
    users, err = api.Users.List(nil)
    if err != nil {
        log.Fatal(err)
    }
    log.Print(users)
}
```

## Sample usage

```go
package main

import (
    "log"

    "github.com/58-facettes/mangopay-go-sdk"
    "github.com/58-facettes/mangopay-go-sdk/model"
)

func main() {
    // mangopay.Config.Mode = mangopay.Production // uncomment this to use the production environment.
    api := mangopay.NewWithBasicAuth("client-id", "client-pass")

    john, err := api.Users.View("someID")
    if err != nil {
        log.Fatal(err)
    }
    log.Print(john)

    // change and update some of his data.
    _, err = api.Users.UpdateNaturalUser("someID", &model.NaturalUserUpdate{
        FirstName: " - CHANGED",
    })
    log.Print(err)

    // get all users (with pagination)
    pagination := model.NewPagination(0, 8)
    users, err := api.Users.List(pagination)
    if err != nil {
        log.Fatal(err)
    }
    log.Print(users)

    // get his bank accounts
    pagination = model.NewPagination(2, 10)
    bankAccount, err := api.BankAccounts.ViewFromUser(john.ID, pagination)
    if err != nil {
        log.Fatal(err)
    }
    log.Print(bankAccount)
}
```

## Sample usage with gin.

You can integrate Mangopay features in a Service in your gin project.

`main.go` :

```go
package main

import (
    "log"
    "net/http"

    "github.com/58-facettes/mangopay-go-sdk"
    "github.com/gin-gonic/gin"
)

var api = mangopay.NewWithBasicAuth("client-id", "client-password")

func main() {
    router := gin.Default()
    router.GET("/users/:userID", HandlerUser)
    router.Run(":8080")
}

func HandlerUser(ctx *gin.Context) {
    jhon, err := api.Users.View(ctx.Param("userID"))
    if err != nil {
        log.Println("error is ", err.Error())
        ctx.AbortWithError(400, err)
        return
    }
    ctx.JSON(http.StatusOK, jhon)
}
```

And run it into the terminal.

```sh
go run main.go
```

## Logging

MangoPay uses the a Logger interface. You can provide your own logger to the API.
Here is a sample showing Monolog integration :

```go
type MyLogger struct{}

func (ml *MyLogger)Debug(args ...interface{})
func (ml *MyLogger)Debugf(template string, args ...interface{})
func (ml *MyLogger)Error(args ...interface{})
func (ml *MyLogger)Errorf(template string, args ...interface{})
func (ml *MyLogger)Fatal(args ...interface{})
func (ml *MyLogger)Fatalf(template string, args ...interface{})
func (ml *MyLogger)Warnf(template string, args ...interface{})

api.Conig.Logger = &MyLogger{};
```

## Verifying rate limits status

According to API docs (https://docs.mangopay.com/guide/rate-limiting), MangoPay is providing a way of 
verifying how many API calls were made, how many are left and when the counter will be reset. 
So there are 4 groups of rate limits available:

1. Last 15 minutes:
2. Last 30 minutes
3. Last 60 minutes
4. Last 24 hours

This information is available from the MangoPayApi instance, like in the following example:

```go
package main

import (
    "github.com/58-facettes/mangopay-go-sdk"
)

func main() {
    api := mangopay.NewWithBasicAuth("your-client-id","your-client-password")
    rl, _ := api.Stats.GetRateLimit()
    log.Println(rl)
}
```

## Idempotency key

You can use an idempotency key with the API, you just have to comply with the interface `data.Manager` in order to save the generated key.

```go
type MyStorage struct{}
func (s *MyStorage)SaveIdempotencyKey(key, url string) error {
    // implement your function
    ...
}
...
// then before connecting to the API
mangopay.Config.UseIdempotency: true
mangopay.Config.DB = myDB
api := mangopay.NewWithBasicAuth("your-client-id","your-client-password")
```
