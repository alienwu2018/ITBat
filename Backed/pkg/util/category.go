package util

import (
	"ITBat/model/book"
)

type Tree struct {
	Label    string      `json:"label"`
	Children []interface{} `json:"children"`
}

//将书籍类型进行分类
//ex:{front:{js,css,html},back{java,python}}
func DoCategory(books []book.Book) []Tree {
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
	var categories []Tree
	count := 0
	for k, _ := range result {
		var t Tree
		t.Label = k
		for _, v := range result[k] {
			if len(v) > 0 {
				t.Children = append(t.Children, map[string]string{"label": v})
			}
		}
		count++
		categories = append(categories, t)
	}
	//排序
	for i := 0; i < len(categories)-1; i++ {
		for j := 0; j < len(categories)-1-i; j++ {
			if len(categories[j].Children) < len(categories[j+1].Children) {
				categories[j], categories[j+1] = categories[j+1], categories[j]
			}
		}
	}
	return categories
}
