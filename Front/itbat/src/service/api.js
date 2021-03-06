var baseUrl = "http://47.100.93.236:8000/api/v1";

export const apiurl = {
    BaseUrl : baseUrl,

    //主页接口
    index : baseUrl+"/index",

    //获取书籍分类总数接口
    docategories : baseUrl+"/categories",

    //查找接口
    search : baseUrl+"/search",

    //书籍的分类接口
    category : baseUrl+"/books/category/",

    //书籍数据接口
    book : baseUrl+"/book",

    //按评分排序接口
    score : baseUrl+"/category",

    //云盘下载链接接口

    yun : baseUrl+"/book/",

    //推荐书籍接口
    recommend : baseUrl+"/book/",
}


