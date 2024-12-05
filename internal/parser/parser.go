package parser

import (
	"douban_movies/model"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// ToHtmlNode 将html文本转换为html.Node
//
// 参数:
//
//   - htmlString: html文本
//
// 返回值:
//
//   - *html.Node: html.Node
//   - error: 错误信息
//
// 示例:
//
// node, err := parser.ToHtmlNode("<html><body><h1>hello world</h1></body></html>")
//
//	if err != nil {
//		panic(err)
//	}
func ToHtmlNode(htmlString string) (doc *goquery.Document, err error) {
	doc, err = goquery.NewDocumentFromReader(strings.NewReader(htmlString))
	if err != nil {
		return nil, err
	}
	return doc, nil
}

// GetMovie 从html.Node中获取电影信息
// 电影信息包括：标题、评分、年份、时长、地区、导演、演员、类型、海报url、电影页面url
//
// 参数:
//
//   - doc: *goquery.Document
//
// 返回值:
//
//   - []model.Movie: 电影信息
//   - error: 错误信息
//
// 示例:
//
// movies, err := parser.GetMovie(doc)
//
//	if err != nil {
//		panic(err)
//	}
//	fmt.Println(movies)
func GetMovie(doc *goquery.Document) ([]model.Movie, error) {
	movies := make([]model.Movie, 0)

	doc.Find(".list-item").Each(func(i int, s *goquery.Selection) {
		movie := model.Movie{}
		movie.Title = s.AttrOr("data-title", "")
		movie.Score = s.AttrOr("data-score", "")
		movie.Year = s.AttrOr("data-release", "")
		movie.Duration = s.AttrOr("data-duration", "")
		movie.Region = s.AttrOr("data-region", "")
		movie.Director = s.AttrOr("data-director", "")
		movie.Actors = s.AttrOr("data-actors", "")
		movie.Category = s.AttrOr("data-category", "")
		movie.PosterUrl = s.Find("img").AttrOr("src", "")
		movie.MoviePageUrl = s.Find("a").AttrOr("href", "")
		movies = append(movies, movie)
	})

	doc.Find(".list-item-hidden").Each(func(i int, s *goquery.Selection) {
		movie := model.Movie{}
		movie.Title = s.AttrOr("data-title", "")
		movie.Score = s.AttrOr("data-score", "")
		movie.Year = s.AttrOr("data-release", "")
		movie.Duration = s.AttrOr("data-duration", "")
		movie.Region = s.AttrOr("data-region", "")
		movie.Director = s.AttrOr("data-director", "")
		movie.Actors = s.AttrOr("data-actors", "")
		movie.Category = s.AttrOr("data-category", "")
		movie.PosterUrl = s.Find("img").AttrOr("src", "")
		movie.MoviePageUrl = s.Find("a").AttrOr("href", "")
		movies = append(movies, movie)
	})

	return movies, nil
}
