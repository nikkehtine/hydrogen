#!/bin/sh
# Basic script to quickly test the program
# Commands: empty, "compile", "run", "cleanup"

compile() {
    go run . test.hyd 
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
        cleanup
        ;;
    "compile")
        compile
        ;;
    "run")
        compile
        run
        ;;
    "cleanup")
        cleanup
        ;;
    *)
        echo "Available arguments:"
        echo "no argument: compiles and runs the program,"
        echo "then cleans up the directory"
        echo "compile: compiles the program without running it"
        echo "run: compiles and runs the program"
        echo "cleanup: cleans up the output files"
        exit 1
esac
