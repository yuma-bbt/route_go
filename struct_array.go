package main

import "fmt"

// テスト用構造体
type example struct {
	Name string
}

// example構造体を格納する配列の型
type examples []*example

// doExample exampleのNameフィールドに引数のstringを指定して返す
func doExample(n string) (r *example) {
	r = new(example)
	r.Name = n
	return r
}

func main() {

	// example構造体配列を宣言
	var exs examples

	// 配列に構造体を格納
	exs = append(exs, doExample("A"))
	exs = append(exs, doExample("B"))

	// 格納した構造体のNameフィールドを出力
	for _, i := range exs {
		fmt.Println(i.Name)
	}
}
