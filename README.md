# adhocore/fast

[![Go Report](https://goreportcard.com/badge/github.com/adhocore/fast)](https://goreportcard.com/report/github.com/adhocore/fast)
[![Donate](https://img.shields.io/badge/donate-paypal-blue.svg?style=flat-square)](https://www.paypal.me/ji10/50usd)
[![Tweet](https://img.shields.io/twitter/url/http/shields.io.svg?style=social)](https://twitter.com/intent/tweet?text=Golang+tool+to+check+internet+speed+right+from+the+terminal&url=https://github.com/adhocore/fast&hashtags=php,jwt,auth)


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

---
### Other projects
My other golang projects you might find interesting and useful:

- [**gronx**](https://github.com/adhocore/gronx) - Lightweight, fast and dependency-free Cron expression parser (due checker), task scheduler and/or daemon for Golang (tested on v1.13 and above) and standalone usage.
- [**urlsh**](https://github.com/adhocore/urlsh) - URL shortener and bookmarker service with UI, API, Cache, Hits Counter and forwarder using postgres and redis in backend, bulma in frontend; has [web](https://urlssh.xyz) and cli client
