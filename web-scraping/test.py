import requests
from bs4 import BeautifulSoup

def main():
    getres = ""
    # URLを取得
    url = "https://www.nikkei.com/"
    html = requests.get(url)

    # パース用のオブジェクトを作成
    soup = BeautifulSoup(html.text, "html.parser")

    # span要素のみを抽出
    span = soup.find_all("span")

    # span要素の中から取得目的のclass"m-miH01C_rate"を取得する
    for tag in span:
        try:
            # span要素からclassをpopする
            string_ = tag.get("class").pop(0)

            # 抽出したclassから"m-miH01C_rate"をチェックする
            if string_ == "m-miH01C_rate":
                # tagの文字列を取得
                getres = tag.string
                break
        except:
            pass

    print(getres)

if __name__ == '__main__':
    main()
