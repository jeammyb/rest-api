# REST API Sample in GO

Simple Golang REST API example that provides a JSON with details about countries.
It consumes the data from the API provided by https://restcountries.eu.

## Run application

Go to the main folder and execute the following commands:
```
go build
./rest-api
```

## Build docker image

`make build`


## Run Docker container

`make run`


## Example
Format: `http://localhost:8080/api/country/{country_name}`

`http://localhost:8080/api/country/guatemala`

Output:
```
[
    {
    name: "Guatemala",
    capital: "Guatemala City",
    region: "Americas",
    subregion: "Central America",
    population: 16176133
    }
]
```
