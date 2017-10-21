# Disk Eject

The **one and only** purpose of Disk Eject is to eject all external disk(s) from macOS.

## Why?

There is applescript like
[this](https://gist.github.com/jaux/b662d31aafefdd837822045674620ae1) trying to
achieve the same goal, but I couldn't get a **clean unmount** for time machine
volume! There were some popup windows showed up during the process saying that
some volumes were ejected, that was very annoying. Therefore I decide to roll
my own program for this task.

## Install

```
go get github.com/jaux/diskeject
```

## Run

```
$GOPATH/bin/diskeject
```

Or, if you have added `$GOPATH/bin` to your `$PATH`, you can simply run

```
diskeject
```



## License

[MIT](/LICENSE)
