/*使用MD5方式加密字符串*/
package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
)

func main() {
	h := md5.New()
	//io.WriteString(h, "The fog is getting thicker!")
	//io.WriteString(h, "And Leon's getting laaarger!")
	io.WriteString(h, "cexoyq1020")
	fmt.Println("%x", h.Sum(nil), "%s", h.Sum(nil))
	io.WriteString(h, "cexoyq1020")
	fmt.Println("%x", h.Sum(nil), "%s", h.Sum(nil))
	fmt.Printf("cexoyq1020:%s\n", hex.EncodeToString(h.Sum(nil))) // 输出加密结果
	data := []byte("These pretzels are making me thirsty.")
	fmt.Println("%x", md5.Sum(data))
	// md5 加密的第一种方法，也可以用
	srcData := []byte("szq")
	cipherText1 := md5.Sum(srcData)
	fmt.Printf("md5 encrypto is \"szq\": %x \n", cipherText1)
	srcData = []byte("cexoyq1020")
	cipherText1 = md5.Sum(srcData)
	fmt.Printf("md5 encrypto is \"cexoyq1020\": %x \n", cipherText1)
	/*用的下面这种没问题**********/
	h = md5.New()
	h.Write([]byte("cexoyq1020"))                                 // 需要加密的字符串为 cexoyq1020
	fmt.Printf("cexoyq1020:%s\n", hex.EncodeToString(h.Sum(nil))) // 输出加密结果
	h.Write([]byte("szq"))                                        // 需要加密的字符串为
	fmt.Printf("szq:%s\n", hex.EncodeToString(h.Sum(nil)))        // 输出加密结果
}
