package utils

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"testing"
)
//const  chunkSize = 1<<(10*2) //1kB
const  chunkSize  = 1<<(10*2) //1kB
const filepath="/Users/mac/Desktop/testfile/vue.mp4"
const New_filepath="/Users/mac/Desktop/testfile/vue2.mp4"
const filechunkpath="/Users/mac/Desktop/testfile/chunk/"


func Test_File(t *testing.T) {
	println(HashCode("842ca630956535c83eca7400e32429eca147d6bf"))
	println(HashCode("842ca630956535c83eca7400e32429eca147d6bf"))
	println(HashCode("842ca630956535c83eca7400e32429eca147d6bf"))
	println(HashCode("842ca630956535c83eca7400e32429eca147d6bf"))
	/*fileinfo,error:=os.Stat(filepath)
	if error!=nil{
		panic(error)
	}
	chunkFile(fileinfo)
	fileMerge(filechunkpath,New_filepath)*/
	//println(HashCode("be166f50301e6268415f714b4e66bd370a89abca"))//189723058
	//println(HashCode("be166f50301e6268415f714b4e66bd370a89abca"))//189723058
	//println(HashCode("be166f50301e6268415f714b4e66bd370a89abca"))
	//println(HashCode("be166f50301e6268415f714b4e66bd370a89abca"))
	//println(HashCode("be166f50301e6268415f714b4e66bd370a89abca"))
	//println(HashCode("be166f50301e6268415f714b4e66bd370a89abca"))
	//println(HashCode("be166f50301e6268415f714b4e66bd370a89abca"))
	//println(HashCode("e165052bd9f4e2bd6a1575dba2c3e806ee55d131"))
	//println(HashCode("e165052bd9f4e2bd6a1575dba2c3e806ee55d131"))
	//println(HashCode("e165052bd9f4e2bd6a1575dba2c3e806ee55d131"))
	//println(HashCode("e165052bd9f4e2bd6a1575dba2c3e806ee55d131"))
	//println(HashCode("e165052bd9f4e2bd6a1575dba2c3e806ee55d131"))
	//println(HashCode("e165052bd9f4e2bd6a1575dba2c3e806ee55d131"))
	//code := HashCode("e165052bd9f4e2bd6a1575dba2c3e806ee55d131")
	//println(CreateDirbySha1("rootpath","tln(HashCode(\"be166f50301e6268415f714b4e66bd370a89abc",1))
	//println(CreateDirbySha1("rootpath","123",1))
	//println(CreateDirbySha1("rootpath","tln(HashCode(\"be166f50301e6268415f714b4e66bd370a89abc",1))
	//println(CreateDirbySha1("rootpath","tln(HashCode(\"be166f50301e6268415f714b4e66bd370a89abc",1))
	//println(CreateDirbySha1("rootpath","tln(HashCode(\"be166f50301e6268415f714b4e66bd370a89abc",1))
	//println(CreateDirbySha1("rootpath","",1))

}


func chunkFile(fileinfo os.FileInfo) {
	flow:=float64(fileinfo.Size()) / chunkSize
	chunkNum := math.Ceil(flow)
	file, error := os.OpenFile(filepath, os.O_RDONLY, os.ModePerm)
	if error!=nil{
		return
	}
	b := make([]byte, chunkSize)
	var i int64=1
	for  ; i <= int64(chunkNum); i++ {
		file.Seek((i-1)*chunkSize, 0)
		if i==int64(chunkNum)  {
			//todo 处理最后一个 不足缓冲去的情况
			lost:=fileinfo.Size()- chunkSize*(i-1)
			if len(b)>int(lost){
				b=make([]byte,lost)
			}
		}
		file.Read(b)
		f, error := os.OpenFile(filechunkpath+strconv.Itoa(int(i))+".chunk", os.O_CREATE|os.O_WRONLY, os.ModePerm)
		if error != nil {
			println(error)
		}
		f.Write(b)
		f.Close()
	}
	file.Close()
}



func fileMerge(chunkpath string,newfile string )  {
	dirfiles, _ := ioutil.ReadDir(chunkpath)
	err := os.Remove(newfile)
	if err==nil{

	}
	fill, error := os.OpenFile(newfile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	if error != nil {
		println(error)
		return
	}
	for index,filetemp:= range dirfiles{
		//ioutil.ReadFile(filetemp.)
		fmt.Printf("index %d filename %s ", index,filetemp.Name())
		fmt.Printf("index %d filepath%s \n" ,index,filetemp.Name())
		file, error := os.OpenFile(chunkpath+strconv.Itoa(index)+".chunk", os.O_RDONLY, os.ModePerm)
		if error!=nil{
			println(error)
		}
		all, _ := ioutil.ReadAll(file)
		fill.Write(all)
		file.Close()
	}
}






