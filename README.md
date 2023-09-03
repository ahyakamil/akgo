# AKGo
Easy and simple Go framework for build API

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
    "code": 40005,
    "message": "Method not allowed"
}
```

Log:
```
Given payload:
{
    "username": "okesips",
    "token": "okeee",
    "oke": "sip"
}

secretKeyword = []string{"token", "password", "auth"}

stdout:
----
2023/09/03 21:57:26  WARN MDC_GROUP=6033b2b6-6e30-491c-a15f-fd0608457997 ":::method=GET :::headers="Accept":"*/*", "Postman-, "Accept-Encoding":"gzip, deflate, br", "Connection":"keep-alive", "Content-Length":"69", "Content-Type":"application/json", "User-Agent":"PostmanRuntime/7.29.2",  :::body={    "username": "okesips",    ",    "oke": "sip"} :::statusCode=405 :::response={"code":40005,"message":"Method not allowed"}" :::secretKeywordsRemovedFromLog=token,
----
```

Any value match with secret keyword automatically removed from stdout.<br>
Currently, only MDC_GROUP is being used as an example. You can add more MDCs as needed.<br>
You can modify the log filter in the file aklog/aklog.go by changing or adding keywords to the secretKeyword variable.
