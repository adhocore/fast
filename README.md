# adhocore/fast

A GO lang command line tool to check internet speed right from the terminal.

> Uses [fast.com](https://fast.com) through headless chrome.

## Usage

First, Make sure you have `chrome` binary available in `$PATH` or `%path%`:
```sh
which chrome
```

If you do not have it, in MacOS, you can do something like this:

```sh
echo '#!/bin/sh\n\n/Applications/Google\ Chrome.app/Contents/MacOS/Google\ Chrome $@' > /usr/local/bin/chrome

chmod +x /usr/local/bin/chrome
```

In other OS, you can do something equivalent to above. The idea is `chrome` command should point to **Chrome Browser**.

Then install `fast`:
```sh
go get github.com/adhocore/fast/cmd/fast
```

Finally, make sure `$GOPATH` or `$HOME/go/bin` is in your `$PATH` or `%path%`, then run
```
fast
```

Wait a while or `Ctrl+C` if you can't. That's all.

## Screen

![FAST](./assets/usage.png)
