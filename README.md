# kiam-redirect

redirects all request from comptandye.fr to kiam.fr :smiley:

### build 
`docker build -t trendev/kiam-redirect .`

### run
`docker run -d -p 80:9000 -e URL="https://www.kiam.fr" trendev/kiam-redirect`

### test
`curl localhost`

> `<a href="https://www.kiam.fr">Permanent Redirect</a>.`