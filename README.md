    docker build -t go-rust .
    docker run --rm -it -v/Users/lenny/go:/go go-rust /go/src/github.com/llofberg/opa-rust/build.sh
