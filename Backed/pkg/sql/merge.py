import pymysql
import configparser
import datetime

class TableMerge:
    def __init__(self):
        cp = configparser.ConfigParser()
        cp.read('../../setting/conf.ini')
        self.conn = pymysql.connect(
            host=cp.get('mysql', 'HOST'),
            user=cp.get('mysql', 'USER'),
            passwd=cp.get('mysql', 'PWD'),
            db=cp.get('mysql', 'DB'),
            port=cp.getint('mysql', 'PORT'),
            charset=cp.get("mysql", 'CHARSET'),
        )
    def mergeBook(self):
        cur = self.conn.cursor()
        cur.execute(
            "select `BookName` ,`PngUrl`,`Author`,`Press`,`PublishTime`,`Pages`,`ISBN`,`DouBanUrl`,`DouBanScore`,`Content`,`BuyUrl`,`YunPanUrl`,`BigCategory`,`SmallCategory` from book_crawl")
        for i in cur.fetchall():
            create_time = datetime.datetime.now().strftime("%Y-%m-%d %H:%M:%S")
            update_time = create_time
            i = [t for t in i]
            i.extend([create_time, update_time])
            data = tuple(i)
            print(data)
            cur.execute(
                "insert into book (`BookName` ,`PngUrl`,`Author`,`Press`,`PublishTime`,`Pages`,`ISBN`,`DouBanUrl`,`DouBanScore`,`Content`,`BuyUrl`,`YunPanUrl`,`BigCategory`,`SmallCategory`,`created_at`,`updated_at`)values {0}".format(
                    data))
            self.conn.commit()

    def mergeUrl(self):
        cur = self.conn.cursor()
        cur.execute(
            "select `bid`,`purl`,`pwd` from downlink_crawl")
        for i in cur.fetchall():
            create_time = datetime.datetime.now().strftime("%Y-%m-%d %H:%M:%S")
            update_time = create_time
            i = [t for t in i]
            i.extend([create_time, update_time])
            data = tuple(i)
            print(data)
            cur.execute(
                "insert into downlink (`bid`,`purl`,`pwd`,`created_at`,`updated_at`)values {0}".format(
                    data))
            self.conn.commit()

if __name__ == '__main__':
    # merge()
    obj = TableMerge()
    obj.mergeUrl()