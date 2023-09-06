# AKGo
Easy and simple Go framework for build REST API

AKGo is a framework equipped with exception handler, log filter, env, and MDC. 
AKGo also has a consistent response template, making it very easy to use.

Endpoint: /hello<br>
Method: POST<br>
Success response:
```
{
    "code": 20000,
    "data": {
        "name": "viola",
        "greeting": "Bonjour!"
    }
}
```

Endpoint: /hello<br>
Method: GET<br>
Error response:
```
{
    "code": 40000,
    "message": "Method not allowed"
}
```

You can use "code" for internal purposes that are not covered by HTTP status codes, such as helping the frontend distinguish response handling.


Log:
```
endpoint: /hello
method: POST
Given payload:
{
    "username": "okesips",
    "token": "okeee",
    "oke": "sip"
}

secretKeyword = []string{"token", "password", "auth"}

stdout:
----
2023/09/06 12:07:39  INFO MDC_GROUP=26ee5fd5-9c60-4595-a2f3-f17ab24fbb36 "log to test mdc" :::secretKeywordsRemovedFromLog=
2023/09/06 12:07:39  INFO MDC_GROUP=26ee5fd5-9c60-4595-a2f3-f17ab24fbb36 ":::method=POST :::statusCode=200 :::uri=localhost:8080/hello :::headers="Content-Type":"application/json", "User-Agent":"PostmanRuntime/7.29.2", "Postman-, "Accept-Encoding":"gzip, deflate, br", "Cache-Control":"no-cache", "Connection":"keep-alive", "Accept":"*/*", "Content-Length":"69",  :::body={    "username": "okesips",    ",    "oke": "sip"} :::response={"code":20000,"data":{"name":"viola","greeting":"Bonjour!"}}" :::secretKeywordsRemovedFromLog=token,
----
```

Any value match with secret keywords automatically removed from stdout.<br>
Currently, only MDC_GROUP is being used as an example. You can add more MDCs as needed.<br>
You can modify the log filter in the file aklog/aklog.go by changing or adding keywords to the secretKeywords variable.<br>
MDC will be very helpful for you in searching and grouping logs.


## Unit Test
Just run "go test" on the "test" folder

Test using cucumber (BDD style), learn more: https://github.com/cucumber/godog
