package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/nguyenthenguyen/docx"
	"github.com/yanyiwu/gojieba"
)

func main() {
	// 远程 .docx 文件的 URL
	url := "https://va-test-1.oss-cn-shanghai.aliyuncs.com/application/8es_mIVHZVhXTjjDqsP_U.docx"

	// 发送 HTTP GET 请求
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("无法下载文件: %v", err)
	}
	defer resp.Body.Close()

	// 读取响应体到内存
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("读取响应体失败: %v", err)
	}

	// 创建 io.ReaderAt 接口
	reader := strings.NewReader(string(data))

	// 使用 ReadDocxFromMemory 解析文档
	r, err := docx.ReadDocxFromMemory(reader, int64(len(data)))
	if err != nil {
		log.Fatalf("解析 .docx 文件失败: %v", err)
	}
	defer r.Close()

	// 获取文档内容
	doc := r.Editable()
	text := doc.GetContent()
	fmt.Println(text)
	return

	// 初始化 gojieba 分词器
	x := gojieba.NewJieba()
	defer x.Free()

	// 使用 gojieba 进行分词
	words := x.Cut(text, true)

	// 统计英文单词数
	englishWords := 0
	for _, word := range words {
		if isEnglish(word) {
			englishWords++
		}
	}

	// 统计中文字符数
	chineseChars := 0
	for _, word := range words {
		if isChinese(word) {
			chineseChars += len(word)
		}
	}

	// 输出统计结果
	fmt.Printf("英文单词数: %d\n", englishWords)
	fmt.Printf("中文字符数: %d\n", chineseChars)
}

// isEnglish 判断字符串是否为英文
func isEnglish(s string) bool {
	for _, r := range s {
		if !((r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z')) {
			return false
		}
	}
	return true
}

// isChinese 判断字符串是否为中文
func isChinese(s string) bool {
	for _, r := range s {
		if r < '\u4e00' || r > '\u9fff' {
			return false
		}
	}
	return true
}
