main.go

```go
package main

import (
	"fmt"

	"github.com/podhmo/go-traceable/httptrace"
)

func main() {
	resp, err := httptrace.Get("https://example.com")
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.Status)
}
```

```console
$ go run main.go
200 OK
$ TRACE=1 go run main.go
GET / HTTP/1.1
Host: example.com
User-Agent: Go-http-client/1.1
Accept-Encoding: gzip

HTTP/2.0 200 OK
Accept-Ranges: bytes
Cache-Control: max-age=604800
Content-Type: text/html
Date: Tue, 29 May 2018 22:26:06 GMT
Etag: "1541025663"
Expires: Tue, 05 Jun 2018 22:26:06 GMT
Last-Modified: Fri, 09 Aug 2013 23:54:35 GMT
Server: ECS (sjc/4E8D)
Vary: Accept-Encoding
X-Cache: HIT

<!doctype html>
<html>
<head>
    <title>Example Domain</title>

    <meta charset="utf-8" />
    <meta http-equiv="Content-type" content="text/html; charset=utf-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1" />
    <style type="text/css">
    body {
        background-color: #f0f0f2;
        margin: 0;
        padding: 0;
        font-family: "Open Sans", "Helvetica Neue", Helvetica, Arial, sans-serif;
        
    }
    div {
        width: 600px;
        margin: 5em auto;
        padding: 50px;
        background-color: #fff;
        border-radius: 1em;
    }
    a:link, a:visited {
        color: #38488f;
        text-decoration: none;
    }
    @media (max-width: 700px) {
        body {
            background-color: #fff;
        }
        div {
            width: auto;
            margin: 0 auto;
            border-radius: 0;
            padding: 1em;
        }
    }
    </style>    
</head>

<body>
<div>
    <h1>Example Domain</h1>
    <p>This domain is established to be used for illustrative examples in documents. You may use this
    domain in examples without prior coordination or asking for permission.</p>
    <p><a href="http://www.iana.org/domains/example">More information...</a></p>
</div>
</body>
</html>
200 OK
```

## patch

```go
package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/podhmo/go-traceable/httptrace"
)

func main() {
	teardown := httptrace.Patch()
	defer teardown()
	resp, err := http.Get("https://example.com")
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.Status)
}
```
