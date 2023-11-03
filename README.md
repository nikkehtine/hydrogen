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

    You can just use the `test.sh` script for now, which should build and run the program

    ```bash
    ./test.sh
    ```

    This script also accepts certain commands for convenience:

    - `compile`: just compile the program, don't run it or remove generated files
    - `clean`: compile and run the program, but delete generated files
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
-   [Idiomatic way in Go to represent a Tagged Union? (r/golang)](https://www.reddit.com/r/golang/comments/13hjevf/idiomatic_way_in_go_to_represent_a_tagged_union/)
-   ["Sum/Union/Variant Type in Go" by haya14busa (Medium)](https://medium.com/@haya14busa/sum-union-variant-type-in-go-and-static-check-tool-of-switch-case-handling-3bfc61618b1e)

## Contact

Shoot me a DM on Mastodon: [@nikkehtine@im-in.space](https://im-in.space/@nikkehtine)
