import urllib.request as request
from bs4 import BeautifulSoup

URL = "http://www.nikkei.com/markets/kabu/"


def main():
    value = ""

    html = request.urlopen(URL)
    soup = BeautifulSoup(html, "html.parser")
    span = soup.find_all("span")

    for tag in span:
        try:
            string = tag.get("class")[0]
            if string in "mkc-stock_prices":
                value = tag.string
                break
        except:
            pass

    print(value)

if __name__ == '__main__':
    main()
