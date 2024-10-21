# Go Unix Shell

A lightweight Unix shell written in Go. This shell provides basic functionalities like running commands, handling arguments, and managing the execution of background processes.

## Features

- Run Unix commands directly
- Handle command arguments and flags
- Execute commands in the background using `&`
- Basic error handling for invalid commands or missing arguments

## Prerequisites

Ensure you have [Go](https://golang.org/dl/) installed (version 1.18 or higher recommended).

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/dGilli/go-unix-shell.git
   ```

2. Navigate to the project directory:

   ```bash
   cd go-unix-shell
   ```

3. Build the shell:

   ```bash
   go build -o goshell
   ```

## Usage

Run the shell by executing the built binary:

```bash
./goshell
```

Once inside the shell, you can use it as follows:

- To run a command:

  ```bash
  ls -la
  ```

- To run a command in the background:

  ```bash
  sleep 10 &
  ```

- To exit the shell (or `CTRL-C`):

  ```bash
  exit
  ```

## Example

```bash
$ ./goshell
goshell> ls -la
goshell> sleep 5 &
goshell> exit
```

## License

This project is licensed under the MIT License.
