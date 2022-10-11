package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

type IndexData struct {
	Title string `json:"title"`
	Desc  string `json:"desc"`
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var indexData IndexData
	indexData.Title = "Xxj的博客"
	indexData.Desc = "不使用开发框架"
	t := template.New("index.html")
	dir, _ := os.Getwd()
	//访问博客首页模板的时候，因为有多个模板的嵌套，解析文件的时候需要将其涉及到的所有模板进行解析
	index := dir + "/template/index.html"
	home := dir + "/template/layout/home.html"
	header := dir + "/template/layout/header.html"
	footer := dir + "/template/layout/footer.html"
	personal := dir + "/template/layout/personal.html"
	post := dir + "/template/layout/post-list.html"
	pagination := dir + "/template/layout/pagination.html"
	t, _ = t.ParseFiles(index, home, header, footer, personal, post, pagination)
	//页面上涉及到的所有数据必须要有定义
	_ = t.Execute(w, indexData)
}

func main() {
	//程序入口，一个项目只能有一个入口
	//web程序，http协议 ip port
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	http.HandleFunc("/", index)

	if err := server.ListenAndServe(); err != nil { //监听这个端口
		log.Println(err)
	}

	str := "kk"
	var p *string = &str
	fmt.Print(p)
}
