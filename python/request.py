import requests

home = requests.get("https://www.louisvuitton.cn/zhs-cn/homepage")
print(home.status_code)
print(home.text)
html = requests.get("https://www.louisvuitton.cn/zhs-cn/products/pochette-metis-monogram-nvprod2130176v")
print(html.status_code)
print(html.text)