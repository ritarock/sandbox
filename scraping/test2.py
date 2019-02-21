import urllib.request as request
from bs4 import BeautifulSoup

URL = "http://www.nikkei.com/markets/kabu/"


def main():
    value = ""
    html = request.urlopen(URL)
    soup = BeautifulSoup(html, "html.parser")
    p = soup.select_one("#CONTENTS_MARROW > div.mk-top_stock_average.cmn-clearfix > div.cmn-clearfix > div.mkc-guidepost > div.mkc-prices > span.mkc-stock_prices").text
    print(p)


if __name__ == '__main__':
    main()
