# Mechanic Store

## Development

## Run database migrations

* Run `go get bitbucket.org/liamstask/goose/cmd/goose` to install goose
* Run `goose up`

### Run the application

* Install [glide](https://github.com/Masterminds/glide) by following the instructions on the website.
* Run `glide install`
* Run `go build`
* Run `./mechanic-store`

### Run the tests
`make tests`

#### Endpoints

##### Create a Company
POST http://localhost:8080/companies

Payload:
```JSON
{
  "name" : "Michael's company",
  "phone" : "+645050505050",
  "email" : "foo@bar.com"
}
```
