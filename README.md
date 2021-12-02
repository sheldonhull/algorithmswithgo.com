# Algorithms With Go + Other Algorithm Challenges

Initially, setup this to perform the challenges in algorithmswithgo.com repo.
However, as I've explored more with Hackerrank, exercism.io and other challenges, I decided I'll use this repo for all of these to log my progress.

Any blog progress on this will be posted on [my blog](sheldonhull.com/microblog)

## Badges

![](https://img.shields.io/badge/day%20📅-6-blue)

![](https://img.shields.io/badge/stars%20⭐-12-yellow)

![](https://img.shields.io/badge/days%20completed-6-red)

## Setup Exercism

[Get token](https://exercism.io/my/settings).

```shell
exercism configure --token=$TOKEN
```

Configure workspace for CLI to use this repo.

```shell
exercism configure --workspace ${PWD}/exercism.io
```

## Advent of Code

To respect the creators, the inputs are not persisted, just the solution.
To download the inputs I'm using:

```go
go install github.com/GreenLightning/advent-of-code-downloader/aocdl@latest
```

Setup a `.aocdlconfig` and populate like below.
This is also part of `.gitignore`.

```json
{
    "session-cookie": "",
    "output": ".inputs/{{.Year}}-{{.Day}}.txt",
    "year": 2021,
    "day": 1
}

```
Use by: `aocdl -day 1`
