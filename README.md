# adhocore/fast

A GO lang command line tool to check internet speed right from the terminal.

> Uses [fast.com](https://fast.com) through headless chrome.

## Prerequistie

Chrome browser must be installed. **`chromedp`** will try to locate the chrome executable automatically from these [paths](https://github.com/chromedp/chromedp/blob/master/allocate.go#L334-L352).

> If you get error regarding chrome availability, and you have chrome in custom path then check [Troubleshooting](#troubleshooting).

## Usage

Install `fast`:
```sh
go get github.com/adhocore/fast/cmd/fast
```

Finally, make sure `$GOPATH` or `$HOME/go/bin` is in your `$PATH` or `%path%`, then run:
```sh
fast
```

Wait a while or `Ctrl+C` if you can't. That's all.

## Screen

![FAST](./assets/usage.png)

## Troubleshooting

In **MacOS**, you can do something like this:

```sh
echo '#!/bin/sh\n\n/Applications/Google\ Chrome.app/Contents/MacOS/Google\ Chrome $@' > /usr/local/bin/chrome

chmod +x /usr/local/bin/chrome
```

In **WSL**, you can symlink chrome from host WinOS like this:
```sh
sudo ln -s /mnt/c/Program\ Files/Google/Chrome/Application/chrome.exe /usr/local/bin/chrome
```

In other OS, you can do something equivalent to above. The idea is `chrome` command should point to **Chrome Browser**.
