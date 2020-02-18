package book

import (
	"ITBat/model"
	"ITBat/pkg/logging"
	"bytes"
)

//主页面书本的信息
type MainPage struct {
	model.Model
	BookName    string  `gorm:"column:BookName" json:"book_name"`
	PngUrl      string  `gorm:"column:PngUrl" json:"png_url"`
	Author      string  `gorm:"column:Author" json:"autor"`
	DouBanScore float64 `gorm:"column:DouBanScore" json:"dou_ban_score"`
	Content     string  `gorm:"column:Content" json:"content"`
}

func (MainPage) TableName() string {
	return "book"
}

//书本详细信息
type Book struct {
	model.Model

	BookName      string  `gorm:"column:BookName" json:"book_name"`
	PngUrl        string  `gorm:"column:PngUrl" json:"png_url"`
	Author        string  `gorm:"column:Author" json:"autor"`
	Press         string  `gorm:"column:Press" json:"press"`
	PublishTime   string  `gorm:"column:PublishTime" json:"publish_time"`
	Pages         int     `gorm:"column:Pages" json:"pages"`
	ISBN          string  `gorm:"column:ISBN" json:"isbn"`
	DouBanUrl     string  `gorm:"column:DouBanUrl" json:"dou_ban_url"`
	DouBanScore   float64 `gorm:"column:DouBanScore" json:"dou_ban_score"`
	Content       string  `gorm:"column:Content" json:"content"`
	BuyUrl        string  `gorm:"column:BuyUrl" json:"buy_url"`
	YunPanUrl     string  `gorm:"column:YunPanUrl" json:"yun_pan_url"`
	BigCategory   string  `gorm:"column:BigCategory" json:"big_category"`
	SmallCategory string  `gorm:"column:SmallCategory" json:"small_category"`
}

func (Book) TableName() string {
	return "book"
}

//主页数据
func QueryIndex(pageNum, pageSize int, Isscore bool) (books []MainPage, err error) {
	if !Isscore {
		if err = model.DB.Debug().Model(Book{}).Offset(pageNum).Limit(pageSize).Find(&books).Error; err != nil {
			logging.DebugLog(err)
			return
		}
		return
	} else {
		if err = model.DB.Debug().Model(Book{}).Order("DouBanScore desc").Offset(pageNum).Limit(pageSize).Find(&books).Error; err != nil {
			logging.DebugLog(err)
			return
		}
		return
	}
}

//总条目
func BooksRow() (row int, err error) {
	if err = model.DB.Debug().Model(Book{}).Count(&row).Error; err != nil {
		logging.DebugLog(err)
		return
	}
	return
}

//大类数据
func QueryBookByBigCategory(pageNum int, pageSize int, BigCategory string) (books []MainPage, err error) {
	if err = model.DB.Debug().Model(Book{}).Where("BigCategory=? ", BigCategory).Order("DouBanScore desc").Offset(pageNum).Limit(pageSize).Find(&books).Error; err != nil {
		logging.DebugLog(err)
		return
	}
	return
}

//大类别总数
func BigCategoryRow(BigCategory string) (row int, err error) {
	if err = model.DB.Debug().Model(Book{}).Where("BigCategory=? ", BigCategory).Count(&row).Error; err != nil {
		logging.DebugLog(err)
		return
	}
	return
}

//小类数据
func QueryBookBySmallCategory(pageNum int, pageSize int, BigCategory, SmallCategory string) (books []MainPage, err error) {
	if err = model.DB.Debug().Model(Book{}).Where("BigCategory=? and SmallCategory=?", BigCategory, SmallCategory).Offset(pageNum).Limit(pageSize).Find(&books).Error; err != nil {
		logging.DebugLog(err)
		return
	}
	return
}

//小类别总数
func SmallCategoryRow(BigCategory, SmallCategory string) (row int, err error) {
	if err = model.DB.Debug().Model(Book{}).Where("BigCategory=? and SmallCategory=?", BigCategory, SmallCategory).Count(&row).Error; err != nil {
		logging.DebugLog(err)
		return
	}
	return
}

//按书名分页查询
func QueryBookByName(pageNum int, pageSize int, BookName string, isScore bool) (books []MainPage, err error) {
	var buffer bytes.Buffer
	buffer.WriteString("%")
	buffer.WriteString(BookName)
	buffer.WriteString("%")
	if !isScore {
		if err = model.DB.Debug().Model(Book{}).Where("BookName Like ?", buffer.String()).Offset(pageNum).Limit(pageSize).Find(&books).Error; err != nil {
			logging.DebugLog(err)
			return
		}
		return
	} else {
		if err = model.DB.Debug().Model(Book{}).Where("BookName Like ?", buffer.String()).Order("DouBanScore desc").Offset(pageNum).Limit(pageSize).Find(&books).Error; err != nil {
			logging.DebugLog(err)
			return
		}
		return
	}
}

func BooksNameRow(BookName string) (row int, err error) {
	var buffer bytes.Buffer
	buffer.WriteString("%")
	buffer.WriteString(BookName)
	buffer.WriteString("%")
	if err = model.DB.Debug().Model(Book{}).Where("BookName Like ?", buffer.String()).Count(&row).Error; err != nil {
		logging.DebugLog(err)
		return
	}
	return
}

//按豆瓣分数排序分页查询
func QueryBookByScore(pageNum int, pageSize int, BigCategory, SmallCategory string, hasSamll bool) (books []MainPage, err error) {
	if hasSamll {
		if err = model.DB.Debug().Model(Book{}).Where("BigCategory=? and SmallCategory=?", BigCategory, SmallCategory).Order("DouBanScore desc").Offset(pageNum).Limit(pageSize).Find(&books).Error; err != nil {
			logging.DebugLog(err)
			return
		}
		return
	} else {
		if err = model.DB.Debug().Model(Book{}).Where("BigCategory=?", BigCategory).Order("DouBanScore desc").Offset(pageNum).Limit(pageSize).Find(&books).Error; err != nil {
			logging.DebugLog(err)
			return
		}
		return
	}
}

//获取书本详细信息
func QueryBookById(BId int) (book Book, err error) {
	if err = model.DB.Debug().Model(Book{}).Where("Id = ?", BId).First(&book).Error; err != nil {
		logging.DebugLog(err)
		return
	}
	return
}

//查询书籍的所有分类
func QueryAllCategory() (book []Book, err error) {
	if err = model.DB.Raw("select `BigCategory`,`SmallCategory` from `book` group by `BigCategory`,`SmallCategory`").Scan(&book).Error; err != nil {
		logging.DebugLog(err)
		return
	}
	return
}

//推荐书籍
func RecommendBooks(bid int) (books []MainPage, err error) {
	//该书籍处于的小类
	var smallcategory string
	var bigcategory string
	var b Book
	if err = model.DB.Model(b).Select("BigCategory,SmallCategory").Where("id=?", bid).First(&b).Error; err != nil {
		logging.DebugLog(err)
		return
	}
	smallcategory = b.SmallCategory
	if smallcategory != "" {
		if err = model.DB.Model(Book{}).Where("SmallCategory =?", smallcategory).Order("RAND()").Limit(5).Find(&books).Error; err != nil {
			logging.DebugLog(err)
			return
		}
		return
	} else {
		bigcategory = b.BigCategory
		if err = model.DB.Model(Book{}).Where("BigCategory =?", bigcategory).Order("RAND()").Limit(5).Find(&books).Error; err != nil {
			logging.DebugLog(err)
			return
		}
		return
	}
}
