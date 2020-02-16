package util

import "ITBat/model/book"

//将书籍类型进行分类
//ex:{front:{js,css,html},back{java,python}}
func DoCategory(books []book.Book) map[string][]string {
	var result = make(map[string][]string)
	flag := ""
	for _, book := range books {
		flag = book.BigCategory
		var tmp []string
		for _, book := range books {
			if book.BigCategory == flag {
				tmp = append(tmp, book.SmallCategory)
			}
		}
		result[book.BigCategory] = tmp
	}
	return result
}
