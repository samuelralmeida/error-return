# error custom response

## description

When someone is programming, sometimes he/she has to test the code receiving from a server the response with specific error (status code), header, and/or JSON.

This program creates a local server that returns a custom response according to URL params.

Although the software can return any kind of response, I believe that the most common use is to return error responses, to understand and test how the main software handlers each error.

## how to use

1. clone the repo
2. run the program `go run main.go`
3. do a request with params

### params

- `status`: --> a response status code. Ex: 404, 505, 302
- `msg`:    --> a simple body response menssage
- `json`:   --> a JSON struct to return in response. Ex: `{"fisrtName":"Samuel","age":30,"weight":68.5,"lastName":"Sousa"}`
- header:   --> a pair key-value to set in response hearder. The pair is separate by dash(`-`) and the program accept many headers separete by comman(`,`). Ex: `key-value,key2-value2`

### request example

```
http://localhost:8080/?status=404&json={"name":"John","age":30,"weight":68.5,"lastName":"Sousa"}&header=key-value,key2-value2
```
