
package main //import "


import (
	"fmt"
	protobuf "github.com/golang/protobuf/proto"
	"gongxianghua/ping_pong/src/ping_pong/data"
	"os"
)

func main() {
	//初始化protobuf数据格式
	msg := &data.HelloWorld{
		Id:     protobuf.Int32(17),
		Name:   protobuf.String("gxh"),
		Opt:    protobuf.Int32(18),

	}

	filename := "./protobuf-test.txt"
	fmt.Printf("使用protobuf创建文件 %s\n",filename)
	fObj,_ := os.Create(filename)
	defer fObj.Close()
	buffer,_ := protobuf.Marshal(msg)
	fObj.Write(buffer)

}
