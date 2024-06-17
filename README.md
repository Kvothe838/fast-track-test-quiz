﻿# Fast Track Test Quiz

 Solution for Fast Track Test Quiz.

Instructions:

- The task is to build a super simple quiz with a few questions and a few alternatives for each question. Each with one correct answer.

Preferred Stack:
- Backend - Golang
- Database - Just in-memory, so no database

Preferred Components:
- REST API or gRPC
- CLI that talks with the API, preferably using https://github.com/spf13/cobra ;( as CLI framework )

User stories/Use cases:
- User should be able to get questions with a number of answers
- User should be able to select just one answer per question.
- User should be able to answer all the questions and then post his/hers answers and get back how many correct answers they had, displayed to the user.
- User should see how well they compared to others that have taken the quiz, eg. "You were better than 60% of all quizzers"

 ## Prerequisites

 - Install Golang: https://go.dev/doc/install
 
 - Install Makefile: https://medium.com/@samsorrahman/how-to-run-a-makefile-in-windows-b4d115d7c516
 
 - Install golangci-lint: https://golangci-lint.run/

 - Create local-env/config.yaml, add same env variables than local-env/config-example.yaml and set their values according to your Postgres connection.

 ## Run

 To run the backend, just execute on the root of the project the following command:

 ```
make run
```

## Test

To run the unit tests, just execute on the root of the project the following command:

```
make test
```

## Lint

To run lint (it help us to write better code), just execute on the root of the project the following command:

```
make lint
```