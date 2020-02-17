# MANGOPAY Go SDK  [![Build Status](https://travis-ci.org/58-facettes/mangopay-go-sdk.svg?branch=master)](https://travis-ci.org/58-facettes/mangopay2-go-sdk) [![License](https://poser.pugx.org/mangopay/go-sdk-v2/license)](https://packagist.org/packages/mangopay/go-sdk-v2)

## Please don't use this SDK yet this is steel under heavy developpement

MangopaySDK is a Go client library to work with [Mangopay REST API](http://docs.mangopay.com/api-references/).

## Compatibility Notes

* Since v2.1 of this SDK, you must be using at least v2.01 of the API
([more information about the changes required](https://docs.mangopay.com/api-v2-01-overview/))
* If you experience problems with authentification and/or the temporary token file following an SDK update
(particuarly updating to v2.0 of the SDK), you may need to just delete your temporary file
(that you specify with `api.TemporaryFolder`) - which allows it to be regenerated correctly the next time it's needed

## Requirements

To use this SDK, you will need (as a minimum):

* Golang v1.13.5 (at least 1.13 is recommended)

## Installation

You can use Mangopay SDK library as a dependency in your project with 

`go get -u github.com/mangopay/mangopay2-go-sdk`

(which is the preferred technique). Follow [these installation instructions]

## License

MangopaySDK is distributed under MIT license,
see the [LICENSE file](https://github.com/Mangopay/mangopay2-go-sdk/blob/master/LICENSE).

## Unit Tests

This project is not yet fully test.
To run test just go into the root directory and run the following command.

```sh
$ make test
```

## Contacts & Bugs

Report bugs or suggest features using
[issue tracker on GitHub](https://github.com/mangopay/mangopay2-go-sdk/issues).


## Account creation

You can get yourself a [free sandbox account](https://www.mangopay.com/signup/create-sandbox/)
or sign up for a [production account](https://www.mangopay.com/signup/production-account/)
(note that validation of your production account can take a few days,
so think about doing it in advance of when you actually want to go live).

## Configuration

Using the credential info from the signup process above, you should then set `client-id`
to your Mangopay ClientId and `client-password` to your Mangopay APIKey.

You also need to set a folder path in `api.TemporaryFolder` that SDK needs 
to store temporary files. This path should be outside your www folder.
It could be `/tmp/` or `/var/tmp/` or any other location that Go can write to. 
**You must use different folders for your sandbox and production environments.**

`BaseUrl` is set to sandbox environment by default. To enable production
environment, set `mangopay.Mode` to `mangopay.Production`.

```go
package main

import(
    "github.com/mangopay/magopay2-go-sdk"
)

func main() {
    // mangopay.Mode = mangopay.Production // uncomment this to use the production environment
    api := mangopay.NewWithBasicAuth("your-client-id","your-client-password")
    api.TempPath = "/some/path/"
    // api.ResponseTimeout = 20 // The cURL response timeout in seconds (its 30 by default)
    // api.ConnectionTimeout = 60;//The cURL connection timeout in seconds (its 80 by default)
    // api.CertificatesFilePath = "" //Absolute path to file holding one or more certificates
    // to verify the peer with (if empty, there won't be any verification of the peer's certificate).
    users, err = api.Users.GetAll()
    if err != nil {
        log.Fatal(err)
    }

    log.Print(users)
}
```

## Sample usage

```go
package main

import(
    "github.com/mangopay/magopay2-go-sdk"
)

func main() {
    // mangopay.Mode = mangopay.Production // uncomment this to use the production environment.
    api := mangopay.NewWithBasicAuth("your-client-id","your-client-password")
    api.TempPath = "/some/path/"

    john, err := api.Users.Get("someID")
    if err != nil {
        log.Fatal(err)
    }
    log.Print(john)

    // change and update some of his data.
    _, err = api.Users.Update("someID", &model.UserUpdate{
        LastName: " - CHANGED",
    })
    log.Print(err)

    // get all users (with pagination)
    pagination := model.NewPagination(1, 8)
    users, err := api.Users.GetAll(pagination)
    if err != nil {
        log.Fatal(err)
    }
    log.Print(users)

    // get his bank accounts
    pagination = model.NewPagination(2, 10)
    users, err := api.Users.GetBankAccounts(john.ID, pagination)
    if err != nil {
        log.Fatal(err)
    }
    log.Print(users)

}
```

## Sample usage with gin.

You can integrate Mangopay features in a Service in your gin project.

`main.go` :

```go

package main

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/mangopay/magopay2-go-sdk"
)

var api = mangopay.NewWithBasicAuth("your-client-id","your-client-password")

func main() {
    api.TempPath = "/some/path/"
    router := gin.Default()
    router.GET("/users/:userID", HandlerUser)
    router.Run(":8080")
}

func HandlerUser(ctx *gin.Context) {
    john, err := api.Users.Get(ctx.Param("userID"))
    if err != nil {
        ctx.AbortWithError(400, err)
        return
    }
    c.JSON(http.StatusOK, jhon)
}
```

And run it into the terminal.

```sh
$ go run main.go
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

api.Logger = &MyLogger{};
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
    "github.com/Mangopay/mangopay2-go-sdk"
)

func main() {
    api := mangopay.NewWithBasicAuth("your-client-id","your-client-password")
    log.Println(api.RateLimits(model.RateLimit15Minutes))
}
```
