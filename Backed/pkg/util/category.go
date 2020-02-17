package util

import (
	"ITBat/model/book"
	"ITBat/pkg/e"
)

type Tree struct {
	Label    string              `json:"label"`
	Children []map[string]string `json:"children"`
}

//将书籍类型进行分类并获得前端真正显示的名称
//ex:category{front:{js,css,html},back{java,python}}
//ex:trueValue{前端{xxx,xxx}}
func DoCategory(books []book.Book) (relation map[string]string) {
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
	//复制一个categories拿来做类名
	var trueName []Tree
	for _, i := range categories {
		label := i.Label
		var c []map[string]string
		for _, j := range i.Children {
			tmp := make(map[string]string)
			for k, v := range j {
				tmp[k] = v
			}
			c = append(c, tmp)
		}
		trueName = append(trueName, Tree{label, c})
	}
	var Relation = make(map[string]string)
	for i := 0; i < len(trueName); i++ {
		Relation[e.Mapper[trueName[i].Label]] = trueName[i].Label
		trueName[i].Label = e.Mapper[trueName[i].Label]
		for j := 0; j < len(trueName[i].Children); j++ {
			Relation[e.Mapper[trueName[i].Children[j]["label"]]] = trueName[i].Children[j]["label"]
			trueName[i].Children[j]["label"] = e.Mapper[trueName[i].Children[j]["label"]]
		}
	}
	return Relation
}
