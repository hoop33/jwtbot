# jwtbot

> A CLI for creating JWTs

## Introduction

jwtbot (pronounced "jot - bot" lets you create JWTs from the command line. As a tool, it's woefully anemic: it supports setting a subject, an algorithm, and a signing key only. Hey -- that's all I needed for right now.

*Note:* Not all "supported" algorithms currently work. I really just created this because I frequently need JWTs for local development, and I don't want to bother with cert generation. So I just use "none" for the algorithm. I used to use <https://jwt.io> to create these JWTs, but they've apparently dropped support for "none." So I know "none" and the "hs*" algorithms work, and the other ones don't yet. One day, I'll probably get to it . . . but I gladly accept PRs!

## Installation

If you have a working Go installation, type:

```sh
$ go get -u github.com/hoop33/jwtbot
```

## Usage

Get help at any time by running:

```sh
$ jwtbot
```

### List Available Algorithms

```sh
$ jwtbot algorithms
```

### Create a JWT

Currently, you can set a subject, choose an algorithm, and set a signing key. That's all.

```sh
$ jwtbot create -s hoop33 -k a-secret -a hs256
```

## Credits

jwtbot uses the following open source libraries -- thank you!

* [Cobra](https://github.com/spf13/cobra)
* [jwt-go](https://github.com/dgrijalva/jwt-go)

Apologies if I've inadvertently omitted any library or any contributor.

## License

Copyright &copy; 2019 Rob Warner

Licensed under the [MIT License](https://hoop33.mit-license.org/)
