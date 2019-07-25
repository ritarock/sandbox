import urllib.request
from bs4 import BeautifulSoup
import sys

DOWNLOAD_PATH = "./dl"

class Soup:
    def __init__(self, url):
        self.html = urllib.request.urlopen(url=url)
        self.soup = BeautifulSoup(self.html, 'html.parser')


def get_index_page(url):
    html_soup = Soup(url)
    elements = html_soup.soup.select('.article-header')

    index_url = []
    for element in elements:
        for link in element.select('a'):
            index_url.append(link.get('href'))
    return index_url[0]


def get_page_list(index_page):
    pages = [index_page]
    url = index_page
    flg = True

    while flg:
        html_soup = Soup(url)
        elements = html_soup.soup.select('.prev')

        for element in elements:
            for prev in element:
                if prev.get('href') != 'None':
                    pages.append(prev.get('href'))
                    url = prev.get('href')
                else:
                    flg = False

    return pages

def get_images(page_list):
    for page in page_list:
        html_soup = Soup(page)
        elements = html_soup.soup.select('.article-body')

        for element in elements:
            for image in element.select('img'):
                download_image(image.get('src').rstrip('/small'))


def download_image(img):
    web_file = urllib.request.urlopen(img).read()
    with open("{}/{}.jpeg".format(DOWNLOAD_PATH,
                img.split("/")[-1]), mode='wb') as local_file:
        local_file.write(web_file)


def main(arg):
    index_page = get_index_page(arg)
    page_list = get_page_list(index_page)
    get_images(page_list)
    print("finish")


if __name__ == '__main__':
    main(sys.argv[1:][0])

