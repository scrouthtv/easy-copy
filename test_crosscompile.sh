export GOOS=linux;     export GOARCH=arm;   echo $GOOS/$GOARCH; go build
export GOOS=linux;     export GOARCH=arm64; echo $GOOS/$GOARCH; go build
export GOOS=linux;     export GOARCH=amd64; echo $GOOS/$GOARCH; go build
export GOOS=linux;     export GOARCH=386;   echo $GOOS/$GOARCH; go build
export GOOS=darwin;    export GOARCH=arm64; echo $GOOS/$GOARCH; go build
export GOOS=darwin;    export GOARCH=amd64; echo $GOOS/$GOARCH; go build
export GOOS=freebsd;   export GOARCH=arm;   echo $GOOS/$GOARCH; go build
export GOOS=freebsd;   export GOARCH=amd64; echo $GOOS/$GOARCH; go build
export GOOS=freebsd;   export GOARCH=386;   echo $GOOS/$GOARCH; go build
export GOOS=openbsd;   export GOARCH=arm;   echo $GOOS/$GOARCH; go build
export GOOS=openbsd;   export GOARCH=amd64; echo $GOOS/$GOARCH; go build
export GOOS=openbsd;   export GOARCH=386;   echo $GOOS/$GOARCH; go build
export GOOS=netbsd;    export GOARCH=arm;   echo $GOOS/$GOARCH; go build
export GOOS=netbsd;    export GOARCH=amd64; echo $GOOS/$GOARCH; go build
export GOOS=netbsd;    export GOARCH=386;   echo $GOOS/$GOARCH; go build
export GOOS=dragonfly; export GOARCH=amd64; echo $GOOS/$GOARCH; go build

