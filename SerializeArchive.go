package goprotobuffer

import (
	"bytes"
	"encoding/binary"
	"reflect"
)

type SerializeArchive struct {
	mIndex  int    //记录读取位置
	mBuffer []byte //接收到的数据
}

func (self *SerializeArchive) Length() int {
	return self.mIndex
}

func (self *SerializeArchive) SetInit(buffer []byte, offset int) {
	self.mBuffer = buffer
	self.mIndex = offset
}

func (self *SerializeArchive) DoSerialize(t IExtensible) {
	t.Serialize(self)
}

func (self *SerializeArchive) DoSomething(t interface{}) {

	var getBytes []byte

	valueTy := reflect.TypeOf(t)
	if valueTy.String() == "*string" {
		getBytes = []byte(*t.(*string))
		var len uint16 = uint16(len(getBytes))
		self.DoSomething(len)
	} else {
		bytebuffer := bytes.NewBuffer([]byte{})          //字节集合
		binary.Write(bytebuffer, binary.LittleEndian, t) //按照二进制写入字节
		getBytes = bytebuffer.Bytes()
	}

	for _, value := range getBytes {
		self.mBuffer[self.mIndex] = value
		self.mIndex++
	}
}
