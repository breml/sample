#!/bin/bash
#
# Generate division implementations based on skelleton in division_uint.go
# Also removes the build tag and the go:generate instructions
#
# Usage:
#   go generate *.go
#
# or
#   go generate -tags generate .
#

if [ "$GOFILE" == "" ]; then
	echo "error: GOFILE variable not set, script is intendet for usage in conjunction with go generate"
	exit 1
fi

if [ "$1" == "" ]; then
	echo "error: no argument given, please provide bit size of int: 8, 16, 32 or 64"
	exit 1
fi

SED=`which sed`
if [ "$SED" == "" ]; then
	echo "error: sed command not found"
	exit 1
fi

GREP=`which grep`
if [ "$GREP" == "" ]; then
	echo "error: grep command not found"
	exit 1
fi

GOFMT=`which gofmt`
if [ "$GOFMT" == "" ]; then
	echo "error: gofmt command not found"
	exit 1
fi

if [[ "$GOFILE" == *_test.go ]]; then
	OUTFILE="${GOFILE%_test\.go}${1}_test.go"
else
	OUTFILE="${GOFILE%\.go}${1}.go"
fi

$SED "s/\([uU]\)int8/\1int${1}/g" $GOFILE | $GREP -vE "//([+]build |go:)generate" | $GOFMT > $OUTFILE
