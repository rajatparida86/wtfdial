# wtfdial
Tracks if a team member is wtf'ing

## Prerequisites

### Install protoc as this project uses Prototbuf

#### On Mac
1. Download the appropriate release here: https://github.com/google/protobuf/releases

2. Unzip the folder

3. Enter the folder and run

```
$ ./autogen.sh && ./configure && make
```

4. If you run into this error: autoreconf: failed to run aclocal: No such file or directory, then:

```
$ run brew install autoconf && brew install automake && brew install libtool
```

And run the command from step 3 again.

5. Then run these other commands. They should run without issues
```
$ make check

$ sudo make install

$ which protoc

$ protoc --version
```

#### On Linux
```
$wget https://github.com/google/protobuf/releases/download/v2.5.0/protobuf-2.5.0.tar.bz2

$tar xvf protobuf-2.5.0.tar.bz2

$cd protobuf-2.5.0

$./configure CC=clang CXX=clang++ CXXFLAGS='-std=c++11 -stdlib=libc++ -O3 -g' LDFLAGS='-stdlib=libc++' LIBS="-lc++ -lc++abi"

$make -j 4

$sudo make install

$protoc --version
```

### Install gogo protobuf and extensions
```
$ go get github.com/gogo/protobuf/proto
$ go get github.com/gogo/protobuf/protoc-gen-gogo
$ go get github.com/gogo/protobuf/gogoproto
```
