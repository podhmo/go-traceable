# go-traceable

## go-run-httptrace

calling [httptrace/_example/github/main.go](httptrace/_example/github/main.go) with trace.

```console
$ go get -v github.com/podhmo/go-traceable/cmd/go-run-httptrace
$ go-run-httptrace _example/github/main.go
2018/12/02 22:03:07 parse httptrace/_example/github/main.go
2018/12/02 22:03:07 transform AST
2018/12/02 22:03:07 format
GET /repos/podhmo/go-traceable/contributors HTTP/1.1
Host: api.github.com
User-Agent: go-github/8
Accept: application/vnd.github.v3+json
Accept-Encoding: gzip

HTTP/1.1 200 OK
Transfer-Encoding: chunked
Access-Control-Allow-Origin: *
Access-Control-Expose-Headers: ETag, Link, Location, Retry-After, X-GitHub-OTP, X-RateLimit-Limit, X-RateLimit-Remaining, X-RateLimit-Reset, X-OAuth-Scopes, X-Accepted-OAuth-Scopes, X-Poll-Interval, X-GitHub-Media-Type
Cache-Control: public, max-age=60, s-maxage=60
Content-Security-Policy: default-src 'none'
Content-Type: application/json; charset=utf-8
Date: Sun, 02 Dec 2018 13:03:08 GMT
Etag: W/"c68f67bb376d5b26356b0f3c3a5b5714"
Last-Modified: Sun, 02 Dec 2018 13:02:20 GMT
Referrer-Policy: origin-when-cross-origin, strict-origin-when-cross-origin
Server: GitHub.com
Status: 200 OK
Strict-Transport-Security: max-age=31536000; includeSubdomains; preload
Vary: Accept
Vary: Accept-Encoding
X-Content-Type-Options: nosniff
X-Frame-Options: deny
X-Github-Media-Type: github.v3; format=json
X-Github-Request-Id: E066:814D:1A64264:223FE3A:5C03D80C
X-Ratelimit-Limit: 60
X-Ratelimit-Remaining: 50
X-Ratelimit-Reset: 1543758654
X-Xss-Protection: 1; mode=block

37e
[{"login":"podhmo","id":80613,"node_id":"MDQ6VXNlcjgwNjEz","avatar_url":"https://avatars1.githubusercontent.com/u/80613?v=4","gravatar_id":"","url":"https://api.github.com/users/podhmo","html_url":"https://github.com/podhmo","followers_url":"https://api.github.com/users/podhmo/followers","following_url":"https://api.github.com/users/podhmo/following{/other_user}","gists_url":"https://api.github.com/users/podhmo/gists{/gist_id}","starred_url":"https://api.github.com/users/podhmo/starred{/owner}{/repo}","subscriptions_url":"https://api.github.com/users/podhmo/subscriptions","organizations_url":"https://api.github.com/users/podhmo/orgs","repos_url":"https://api.github.com/users/podhmo/repos","events_url":"https://api.github.com/users/podhmo/events{/privacy}","received_events_url":"https://api.github.com/users/podhmo/received_events","type":"User","site_admin":false,"contributions":1}]
0

2018/12/02 22:03:08 rollback httptrace/_example/github/main.go
```
