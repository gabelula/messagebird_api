# INSTALL

Go Version 1.8.3

Set environment variables:

```Shell
export HOST=localhost:3000
export MESSAGEBIRD_API_KEY="4hr42b0g9s3tKbRjdY8mv6Hp5"
```

To run the webserver:

```GO
    go run app/main.go
```

The POST goes to [http://127.0.0.1:3000/v1/message](http://127.0.0.1:3000/v1/message) with body

```JSON
{
 "body": "THE MESSAGE",
 "originator": "NAME",
 "recipient": PHONE NUMBER
}
```