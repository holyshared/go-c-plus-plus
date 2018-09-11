# go-c-plus-plus

swigのインターフェースファイル（.swigcxx）の％{...％}ブロックのすべては、SWIGによって作成された結果のラッパーファイルにそのままコピーされます。

## Setup

```shell
brew install go swig
```

## CPP build

```shell
cd cpp
make
```

### Swig version

```shell
vagrant@ubuntu-xenial:~$ swig -version

SWIG Version 3.0.8

Compiled with g++ [x86_64-pc-linux-gnu]

Configured options: +pcre

Please see http://www.swig.org for reporting bugs and further information
```

### Go version

```shell
go version go1.6.2 linux/amd64
```
