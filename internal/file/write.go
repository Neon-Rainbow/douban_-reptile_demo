package file

import (
	"douban_movies/config"
	"douban_movies/internal/fetcher"
	"douban_movies/model"
	"encoding/csv"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"sync"
)

// WriteToCsv 写入csv文件
//
// 参数:
//
//   - path: csv文件路径
//   - data: 电影信息
//
// 返回值:
//
//   - error: 错误信息
//
// 示例:
//
//	err := file.WriteToCsv("movies.file", movies)
//
//	if err != nil {
//		panic(err)
//	}
func WriteToCsv(path string, data []model.Movie) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// 写入表头
	err = writer.Write([]string{"标题", "评分", "年份", "时长", "地区", "导演", "演员", "类型", "海报url", "电影页面url"})
	if err != nil {
		return err
	}

	// 写入数据
	for _, movie := range data {
		err = writer.Write([]string{
			movie.Title,
			movie.Score,
			movie.Year,
			movie.Duration,
			movie.Region,
			movie.Director,
			movie.Actors,
			movie.Category,
			movie.PosterUrl,
			movie.MoviePageUrl,
		})
		if err != nil {
			return err
		}
	}
	return nil
}

var (
	ch = make(chan model.Movie, 10)
	wg sync.WaitGroup
)

// SavePicturesAndPages 保存图片和电影详情页
//
// 参数:
//   - movies: 电影信息
//
// 返回值:
//   - error: 错误信息
//
// 示例:
//
//	err := file.SavePicturesAndPages(movies)
//
//	if err != nil {
//		panic(err)
//	}
func SavePicturesAndPages(movies []model.Movie) error {
	go func() {
		save(config.GetConfig().Goroutines)
	}()

	for _, movie := range movies {
		wg.Add(1)
		ch <- movie
	}

	close(ch)
	return nil
}

func save(n int) {
	// 使用 n 个 goroutine 从 ch 中读取数据并保存图片和电影详情页
	for i := 0; i < n; i++ {
		go func() {
			for movie := range ch {
				// 这里使用匿名函数来保证每个goroutine在处理完一个电影后都能调用wg.Done()来通知主goroutine
				func() {
					defer wg.Done()
					err := os.Mkdir(fmt.Sprintf("output/%s", movie.Title), os.ModePerm)
					if err != nil {
						fmt.Println(err)
					}
					// 保存图片
					err = downloadImage(movie.PosterUrl, fmt.Sprintf("output/%s/海报.jpg", movie.Title))
					if err != nil {
						fmt.Println(err)
					}

					// 保存电影详情页
					err = saveMoviePage(movie.MoviePageUrl, fmt.Sprintf("output/%s/详情页.html", movie.Title))
					if err != nil {
						fmt.Println(err)
					}
				}()
			}
		}()
	}
}

// downloadImage 根据图片 URL 下载并保存到本地
//
// 参数:
//   - imageURL: 图片地址
//   - savePath: 图片保存位置
//
// 返回值:
//   - error: 错误信息
func downloadImage(imageURL, savePath string) error {
	// 访问豆瓣的图片无需 Cookie,只需要发送 GET 请求即可,因此直接创建 HTTP GET 请求
	resp, err := http.Get(imageURL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// 检查 HTTP 响应状态
	if resp.StatusCode != http.StatusOK {
		return err
	}

	// 创建保存文件
	file, err := os.Create(savePath)
	if err != nil {
		return err
	}
	defer file.Close()

	// 将响应的内容写入文件
	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return err
	}

	return nil
}

// saveMoviePage 保存电影详情页
//
// 参数:
//   - url: 电影详情页地址
//   - savePath: 保存路径
//
// 返回值:
//   - error: 错误信息
func saveMoviePage(url string, savePath string) error {
	s, err := fetcher.FetchHtml(url)
	if err != nil {
		return err
	}
	file, err := os.Create(savePath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, strings.NewReader(s))
	if err != nil {
		return err
	}
	return nil
}
