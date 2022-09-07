package lib

import (
	"fmt"
	"html/template"
	"io"
	"io/ioutil"

	"github.com/cbroglie/mustache"
	"github.com/russross/blackfriday/v2"
)

func IndexRender(w io.Writer, context ...interface{}) {
	indexMd, _ := ioutil.ReadFile("src/pages/index.md")
	md2, _ := ioutil.ReadFile("src/pages/articles/202208/20220801.md")
	md3, _ := ioutil.ReadFile("src/pages/articles/202208/20220802.md")
	indexMd = multiBytesAppend(indexMd, md2, md3)
	articleDirs, _ := ioutil.ReadDir("src/pages/articles")
	for _, f := range articleDirs {
		dirs, _ := ioutil.ReadDir("src/pages/articles/" + f.Name())
		for _, fa := range dirs {
			md, _ := ioutil.ReadFile("src/pages/articles/" + f.Name() + "/" + fa.Name())
			indexMd = multiBytesAppend(indexMd, md)
			fmt.Println(fa.Name())
		}
	}

	// Markdownをmustache HTMLに変換
	mustacheHtml := string(blackfriday.Run(indexMd))
	// mustache HTMLからHTMLを生成
	html, _ := mustache.Render(mustacheHtml, context...)

	// 描画
	tmpl := template.Must(template.ParseFiles("src/template/md.html"))
	tmpl.ExecuteTemplate(w, "md", template.HTML(html))
}

func multiBytesAppend(b ...[]byte) []byte {
	var result []byte
	for _, v := range b {
		result = append(result, v...)
		lfc := []byte{10}
		result = append(result, lfc...)
	}
	return result
}
