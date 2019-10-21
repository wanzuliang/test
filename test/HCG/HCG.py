__author__ = 'shirokuma'
#coding=utf-8


# proxies = {
#   "http": "0.10.1.10:3128",
#   "https": "10.10.1.10:1080",
# }
#
# requests.get("http://example.org", proxies=proxies)

import urllib
from urllib import request
from urllib import error
import re
import requests
import os
import threading

user_agent = "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/51.0.2704.106 Safari/537.36"
headers = {'User-Agent': 'use_agent', 'Referer': 'http://www.dakemakura.com/kakoitiran.html'}


def down_pic():
    fp = open('pic/pic_number.txt', 'r')
    numall = fp.readlines()
    for num in numall:
        num = num.replace('\n', '')
        if os.path.isfile("pic/"+str(num)+".jpg"):
            print("可能存在"+str(num)+"文件")
        else:
            try:
                fp = open("pic/"+str(num)+".jpg", "wb")
                fp.close()
                print("从文件读取链接并获取"+str(num)+"图片 . . . . . . ")
                pic = requests.get('http://www.dakemakura.com/image/zentai/'+str(num)+'-2.jpg', headers=headers)
                print("获取图片完成。")
            except requests.HTTPError as e:
                print("读取源文件失败！错误是 ：" + e)
                break
            print(str(num)+".jpg 写入中.......")
            fp = open("pic/"+str(num)+".jpg", "wb")
            fp.write(pic.content)
            fp.close()
            print(str(num)+".jpg 写入成功！")
            print("--------------------------------------------------------------------"+num+".jpg")


def down(num):
    urllib.request.urlretrieve('http://www.dakemakura.com/image/zentai/'+str(num)+'-2.jpg', "pic/"+str(num)+".jpg")


def getLink():
    # url = "http://www.dakemakura.com/top.html"
    url = 'http://www.dakemakura.com/kakoitiran.html'
    # numstr = '<A href="(.*?)_shousai.html">'
    numstr = '<A href="(.*?)_shousai.html"><IMG src="image/cart/.*?-1.jpg" width="63" height="75" border="0">'

    print("获取图片链接 . . . . . .")

    request = urllib.request.Request(url)
    response = urllib.request.urlopen(request)
    html = response.read().decode('gbk')
    recompile = re.compile(numstr)
    allpic = re.findall(recompile, html)

    print("链接获取完成！")
    for i in allpic:
        f = open('pic/pic_number.txt', 'a')
        f.write(str(i)+'\n')
        f.close()
    print("链接已保存至 pic/pic_number.txt ")
getLink()
# down_pic()
