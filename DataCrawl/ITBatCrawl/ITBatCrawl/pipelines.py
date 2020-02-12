# -*- coding: utf-8 -*-

# Define your item pipelines here
#
# Don't forget to add your pipeline to the ITEM_PIPELINES setting
# See: https://docs.scrapy.org/en/latest/topics/item-pipeline.html

import configparser
import pymysql

class ItbatcrawlPipeline(object):

    def open_spider(self,spider):
        cp = configparser.ConfigParser()
        cp.read('scrapy.cfg')
        self.conn = pymysql.connect(
            host=cp.get('mysql', 'HOST'),
            user=cp.get('mysql', 'USER'),
            passwd=cp.get('mysql', 'PWD'),
            db=cp.get('mysql', 'DB'),
            port=cp.getint('mysql', 'PORT'),
            charset=cp.get("mysql", 'CHARSET'),
        )
        self.cur = self.conn.cursor()

        pass

    def close_spider(self,spider):
        self.conn.commit()
        self.conn.close()
        pass

    def process_item(self, item, spider):
        bookName=item['BookName']
        bookPng=item['PngUrl']
        author=item['Author']
        press=item['Press']
        publishtime=item['PublishTime']
        pages=item['Pages']
        isbn=item['ISBN']
        dburl=item['DouBanUrl']
        dbscore=item['DouBanScore']
        desc=item['Content']
        burl=item['BuyUrl']
        durl=item['YunPanUrl']
        bigCategory=item['BigCategory']
        smallCategory=item['SmallCategory']

        insert_data = (bookName,bookPng,author,press,publishtime,pages,isbn,dburl,dbscore,desc,
                       burl,durl,bigCategory,smallCategory)
        self.cur.execute("Insert into `book`(`BookName` ,`PngUrl`,`Author`,`Press`,`PublishTime`,`Pages`,`ISBN`,`DouBanUrl`,`DouBanScore`,`Content`,`BuyUrl`,`YunPanUrl`,`BigCategory`,`SmallCategory`)values {0}".format(insert_data))
        self.conn.commit()
        return item
