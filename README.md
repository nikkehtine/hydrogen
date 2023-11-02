# Hydrogen compiler

Compiler in Go following a tutorial by [@orosmatthew](https://github.com/orosmatthew) - [Creating a Compiler](https://www.youtube.com/playlist?list=PLUDlas_Zy_qC7c5tCgTMYq2idyyT241qs) (originally in C++)

[Link to original project repo](https://github.com/orosmatthew/hydrogen-cpp)

## Usage

Requirements:

-   Go >= 1.21
-   `nasm`
-   `ld`

> [!WARNING]
> Currently this program only works on Linux

1. Clone the repo

    ```bash
    git clone https://github.com/nikkehtine/hydrogen.git
    ```

2. Build and run the program

    You can just use the `test.sh` script for now, which build the program, runs it, and then deletes the generated output files

    ```bash
    ./test.sh
    ```

    This script also accepts certain commands for convenience:

    - `compile`: just compile the program, don't run it or remove generated files
    - `run`: compile and run the program, but don't delete any files
    - `cleanup`: just delete any generated output files that exist

## What I've learned and want to learn

I've learned:

-   how compilers work (things such as tokenizer, lexer, outputting to assembly, linking)
-   how to turn code into tokens, and then translate tokens into different code
-   how programs operate on low level and communicate with the operating system
-   more advanced concepts about programming and the Go language, such as buffers, enums, operations on files, accepting cli arguments

I want to:

-   learn the differences in the way OSes (Windows, macOS, Linux) communicate with programs
-   learn more on how certain operations and language features work on a lower level
-   learn more advanced concepts about programming
-   be able to brag about having written my own compiler :)

## Useful resources

-   [Go package registry](https://pkg.go.dev)
    -   [`os`](https://pkg.go.dev/os)
    -   [`fmt`](https://pkg.go.dev/fmt)
    -   [`unicode`](https://pkg.go.dev/unicode)
    -   [`bytes`](https://pkg.go.dev/bytes)
-   ["How to do enums in Go" by Marco Franssen](https://marcofranssen.nl/how-to-do-enums-in-go)
-   [Go by Example](https://gobyexample.com)
    -   [Command-line arguments](https://gobyexample.com/command-line-arguments)
    -   [Reading files](https://gobyexample.com/reading-files)

## Contact

Shoot me a DM on Mastodon: [@nikkehtine@im-in.space](https://im-in.space/@nikkehtine)
