package fetcher

import (
	"io"
	"net/http"
)

// FetchHtml 从url获取html
//
// 参数:
//
//   - url: 要获取的url
//
// 返回值:
//
//   - string: 获取到的html文本
//   - error: 错误信息
//
// 示例:
//
// doc, err := fetcher.FetchHtml("https://www.baidu.com")
//
//	if err != nil {
//		panic(err)
//	}
//
// fmt.Println(doc)
func FetchHtml(url string) (string, error) {
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		return "", err
	}
	req.Header.Add("Cookie", "bid=sS4jUqxnFX8; viewed=\"36081529\"; dbcl2=\"181192123:84PKSkyYuFg\"; ck=H94U; ll=\"108296\"; push_noty_num=0; push_doumail_num=0; bid=vSgfKATO8mE")
	// 豆瓣电影网站会检查User-Agent，因此需要设置一个合法的User-Agent
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36")

	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}
