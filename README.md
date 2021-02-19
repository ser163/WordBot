# WordBot
## 项目介绍
  WordBot是一个go语言版本的随机单词生成器。

   有此项目诸君不用重复造轮子。
## 说明
### 示例
```go
package main

import (
	"fmt"
	"github.com/ser163/WordBot/generate"
)

func main() {
	for i := 0; i < 5; i++ {
		w, err := generate.GenRandomWorld(0, "none")
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println(w.Chance, w.Word)
	}
}
```
### 参数说明
```go
func GenRandomWorld(length int, model string) (WorldList, error)
```
- length 为要生成的字符串长度。0时随机单词长度为1-19之间。指定小于19的数字可以获取定长的字符串。
- model 为生成的字符串大小写模式
  1. none 为小写模式。与GenRandomNone()功能相同
  2. upper 大写模式。与GenRandomUpper()功能相同
  3. lower 小写模式，与none相同，与GenRandomLower()功能相同
  4. mix 大小写混合模式，与GenRandomMix()功能相同
  5. title 首字母大写模式，与GenRandomTitle()功能相同

## 感谢
此项目算法参考： https://namefull.github.io

namefull--一个js版本随机单词生成器
