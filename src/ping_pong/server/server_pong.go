package main
import (
	"fmt"
	"github.com/Golang_Puzzlers/src/puzzlers/ping_pong/data"
	protobuf "github.com/golang/protobuf/proto"
	"io"
	"os"
)


func checkError(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(-1)
	}
}

func main() {
	filename := "protobuf-test.txt"
	file,fileErr := os.Open(filename)
	checkError(fileErr)

	defer file.Close()
	fs,fsErr := file.Stat()
	checkError(fsErr)
	buffer := make([]byte,fs.Size())
	//把file文件内容读取到buffer
	_,readErr := io.ReadFull(file,buffer)
	checkError(readErr)

	//初始化pb结构体对象并将buffer中的文件内容读取到pb结构体中
	msg := &data.HelloWorld{}
	pbErr := protobuf.Unmarshal(buffer, msg)
	checkError(pbErr)
	fmt.Printf("读取文件:%s \r\nname:%s\nid:%d\nopt:%d\n",filename,msg.GetName(),msg.GetId(),msg.GetOpt())
}