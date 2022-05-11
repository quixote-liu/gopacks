package main

import (
	"fmt"
	"strconv"
)

func main() {
	// text := "hello 你好"
	// res, err := zhToUnicode([]byte(text))
	// dealError(err)
	// fmt.Println("res:", res)

	// res2, err := unicodeToZh([]byte(res))
	// dealError(err)
	// fmt.Println("res2:", res2)

	// text1 := `\nhello, \u4f60\u597d\u6211\u53eb\u6e56\u5317\u5927\u5b66`
	// res1, err := unicodeToZh(text1)
	// dealError(err)
	// fmt.Println(res1)
	// text2 := "Resource not found: [GET http://10.0.220.11:9292/v2/images/823b4ea5-cb8b-4a38-8b3e-fa9581fd87a], error message: {\"message\": \"\u627e\u4e0d\u5230\u4efb\u4f55\u5177\u6709\u6807\u8bc6 823b4ea5-cb8b-4a38-8b3e-fa9581fd87a \u7684\u6620\u50cf<br /><br />\n\n\n\", \"code\": \"404 Not Found\", \"title\": \"Not Found\"}"
	// fmt.Printf("%v\n", text2)

	// err2 := errors.New(text2)
	// err2Str := fmt.Sprint(err2.Error())
	// fmt.Println("err2Str:", err2Str)

	// byte2 := []byte(text2)
	// byte2Str := fmt.Sprintf("%q", byte2)
	// byte2Str := fmt.Sprint(string(byte2))
	// fmt.Println(byte2Str)

	// fmt.Println(text2New)
	// res2, err := unicodeToZh(text2)
	// dealError(err)
	// fmt.Println(res2)

	UnicodeToZhDemo()
}

func dealError(err error) {
	if err != nil {
		panic(err)
	}
}

func zhToUnicode(raw []byte) (string, error) {
	text := strconv.QuoteToASCII(string(raw))
	return text[1 : len(text)-1], nil
}

// func unicodeToZh(raw []byte) (string, error) {
// 	// first
// 	s1 := strconv.Quote(string(raw))
// 	fmt.Println("s1:", s1)

// 	str, err := strconv.Unquote(s1)
// 	if err != nil {
// 		return "", err
// 	}
// 	return str, nil
// }

func UnicodeToZhDemo() {
	text1 := "\nhello, \u4f60\u597d\uff0c\u6211\u662f\u5218\u843d\u8f69"
	s1, err := unicodeToZh(text1)
	dealError(err)
	fmt.Println("s1:", s1)
}

func unicodeToZh(raw string) (string, error) {
	// first
	s1 := strconv.Quote(raw)
	fmt.Println("s1:", s1)

	str, err := strconv.Unquote(s1)
	if err != nil {
		return "", err
	}
	return str, nil
}
