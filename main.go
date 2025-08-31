package main

import (
	"crypto/rand"
	"fmt"
	"io/ioutil"
	"math/big"
	"os"
	"path/filepath"
	"strconv"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// 生成随机字符串
func randomString(length int) string {
	result := make([]byte, length)
	for i := 0; i < length; i++ {
		n, _ := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		result[i] = charset[n.Int64()]
	}
	return string(result)
}

// 打印帮助信息
func printHelp() {
	exe := filepath.Base(os.Args[0])
	fmt.Printf(`用法:
1. 直接运行: %s
   - 打印帮助信息并生成一个随机32位字符串

2. %s <数量>
   - 生成指定数量的随机字符串 (长度默认为32)

3. %s <数量> <长度>
   - 生成指定数量、指定长度的随机字符串

4. %s <数量> <长度> <前缀>
   - 生成指定数量、指定长度的随机字符串，且以指定前缀开头

5. %s <...参数...> -o <文件路径>
   - 将生成的字符串写入到文件(覆盖)
`, exe, exe, exe, exe, exe)
}

func main() {
	args := os.Args[1:]

	// 无参数
	if len(args) == 0 {
		printHelp()
		fmt.Println("示例输出:", randomString(32))
		return
	}

	// 判断是否有 -o 参数
	outputPath := ""
	for i, arg := range args {
		if arg == "-o" && i+1 < len(args) {
			outputPath = args[i+1]
			args = append(args[:i], args[i+2:]...) // 移除 -o 和路径
			break
		}
	}

	count := 1
	length := 32
	prefix := ""

	// 解析参数
	if len(args) >= 1 {
		if n, err := strconv.Atoi(args[0]); err == nil {
			count = n
		} else {
			fmt.Println("错误: 第一个参数必须是数字")
			return
		}
	}
	if len(args) >= 2 {
		if n, err := strconv.Atoi(args[1]); err == nil {
			length = n
		} else {
			fmt.Println("错误: 第二个参数必须是数字")
			return
		}
	}
	if len(args) >= 3 {
		prefix = args[2]
		if len(prefix) > length {
			fmt.Println("错误: 前缀长度不能大于总长度")
			return
		}
	}

	// 生成结果
	results := make([]string, 0, count)
	for i := 0; i < count; i++ {
		remain := length - len(prefix)
		if remain < 0 {
			remain = 0
		}
		str := prefix + randomString(remain)
		results = append(results, str)
	}

	// 输出到控制台或文件
	if outputPath != "" {
		content := ""
		for _, s := range results {
			content += s + "\n"
		}
		if err := ioutil.WriteFile(outputPath, []byte(content), 0644); err != nil {
			fmt.Println("写入文件失败:", err)
			return
		}
		fmt.Println("已写入文件:", outputPath)
	} else {
		for _, s := range results {
			fmt.Println(s)
		}
	}
}
