# Hydrogen compiler

Compiler in Go following a tutorial by [Pixeled](https://www.youtube.com/@pixeled-yt) - [Creating a Compiler](https://www.youtube.com/playlist?list=PLUDlas_Zy_qC7c5tCgTMYq2idyyT241qs) (originally in C++)

[Link to original project repo](https://github.com/orosmatthew/hydrogen-cpp)

## Usage

Requirements:

- Go >= 1.21
- `nasm`
- `ld`

**TODO**: Currently compiles for Linux

1. Clone the repo

```bash
git clone https://gitlab.com/nikkehtine/hydrogen.git
```

2. Build and run the program

   There are two ways:

   1. ```bash
      go build -o hydro<.exe> # .exe if you're on Windows
      ./hydro<.exe> <input file>
      ```

   2. ```bash
      go run . <input file>
      ```

3. Assemble and link the compiled program

```bash
nasm -f <format> out.asm # format: elf64 for Linux, win64 for Windows
ld -o out<.exe> out.o<bj> # .o on Linux, .obj on Windows
```

## What I've learned and want to learn

- I've learned how compilers work (things such as tokenizer, lexer, outputting to assembly, linking)
- I've learned how to turn code into tokens, and then translate tokens into different code
- I've learned how programs operate on low level and communicate with the operating system
- I've learned more advanced concepts about programming and the Go language, such as buffers, enums, operations on files, accepting cli arguments
- I hope to learn the differences in the way OSes (Windows, macOS, Linux) communicate with programs
- I hope to learn more on how certain operations and language features work on a lower level
- I hope to learn more advanced concepts about programming
- I want to be able to brag about having written my own compiler :)

## Useful resources

- [Go package registry](https://pkg.go.dev)
  - [`os`](https://pkg.go.dev/os)
  - [`fmt`](https://pkg.go.dev/fmt)
  - [`unicode`](https://pkg.go.dev/unicode)
  - [`bytes`](https://pkg.go.dev/bytes)
- ["How to do enums in Go" by Marco Franssen](https://marcofranssen.nl/how-to-do-enums-in-go)
- [Go by Example](https://gobyexample.com)
  - [Command-line arguments](https://gobyexample.com/command-line-arguments)
  - [Reading files](https://gobyexample.com/reading-files)

## Contact

Shoot me a DM on Mastodon: [@nikkehtine@im-in.space](https://im-in.space/@nikkehtine)
