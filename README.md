### Installation
Download the built binary from the release page and move it to your bin directory [release page](https://github.com/joechea-aupp/passgen/releases)
put the binary in your bin directory


### Usage
By default, passgen will generate 5 words with `-` as separator, generate password will automatically copy to clipboard.
```bash
passgen -n=5 -sep=~
# output Euphrates~Hellenism's~Wales~sole~legmen
```

You can also use it to generate a token with specific lenght
```bash
passgen -phrases=false -c=8
# output W#LuJMB)
```

### Development
make sure you have `wamerican` or `wbritish` installed on your system, you can install it by running
```bash
sudo apt-get install wamerican
```

to compile the code, run the following command
build for linux
```bash
go mod tidy
go build -o passgen
```

build for windows
```bash
GOOS=windows GOARCH=amd64 go build -o passgen.exe
```
