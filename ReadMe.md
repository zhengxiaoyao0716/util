# util
## Golang utilities kit.

> Those utils come from my actual production, they are all simple, shot but usefully. Each util was built in a separate module and follow with unit test (example), pursue of simple and convenience.

***
### Install

    go get github.com/zhengxiaoyao0716/util/<package you want>

    // Note that the project path `util` is only a directory,
    // you should choose a package in child path to import.


***
### Document

See [godoc](https://godoc.org/github.com/zhengxiaoyao0716/util)


***
### Example

See unit test in packages.

***
### Change log

- 2017/08/02: The interface of `Event` module has a bit change. the `Key` has change from `struct{Type, Name}` to `[2]string{Type, Name}`. 
