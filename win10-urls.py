import lxml.html
import requests


def main():
    """
    Displays URLs for Windows 10 ISO download from Microsoft.

    https://www.microsoft.com/en-us/software-download/windows10
    actually gives you 24 hour HTTPS URLs of the Windows 10 ISO, if you're
    not using Windows. It redirects you to a URL ending /windows10ISO
    
    (If on Windows, the /windows10ISO URL.)

    The URL used in the request is for the Windows 10 English release.

    Alternatively, it may be possible to spoof your user agent in browser to
    trick Microsoft to show you the ISO download page.
    """
    r = requests.post('https://www.microsoft.com/en-us/api/'
                      'controls/contentinclude/html?'
                      'pageId=cfa9e580-a81e-4a4b-a846-7b21bf4e2e5b'
                      '&host=www.microsoft.com&'
                      'segments=software-download,windows10ISO'
                      '&query=&action=GetProductDownloadLinksBySku'
                      '&skuId=3864&language=English&sdVersion=2')
    root = lxml.html.fromstring(r.text)
    for url in root.xpath('//div[@class]/a/@href'):
        print(url)


if __name__ == '__main__':
    main()
