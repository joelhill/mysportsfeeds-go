#!/usr/bin/env bash

LIB="github.com/joelhill/go-rest-http-blaster"
REPO="https://github.com/joelhill/go-rest-http-blaster/blob/master/"
FILE="GODOC.md"
SRCTGT="/src/target/"
SRC="/src/github.com/joelhill/go-rest-http-blaster/"

command -v godoc2md >/dev/null 2>&1 || { echo "*** gendoc requires godoc2md ***" >&2; echo "Get it at https://github.com/davecheney/godoc2md" >&2; exit 1; }
godoc2md ${LIB} > ./GODOC_TMP

cat ./GODOC_TMP | sed 's,'"${SRCTGT}"','"${REPO}"',g' > ./GODOC_TMP2
cat ./GODOC_TMP2 | sed 's,'"/src/${LIB}/"','"${REPO}"',g' > ./GODOC.md
rm -f ./GODOC_TMP
rm -f ./GODOC_TMP2
