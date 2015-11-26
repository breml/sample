#!/bin/bash
#
# Generate division implementations based on skelleton in division_uint.go
# Also removes the build tag and the go:generate instructions
#
# Usage:
#   go generate *.go
#
if [[ "$GOFILE" == *_test.go ]]; then
	OUTFILE="${GOFILE%_test\.go}${1}_test.go"
else
	OUTFILE="${GOFILE%\.go}${1}.go"
fi

sed "s/\([uU]\)int8/\1int${1}/g" $GOFILE | grep -vE "//([+]build |go:)generate" | gofmt > $OUTFILE
