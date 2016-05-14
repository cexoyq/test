/*结构转字节，字节转结构*/
package main

import (
	"bytes"
	"encoding/binary"
	"log"
)

type A struct { // should be exported member when read back from buffer
	One int32
	Two int32
}
type B struct {
	One int16
	Two int16
	str [16]byte
}

func main() {
	var a A
	a.One = int32(1)
	a.Two = int32(2)
	log.Println("a's size is: ", binary.Size(a))
	buf := new(bytes.Buffer)

	binary.Write(buf, binary.LittleEndian, a) //把a写到缓冲区
	log.Println("after write ，buf is:", buf.Bytes())

	var aa A
	binary.Read(buf, binary.LittleEndian, &aa) //读缓冲区数据到结构体
	log.Println("after aa is ", aa)

	var b B
	b.One = 1
	b.Two = 2
	copy(b.str[:], "abcdefghijklmns!") //字符串string 转 []byte
	log.Println("b's size is: ", binary.Size(b))
	//不能使用未知长度的数据类型，如int,string,切片等。必须指明数据类型长度
	binary.Write(buf, binary.LittleEndian, b) //把结构体b写到缓冲区
	log.Println("buf is :", buf.Bytes())

	ret := buf.Bytes()
	log.Println("buf string is :", string(ret[4:binary.Size(b)])) //字节转str
}
