# Passcode

This rube module contains all you need a simple [Go](https://golang.org) based keyboard entry program. When running, it simply listens to all keyboard input awaiting for a preset keystroke combination (e.g. "12345"). Upon entry a http request is made to the Node-RED back-end that can trigger an action of your choice.

## Installation

Ensure you that have [Go](https://golang.org/dl/) installed on the target machine.

Clone the repo:

```
git clone https://github.com/projectrube/passcode
```

## Running

You can either run the application directly with `go run passcode.go` or build it first to create the binary with `go build`.

## Game & Challenge Ideas

Below are some ideas for using this rube...

### Treasure Map

This challenge uses `stages.go` which can be run with `go run stages.go`.

- Welcome to the treasure hunt 🎁
- To crack the code, you must find three clues scattered around the house 🏡
- Each clue will contain a secret 3 digit code (more on that in a moment) and the details of the next clue 🕵️‍♀️
- Type each 3 digit code into the keypad. If it's correct, you'll see a success message 🎉
- The last one will give you give you the location of the treasure 🏆
