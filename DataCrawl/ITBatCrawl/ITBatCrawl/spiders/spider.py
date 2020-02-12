# -*- coding: utf-8 -*-
import scrapy
from ..items import ItbatcrawlItem
from urllib import parse
import re

class SpiderSpider(scrapy.Spider):
    name = 'spider'
    allowed_domains = ['itpanda.net']
    start_urls = ['https://itpanda.net/']
    pattern = "(.+)\/(.+)"

    def parse(self, response):
        CategoryUrl=response.xpath('//li[@class="nav-item"]/a/@href').extract()
        for url in CategoryUrl:
            # print(url)
            yield scrapy.Request(url="https://itpanda.net"+url,meta={"mainUrl":url},callback=self.parseBookUrl,dont_filter=True)

    #获取书籍的主页链接
    def parseBookUrl(self,response):
        # 书籍url列表
        bookslist = response.xpath('//li[@class="media mb-3"]/div[@class="media-body"]/h5/a/@href').extract()
        for url in bookslist:
            yield scrapy.Request(url="https://itpanda.net"+url,meta={"mainUrl":response.meta["mainUrl"]},callback=self.parseBookData)
        #爬取下一页
        nextPage = response.xpath('//li[@class="page-item"]/a/@href').extract()
        for i in nextPage:
           yield scrapy.Request(url="https://itpanda.net"+i,meta={"mainUrl":response.meta["mainUrl"]},callback=self.parseBookUrl)


    #获取书籍的书籍
    def parseBookData(self,response):
        #书名
        bookName = response.xpath('//h4[@class="my-3"]/text()').extract()[0]
        #书籍链接
        bookPng = response.xpath('//img[@class="mr-3"]/@src').extract()[0]
        count=1
        #作者
        try:
            author = response.xpath('//div[@class="media-body"]/p[{0}]/text()'.format(count)).extract()[0]
            count+=1
        except:
            author=""
            count+=1
        #出版社
        try:
            press = response.xpath('//div[@class="media-body"]/p[{0}]/text()'.format(count)).extract()[0]
            count+=1
        except:
            press =""
            count+=1
        # #出版时间
        try:
            name = response.xpath('//div[@class="media-body"]/p[{0}]/span/text()'.format(count)).extract()[0]
            if name == "副标题:":
                publishtime = response.xpath('//div[@class="media-body"]/p[{0}]/text()'.format(count+1)).extract()[0]
                count+=2
            else:
                publishtime = response.xpath('//div[@class="media-body"]/p[{0}]/text()'.format(count)).extract()[0]
                count+=1
        except:
            publishtime = ""
            count+=1
        #页数
        try:
            pages = response.xpath('//div[@class="media-body"]/p[{0}]/text()'.format(count)).extract()[0]
            count+=1
        except:
            pages = 0
            count+=1
        #ISBN
        try:
            isbn = response.xpath('//div[@class="media-body"]/p[{0}]/text()'.format(count)).extract()[0]
            count+=1
        except:
            isbn = ""
            count+=1
        #豆瓣链接
        try:
            dburl = response.xpath('//div[@class="media-body"]/p[{0}]/a/@href'.format(count)).extract()[0]
            count+=1
        except:
            dburl = ""
            count+=1
        #豆瓣评分
        try:
            dbscore = response.xpath('//div[@class="media-body"]/p[{0}]/text()'.format(count)).extract()[0]
        except:
            dbscore = 0.0
        #内容简介
        try:
            desc = response.xpath('//p[@style="white-space: pre-line;"]/text()').extract()[0]
        except:
            desc = ""
        try:
            double = response.xpath('//a[@class="text-danger mr-2"]/@href').extract()
        except:
            double = ""
        #购买链接
        try:
            burl = double[0]
        except:
            burl = ""
        #下载链接
        try:
            durl = double[1]
        except:
            durl = ""
        #书籍的大类别
        target = parse.urlparse(url=response.meta["mainUrl"]).path[15:]
        pattern = "([^/]+)\/?(.*)"
        tmp =re.compile(pattern).findall(target)
        bigCategory =  tmp[0][0]
        #书籍的小类别
        try:
            smallCategory = tmp[0][1]
        except:
            smallCategory = ""

        item = ItbatcrawlItem()
        item['BookName'] = bookName
        item['PngUrl'] = bookPng
        item['Author'] = author
        item['Press'] = press
        item['PublishTime'] = publishtime
        item['Pages'] = int(pages)
        item['ISBN'] = isbn
        item['DouBanUrl'] = dburl
        item['DouBanScore'] = float(dbscore[:-1])
        item['Content'] = desc
        item['BuyUrl'] = burl
        item['YunPanUrl'] = durl
        item['BigCategory'] = bigCategory
        item['SmallCategory'] = smallCategory
        yield item

