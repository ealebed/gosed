### gosed

Commmand line tool for reading input from the console, replacing symbols `u2016` to `\n` and passing the output.

## Get
```
git clone git@github.com:ealebed/gosed.git
cd gosed
go build -v ./
```

## Usage
```
kubectl logs -f [pod_name] | ./gosed
```