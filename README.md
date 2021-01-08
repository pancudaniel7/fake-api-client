# fake-api-client

- fake-api-client is a dependency library used for interacting with rest fake-api using Golang. 

### Get started

- In order to get started be sure that you have github **ssh key setup** ([link](https://docs.github.com/en/free-pro-team@latest/github/authenticating-to-github/generating-a-new-ssh-key-and-adding-it-to-the-ssh-agent)), 
then in order to import library packages set **GOPROXY** and **GOPRIVATE** as following:
``` sh
    go env -w GOPROXY=direct GOPRIVATE=github.com/pancudaniel7
```
- After that you can add the dependency to your **go.mod** file like this:
```go
    //go.mod

    require github.com/pancudaniel7/fake-api-client v1.0.0
```
- Then you can access api **Resource** and make any operations like **create** using the object:
```go
    	acc := api.Account{
                ID:             "9095b5f1-aa8d-40e2-8442-f4da2f576b4e",
                CreatedOn:      time.Time{},
                ModifiedOn:     time.Time{},
                OrganisationID: "afa14cf7-b28e-421f-b3af-6a47b4121fc2",
                Type:           "Accounts",
                Version:        0,
                Attributes: api.Attributes{
                AccountNumber:               "10000004",
                AccountClassification:       "Personal",
                AccountMatchingOptOut:       false,
                AlternativeBankAccountNames: nil,
                BankID:                      "400302",
                BaseCurrency:                "GBP",
                Bic:                         "BARCGB22XXX",
                Country:                     "GB",
                CustomerID:                  "987788",
                JointAccount:                false,
                Iban:                        "GB33BUKB20201555555555",}}

        res, err := acc.Create()
        if err != nil {
            log.Fatalf("Fail to create account: %s", err)
        }

        log.Printf("Successful created account: %s", res)
```

### Environment variables

Name | Default Value | Description 
--------- | --------- | --------- |
BASE_API_URL  | http://localhost:8080/v1  | Default base api url |
HTTP_CLIENT_REQ_TIME_OUT  | 1 minute | Request time out value in nanoseconds |
HTTP_RECORD_VERSION  | 0 | Api record version |
HTTP_DEFAULT_PAGE_SIZE  | 2 | List resources page size |

### Testing

- In order to see the library tests you can run inside the main project folder the command:

```sh
    docker-compose up
```
- Or you can start **fake-api** services using docker and run separately tests using **go test** command:

```sh
  docker-compose up -d && go test -v --tags=integration ./...
```

### Api operations

- Api operations are described by **Resource** interface that is the parent of every api resource entity:

```go
  //resource.go

  type Resource interface {
      Create() (Resource, error)
      List(pageNum, pageSize string) ([]Resource, error)
      ListById() (Resource, error)
      Delete() error
  }
```

- Also, client library supports **fetch** operations using **ResourcePromise** as in the example below:
```go
    expAcc := api.Account{...}
    p := api.NewResourcePromise(expAcc.Create)

    p.Then(func(res api.Resource) {
        // successful promise function code...
    }

    p.Cache(func(err error) {
        // fail promise function code...
    })
```

***
**Name: Pancu Daniel**<br>
**I am new to Go**
