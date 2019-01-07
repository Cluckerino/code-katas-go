# code-katas-go

My Go Kata solutions for Codewars

Boy, golang is hard

I use Go's built-in [testing](https://golang.org/pkg/testing/) framework extended with the [Testify](https://github.com/stretchr/testify) library

To create and use a Docker image to build and run the test:

```sh
docker build -t go-test-runner . \
    && docker run --rm -it --name cpp-tests go-test-runner
```
