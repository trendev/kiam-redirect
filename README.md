# kiam-redirect

redirects all request from comptandye.fr to kiam.fr :smiley:

### build 
`docker build -t trendev/kiam-redirect .`

### run
`docker run -p 80:9000 -e URL="https://github.com/trendev" trendev/kiam-redirect`

### test
`curl -vvv localhost`

```
    * Trying ::1...
    * TCP_NODELAY set
    * Connected to localhost (::1) port 80 (#0)
    > GET / HTTP/1.1
    > Host: localhost
    > User-Agent: curl/7.64.1
    > Accept: */*
    > 
    < HTTP/1.1 308 Permanent Redirect
    < Content-Type: text/html; charset=utf-8
    < Location: https://github.com/trendev
    < Date: Fri, 10 Apr 2020 11:54:32 GMT
    < Content-Length: 62
    < 
    <a href="https://github.com/trendev">Permanent Redirect</a>.

    * Connection #0 to host localhost left intact
    * Closing connection 0`