package main

import (
	"douban_movies/config"
	"douban_movies/internal/fetcher"
	"douban_movies/internal/file"
	"douban_movies/internal/parser"
	"fmt"
	"os"
)

func main() {
	// 获取html文本
	htmlString, err := fetcher.FetchHtml(config.GetConfig().Url)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 解析html文本
	doc, err := parser.ToHtmlNode(htmlString)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 获取电影信息
	movies, err := parser.GetMovie(doc)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 写入csv文件
	err = os.Mkdir(config.GetConfig().OutputPath, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = file.WriteToCsv(fmt.Sprintf("%v/movies.csv", config.GetConfig().OutputPath), movies)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = file.SavePicturesAndPages(movies)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("写入成功")
}
