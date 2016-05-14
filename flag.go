/*使用flag取得命令行标志类参数的值
最后一个参数是help信息：go run test.go --help，也可以用-help，-h, --h
*/
package main

import (
	"flag"
	"fmt"
)

func main() {
	ok := flag.Bool("ok", false, "is ok")
	id := flag.Int("id", 0, "id")
	port := flag.String("port", ":8080", "http listen port")
	var name string
	flag.StringVar(&name, "name", "123", "name")

	flag.Parse()

	fmt.Println("ok:", *ok)
	fmt.Println("id:", *id)
	fmt.Println("port:", *port)
	fmt.Println("name:", name)
}
