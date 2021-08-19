### gosed

Commmand line tool for reading input from the console, replacing symbols `u2016` to `\n` and passing the output.

## Get
```
git clone git@github.com:ealebed/gosed.git
cd gosed
go install github.com/ealebed/gosed
```

## Building inside docker image:
```
docker run --rm -e GOOS=linux -v $(pwd):/project -w /project golang:1.17 go build -o build/gosed
```
or (if you use MacOS):
```
docker run --rm -e GOOS=darwin -v $(pwd):/project -w /project golang:1.17 go build -o build/gosed
```

## Usage
```
kubectl logs -f [pod_name] | gosed
```
