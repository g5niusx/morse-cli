package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strings"
)

var morseMap = make(map[string]string)
var textMap = make(map[string]string)

func init() {

	bytes, e := ioutil.ReadFile(`morse.txt`)
	if e != nil {
		fmt.Printf("读取[morse.txt]异常: %v\n", e)
		return
	}
	split := strings.Split(string(bytes), "\n")
	for _, line := range split {
		morse := strings.Split(line, "|")
		for range morse {
			morseMap[strings.TrimSpace(morse[0])] = strings.TrimSpace(morse[1])
			textMap[strings.TrimSpace(morse[1])] = strings.TrimSpace(morse[0])
		}
	}
}

func main() {
	var message string
	flag.StringVar(&message, `c`, `hello world`, `请输入需要转换的英文字符串😄`)
	flag.Parse()
	message = strings.TrimSpace(message)
	msgRunes := []rune(message)
	morseLength := 0
	for _, v := range msgRunes {
		if strings.EqualFold(string(v), ".") || strings.EqualFold(string(v), "-") || strings.EqualFold(string(v), " ") {
			morseLength++
		}
	}
	isMorse := false
	if morseLength == len(message) {
		isMorse = true
	}
	if !isMorse {
		fmt.Printf("[%v] 的莫斯电码是 [%v]\n", message, msg2morse(message))
	} else {
		fmt.Printf("莫斯电码 [%v] 翻译结果是 [%v]\n", message, morse2msg(message))
	}
}

func msg2morse(s string) string {
	var mouseContent string
	split := strings.Split(s, " ")
	for index, value := range split {
		src := []rune(value)
		var mouseWorld string
		for i, c := range src {
			word := strings.ToUpper(string(c))
			if strings.EqualFold(word, ` `) {
				continue
			}
			word = morseMap[word]
			if i == 0 {
				mouseWorld = word
			} else {
				mouseWorld = mouseWorld + " " + word
			}
		}
		if index == 0 {
			mouseContent = mouseWorld
		} else {
			mouseContent = mouseContent + "  " + mouseWorld
		}
	}
	return mouseContent
}

func morse2msg(s string) string {
	split := strings.Split(s, " ")
	var content string
	for _, value := range split {
		if strings.EqualFold(value, "") {
			content = content + " "
			continue
		}
		text := textMap[value]
		content = content + text
	}
	return content
}
