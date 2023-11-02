#!/bin/sh
# Basic script to quickly test the program
# Commands: empty, "compile", "run", "cleanup"

compile() {
    go run . test.hy
    nasm -felf64 out.asm
    ld -o out out.o
}

run() {
    ./out
    echo $?
}

cleanup() {
    rm out.asm
    rm out.o
    rm out
}

case $1 in
    "")
        compile
        run
        ;;
    "compile")
        compile
        ;;
    "clean")
        compile
        run
        cleanup
        ;;
    "cleanup")
        cleanup
        ;;
    *)
        echo "Available arguments:"
        echo "no argument: compiles and runs the program"
        echo "compile: compiles the program without running it"
        echo "clean: compiles and runs the program, then deletes output files"
        echo "cleanup: just cleans up the output files"
        exit 1
esac
