package handler

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/gin"
	"io"
	"lyp-go/model"
	"lyp-go/output"
	"net/http"
	"regexp"
	"strings"
)

func GitHubTrendingHandler(c *gin.Context) {

	url := "https://api.lyp.ink/proxy/github.com"

	// 发送 HTTP 请求
	res, err := http.Get(url + "/trending")
	if err != nil {
		c.JSON(http.StatusOK, output.Err(model.ErrorCode, fmt.Sprintf("Failed to fetch data: %s", err), nil))
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(res.Body)

	// 检查响应状态码
	if res.StatusCode != 200 {
		c.JSON(http.StatusOK, output.Err(model.ErrorCode, fmt.Sprintf("Failed to fetch data: status code %d", res.StatusCode), nil))
	}

	// 使用 goquery 解析 HTML
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		c.JSON(http.StatusOK, output.Err(model.ErrorCode, fmt.Sprintf("Failed to parse HTML: %s", err), nil))
	}

	var collection []map[string]interface{}
	// 查找所有的仓库条目
	doc.Find("article.Box-row").Each(func(i int, s *goquery.Selection) {
		// 提取仓库名称
		repoName := strings.TrimSpace(s.Find("h2").Text())
		replacer := strings.NewReplacer("\n", "", "\t", "")
		repoName = replacer.Replace(repoName)
		// 编译正则表达式，匹配多个空格
		re := regexp.MustCompile(`\s+`)
		// 将多个空格替换为单个空格
		repoName = re.ReplaceAllString(repoName, " ")

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
		forks := strings.TrimSpace(s.Find("a[href*='forks']").Text())

		href, exists := s.Find("a[class*='Link']").Attr("href")
		if exists {
			href = url + href
		}

		replacer = strings.NewReplacer(",", "")
		collection = append(collection, map[string]interface{}{
			"repository":  repoName,
			"description": repoDescription,
			"language":    repoLanguage,
			"stars":       replacer.Replace(stars),
			"forks":       replacer.Replace(forks),
			"href":        href,
		})
	})

	c.JSON(http.StatusOK, output.Suc("trending", collection))
}
