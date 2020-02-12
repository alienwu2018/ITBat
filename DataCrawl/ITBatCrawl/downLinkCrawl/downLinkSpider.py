import requests
import pymysql
import configparser
from lxml import etree


class downLinkSpider:
    conn = None
    cur = None
    falseUrl = []

    def connMysql(self):
        cp = configparser.ConfigParser()
        cp.read('../scrapy.cfg')
        self.conn = pymysql.connect(
            host=cp.get('mysql', 'HOST'),
            user=cp.get('mysql', 'USER'),
            passwd=cp.get('mysql', 'PWD'),
            db=cp.get('mysql', 'DB'),
            port=cp.getint('mysql', 'PORT'),
            charset=cp.get("mysql", 'CHARSET'),
        )
        self.cur = self.conn.cursor()
        return

    def Crawl(self,bid,url):
        proxy_dict = {
            "http": "http://127.0.0.1:1080/",
            "https": "http://127.0.0.1:1080/"
        }
        try:
            resp = requests.get(url=url,proxies=proxy_dict).text
            html = etree.HTML(resp)
        except Exception as e:
            print(e)
            self.falseUrl.append(url)
            return
        try:
            url = html.xpath('//a[@class="text-danger alert-link"]/@href')[1]
        except:
            url = ''
        try:
            code = html.xpath('//div[@class="alert alert-success mt-3"]/p/text()')[2][6:]
        except:
            code = ''
        data = (bid, url, code)
        try:
            self.cur.execute("insert into `downlink` (`bid`,`purl`,`pwd`)values {0}".format(data))
            self.conn.commit()
        except Exception as e:
            print(e)
        return



if __name__ == '__main__':
    obj = downLinkSpider()
    obj.connMysql()
    conn,cur = obj.conn,obj.cur
    cur.execute("select `id`,`YunPanUrl` from `book`")
    tpool = []
    for i in cur.fetchall():
        if i[1]!='':
            obj.Crawl(i[0],"https://itpanda.net"+i[1])
    conn.commit()
    conn.close()
    print(obj.falseUrl)
