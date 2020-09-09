from bs4 import BeautifulSoup
import urllib.request

url = "https://github.com/microsoft/PowerToys"
url1 = "https://www.microsoft.com/zh-cn/software-download/windows10"

html_page = urllib.request.urlopen()
soup = BeautifulSoup(html_page, "html.parser")
for link in soup.findAll('a'):
    print(link.get('href'))