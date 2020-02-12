import pymysql
import configparser
#将爬取到的数据进行清洗，爬YunPanUrl中的bookid修改
class DataCleaning:

    def __init__(self):
        cp = configparser.ConfigParser()
        cp.read('../DataCrawl/ITBatCrawl/scrapy.cfg')
        self.conn = pymysql.connect(
            host=cp.get('mysql', 'HOST'),
            user=cp.get('mysql', 'USER'),
            passwd=cp.get('mysql', 'PWD'),
            db=cp.get('mysql', 'DB'),
            port=cp.getint('mysql', 'PORT'),
            charset=cp.get("mysql", 'CHARSET'),
        )

    def cleaning(self):
        cur = self.conn.cursor()
        cur.execute("select `Id`,`YunPanUrl` from `book`")
        for i in cur.fetchall():
            if i[1]!='':
                tmp = '"/book/{0}/download/{1}"'.format(i[0],i[0])
                cur.execute("update  `book` set `YunPanUrl`={0}  where `Id`={1}".format(tmp,i[0]))
                self.conn.commit()

if __name__ == '__main__':
    obj = DataCleaning()
    obj.cleaning()
