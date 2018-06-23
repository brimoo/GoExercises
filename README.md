# GoExercises

## About
Hello! This repo contains some of the exercises I used to get familiarized with the Go language. 

## How to Run
Make sure Go is installed first, then you can run the code in one of two ways. Firstly, you can clone this
repository into the src folder of your Go workspace. Then each package can be installed by running `go install` on 
the package you wish to install, then an executable will be generated in the bin folder of your workspace. 
Alternatively, if you wish to build an individual go file without setting up a workspace you can simply 
clone the repo anywhere and use`go run` on the file you wish to run.

## Package Descriptions

* __Hello__: The classic "Hello, World!" program

* __Factorial__: Computes the factorial of the user entered number using dynamic programming (bottom up)

* __Hamming__: Computes the Hamming distance between two DNA strands

* __Lettercount__: Counts the occurrences of a letter in multiple files using concurrency

* __Anagrams__: User enters a word and a list of words. Outputs words from the list that are anagrams of the first word.

* __HelloTCP__: Listens on port 4000 and prints "Hello!" to the connection.

* __EchoTCP__: Listens on port 4000 and echoes anything that is written on the connection.

* __ChatTCP__: Listens on port 4000 and waits for two users to connect. The users are paired and may chat over the connection.