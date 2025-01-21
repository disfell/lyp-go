package tests

import (
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/PuerkitoBio/goquery"
)

// TestFetchGitHubTrending 爬取 GitHub Trending 页面
func TestFetchGitHubTrending(t *testing.T) {

	url := "https://github.com/trending"

	// 发送 HTTP 请求
	res, err := http.Get(url)
	if err != nil {
		t.Fatal("Failed to fetch data: ", err)
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(res.Body)

	// 检查响应状态码
	if res.StatusCode != 200 {
		t.Fatalf("Failed to fetch data: status code %d", res.StatusCode)
	}

	// 使用 goquery 解析 HTML
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		t.Fatal("Failed to parse HTML: ", err)
	}

	// 查找所有的仓库条目
	doc.Find("article.Box-row").Each(func(i int, s *goquery.Selection) {
		// 提取仓库名称
		repoName := strings.TrimSpace(s.Find("h2").First().Find("a").Text())

		// 提取仓库描述
		repoDescription := strings.TrimSpace(s.Find("p.col-9").Text())
		if repoDescription == "" {
			repoDescription = "No description"
		}

		// 提取编程语言
		repoLanguage := strings.TrimSpace(s.Find("span[itemprop='programmingLanguage']").Text())
		if repoLanguage == "" {
			repoLanguage = "Unknown"
		}

		// 提取 stars 和 forks 数量
		stars := strings.TrimSpace(s.Find("a[href*='stargazers']").Text())
		forks := strings.TrimSpace(s.Find("a[href*='network/members']").Text())

		// 打印仓库信息
		t.Logf("Repository: %s\n", repoName)
		t.Logf("Description: %s\n", repoDescription)
		t.Logf("Language: %s\n", repoLanguage)
		t.Logf("Stars: %s\n", stars)
		t.Logf("Forks: %s\n", forks)
		t.Logf(strings.Repeat("-", 50))
	})
}
