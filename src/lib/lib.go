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
	indexMd, _ := ioutil.ReadFile("pages/index.md")
	md2, _ := ioutil.ReadFile("pages/articles/202208/20220801.md")
	md3, _ := ioutil.ReadFile("pages/articles/202208/20220802.md")
	indexMd = iterappend(indexMd, md2, md3)
	articleDirs, _ := ioutil.ReadDir("pages/articles")
	for _, f := range articleDirs {
		fmt.Println(f.Name())
		dirs, _ := ioutil.ReadDir("pages/articles/" + f.Name())
		for _, fa := range dirs {
			md, _ := ioutil.ReadFile("pages/articles/" + f.Name() + "/" + fa.Name())
			indexMd = append(indexMd, md...)
			indexMd = append(indexMd, []byte{10}...)
			fmt.Println(fa.Name())
		}
	}

	fmt.Println(indexMd)
	fmt.Println(md2)

	// Markdownをmustache HTMLに変換
	mustacheHtml := string(blackfriday.Run(indexMd))
	fmt.Println(mustacheHtml)
	// mustache HTMLからHTMLを生成
	html, _ := mustache.Render(mustacheHtml, context...)

	// 描画
	tmpl := template.Must(template.ParseFiles("template/md.html"))
	tmpl.ExecuteTemplate(w, "md", template.HTML(html))
}

func iterappend(b ...[]byte) []byte {
	var result []byte
	for _, v := range b {
		result = append(result, v...)
		s := []byte{10}
		result = append(result, s...)
	}
	return result
}
