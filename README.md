# auth-faker

auth-faker is a simple HTTP server that lets you forge a JWT token containing the claims you want, and the public key to read it. 
It can be used to mock authentication services in different scenarios, ranging from local development to testing, particularly in 
a microservices architecture.

## How to run 
Currently the server can be run using Docker with the following commands:

```
docker build . -t auth-faker
docker run -t auth-faker -p 8080:80 docker.io/library/auth-faker
```