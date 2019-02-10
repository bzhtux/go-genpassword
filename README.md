```
  ____        ____            ____                                     _ 
 / ___| ___  / ___| ___ _ __ |  _ \ __ _ ___ _____      _____  _ __ __| |
| |  _ / _ \| |  _ / _ \ '_ \| |_) / _` / __/ __\ \ /\ / / _ \| '__/ _` |
| |_| | (_) | |_| |  __/ | | |  __/ (_| \__ \__ \\ V  V / (_) | | | (_| |
 \____|\___/ \____|\___|_| |_|_|   \__,_|___/___/ \_/\_/ \___/|_|  \__,_|
                                                                         
```

Golang generate password

# Prerequisites

I assume you have a clean Go installation, that includes `$GOPATH` defined:

```bash
# for macos
echo $GOPATH
/Users/<login>/go

# for linux
echo $GOPATH
/home/<login>/go
```

# Build

For Linux and MacOS only:

```bash
# clone repository
git clone git@github.com:bzhtux/go-genpassword.git

# get dependancies
go get ./...

# build binary
cd go-genpassword
go build -o gen_password
sudo cp gen_password /usr/local/bin

# test gen_password binary
gen_password -h
.:HELP:.
--------
 -l		password's length (int)
 -n		numerics' count (int)
 -s		symbols' count (int)
Example: gen_password -l 24 -n 6 -s 4
```

# Docker

If you need to generate password remotly you may want to use a docker container (e.g for concourse-ci):

```bash
# clone repository
git clone git@github.com:bzhtux/go-genpassword.git

# build docker image
cd go-genpassword
docker build . -t gen-password:1.0.0

# test gen_password binary
docker run -ti --rm gen-password:1.0.0 -h
.:HELP:.
--------
 -l		password's length (int)
 -n		numerics' count (int)
 -s		symbols' count (int)
Example: gen_password -l 24 -n 6 -s 4
```

