//测试继承
package main

import (
	"flag"
	"log"
	"net/url"
	"github.com/gorilla/websocket"
)

var addr = flag.String("addr", "192.168.1.112:80", "http service address")
type Conn struct {
	conn *websocket.Conn
}

func (this *Conn) Close(){
	this.conn.Close()
}
func (this *Conn)Init() error{
	var err error
	u := url.URL{Scheme: "ws", Host: *addr, Path: "/ws"}
	log.Printf("connecting to %s", u.String())
	this.conn, _, err = websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
		return err
	}
	return nil
}

func (this *Conn)SendServerStat(msg string) error {
	var err error
	err = this.conn.WriteJSON(msg)
	if err != nil {
		log.Println("write:", err)
		return err
	}
	return nil
}

func (this *Conn)ReadServerStat() error {
	var err error
	_, message, err := this.conn.ReadMessage()
	if err != nil {
		return err
	}
	log.Printf("read:%s",message)
	return nil
}

func main() {
	var (
		//c *websocket.Conn
		conn Conn
	)
	/*u := url.URL{Scheme: "ws", Host: *addr, Path: "/ws"}
	log.Printf("connecting to %s", u.String())
	c, _, err = websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()*/
	conn.Init();
	defer conn.Close();
	msg := "ehehe"
	log.Printf("write:",msg)
	//conn=c;
	conn.SendServerStat(msg)
	conn.ReadServerStat()
	return
}