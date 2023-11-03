#!/bin/sh
# Basic script to quickly test the program
# Commands: empty, "compile", "run", "cleanup"


# Variables
OUTDIR=tmp
FILE=$OUTDIR/test

# Helper function
checkErr() {
    if [ $? -ne 0 ]; then
        echo "!!! script failed during $1"
        exit 1
    fi
}

# Actual steps

compile() {
    go run . test.hy
    checkErr "compilation"

    if [ ! -d $OUTDIR ]; then
        mkdir $OUTDIR
    fi
    mv out.asm $FILE.asm

    nasm -felf64 $FILE.asm -o $FILE.o
    checkErr "assembling"

    ld -o $FILE $FILE.o
    checkErr "linking"
}

run() {
    ./$FILE
    echo "Program exit code: $?"
}

cleanup() {
    if [ -d $OUTDIR ]; then
        rm -r $OUTDIR
    fi
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
