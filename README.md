# Introduction to GoLang

## Intro

- Compiled language
- Runtime is baked into the final product => No VM required like Java
- Executables are exported at the compile time
- Executables are different for OS
- Where to expect Go
  - System apps to web apps - cloud
- It is already in production
- Don't bring baggage of previous languages
- Object Oriented?
  - Yes and No
  - Object oriented language should have classes and objects, Go does'nt have classes (we have structs though)
  - There is no Overriding/overloading in Go
  - No try-catch
  - lexer does a lot of work

## Hello World

- write `go mod init hello` in the terminal
- write `go run main.go` in the terminal to run the file
  - This going to run the file, it will not give any executables.

## GOPATH and reading go docs

- write `go help` in the terminal to get some help for Go
- write `go env GOPATH` and explore other commands

## Lexer in Golang and Types

- ### Lexer
  - There is something called Lexer
    - The job of the lexer is to make sure grammer of the language is being followed at pre-compile time
    - If intellisense is installed, then Lexer will automatically kick in and put semicolons
- ### Types
  - Case insensitive
    - Define if public or private
      - First letter defines if the thing is public or private
  - Variable type should be known in advance
  - Everything is a Type
  - #### Simple Types

# Basics of GoLang
