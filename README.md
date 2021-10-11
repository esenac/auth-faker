# auth-faker

auth-faker is a simple HTTP server that returns a signed JWT and the public key to read it. 

## How to run 
Currently the server can be run using Docker with the following commands:

```
docker build . -t auth-faker
docker run -p 8080:80 docker.io/library/auth-faker
```