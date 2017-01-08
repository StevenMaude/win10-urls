# win10-urls

Displays URLs for the Windows 10 English (not English International) ISO
download from Microsoft, either in Python or Go, whatever OS you're
currently running.

You should see two URLs: one for the 32-bit version and one for the
64-bit version. These are valid for 24 hours.

## Running

### Python

Install the requirements:

```sh
pip install --user lxml requests
```

(omit `--user` if working in a virtualenv)

and then run `python win10-urls.py`.

### Go

Install:

```sh
go get github.com/StevenMaude/win10-urls
```

and then run `win10-urls`.

## Why?

For whatever reason, Microsoft make it unintentionally difficult for
non-MSDN subscriber users to directly download an ISO of Windows 10.

Visiting the [Windows 10 download
URL](https://www.microsoft.com/en-us/software-download/windows10) on a
browser in Windows points you to the Media Creation tool, which requires
a working Windows install.

However, if you're not on Windows, visiting that URL redirects you
[elsewhere](https://www.microsoft.com/en-us/software-download/windows10ISO)
(note that this URL will just redirect you back to the /windows10 URL if
you're on Windows) and they tell you:

> You’ve been routed to this page because the operating system you’re
> using won’t support the Windows 10 media creation tool and we want to
> make sure you can download Windows 10. To use the media creation tool,
> visit the Microsoft Software Download Windows 10 page from a Windows
> 7, Windows 8.1 or Windows 10 device.

> You can use this page to download a disc image (ISO file) that can be
> used to install or reinstall Windows 10. The image can also be used to
> create installation media using a USB flash drive or DVD.

Alternatively, it may be possible to spoof your user agent in browser on
Windows to trick Microsoft to show you the ISO download page, but the
code I've written here is perhaps a simpler fix.

Note that it is possible to download the other language versions of
Windows 10 this way. There's nothing special about the English version.
It just happens to be the one I'd use.
