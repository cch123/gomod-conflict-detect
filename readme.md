# gomod-conflict-detect


Install:

```shell
go get -v -u github.com/cch123/gomod-conflict-detect

cd ${YOUR PROJECT PATH}
```

Usage:

In your project path, execute `gomod-conflict-detect` to get possible path conflict.

If you want to find out which dependency imports a library, execute `gomod-conflict-detect | grep LIBNAME`

The output is like:

```
Conflict in pkg github.com/klauspost/compress paths are:
 mosn.io/mosn -> github.com/klauspost/compress@v1.7.5
 mosn.io/mosn -> github.com/valyala/fasthttp@v1.2.0 -> github.com/klauspost/compress@v1.4.0
```

TODO

Convert output to dot svg

