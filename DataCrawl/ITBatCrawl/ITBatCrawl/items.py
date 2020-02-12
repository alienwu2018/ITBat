# -*- coding: utf-8 -*-

# Define here the models for your scraped items
#
# See documentation in:
# https://docs.scrapy.org/en/latest/topics/items.html

import scrapy


class ItbatcrawlItem(scrapy.Item):
    # define the fields for your item here like:
    BookName = scrapy.Field()
    PngUrl = scrapy.Field()
    Author = scrapy.Field()
    Press = scrapy.Field()
    PublishTime = scrapy.Field()
    Pages = scrapy.Field()
    ISBN = scrapy.Field()
    DouBanUrl = scrapy.Field()
    DouBanScore = scrapy.Field()
    Content = scrapy.Field()
    BuyUrl = scrapy.Field()
    YunPanUrl = scrapy.Field()
    BigCategory = scrapy.Field()
    SmallCategory = scrapy.Field()
    pass
