package goprotobuffer

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"reflect"
)

type DeserializeArchive struct {
	mIndex  int    //记录读取位置
	mBuffer []byte //接收到的数据
}

func (self *DeserializeArchive) SetInit(buffer []byte, offset int) {
	self.mBuffer = buffer
	self.mIndex = offset
}

func (self *DeserializeArchive) DoSerialize(t IExtensible) {
	t.Serialize(self)
}

func (self *DeserializeArchive) DoSomething(t interface{}) {
	valueTy := reflect.TypeOf(t)
	typeName := valueTy.String()
	switch typeName {
	case "*int8":
		Bytes := make([]byte, 1)
		copy(Bytes, self.mBuffer[self.mIndex:self.mIndex+1])
		bytebuffer := bytes.NewBuffer(Bytes)
		binary.Read(bytebuffer, binary.LittleEndian, (t.(*int8)))
		self.mIndex++
	case "*uint8":
		Bytes := make([]byte, 1)
		copy(Bytes, self.mBuffer[self.mIndex:self.mIndex+1])
		bytebuffer := bytes.NewBuffer(Bytes)
		binary.Read(bytebuffer, binary.LittleEndian, (t.(*uint8)))
		self.mIndex++
	case "*int16":
		Bytes := make([]byte, 2)
		copy(Bytes, self.mBuffer[self.mIndex:self.mIndex+2])
		bytebuffer := bytes.NewBuffer(Bytes)
		binary.Read(bytebuffer, binary.LittleEndian, (t.(*int16)))
		self.mIndex += 2
	case "*uint16":
		Bytes := make([]byte, 2)
		copy(Bytes, self.mBuffer[self.mIndex:self.mIndex+2])
		bytebuffer := bytes.NewBuffer(Bytes)
		binary.Read(bytebuffer, binary.LittleEndian, (t.(*uint16)))
		self.mIndex += 2
	case "*int32":
		Bytes := make([]byte, 4)
		copy(Bytes, self.mBuffer[self.mIndex:self.mIndex+4])
		bytebuffer := bytes.NewBuffer(Bytes)
		binary.Read(bytebuffer, binary.LittleEndian, (t.(*int32)))
		self.mIndex += 4
	case "*uint32":
		Bytes := make([]byte, 4)
		copy(Bytes, self.mBuffer[self.mIndex:self.mIndex+4])
		bytebuffer := bytes.NewBuffer(Bytes)
		binary.Read(bytebuffer, binary.LittleEndian, (t.(*uint32)))
		self.mIndex += 4
	case "*int64":
		Bytes := make([]byte, 8)
		copy(Bytes, self.mBuffer[self.mIndex:self.mIndex+4])
		bytebuffer := bytes.NewBuffer(Bytes)
		binary.Read(bytebuffer, binary.LittleEndian, (t.(*int64)))
		self.mIndex += 8
	case "*uint64":
		Bytes := make([]byte, 8)
		copy(Bytes, self.mBuffer[self.mIndex:self.mIndex+4])
		bytebuffer := bytes.NewBuffer(Bytes)
		binary.Read(bytebuffer, binary.LittleEndian, (t.(*uint64)))
		self.mIndex += 8
	case "*float32":
		Bytes := make([]byte, 4)
		copy(Bytes, self.mBuffer[self.mIndex:self.mIndex+4])
		bytebuffer := bytes.NewBuffer(Bytes)
		binary.Read(bytebuffer, binary.LittleEndian, (t.(*float32)))
		self.mIndex += 4
	case "*float64":
		Bytes := make([]byte, 8)
		copy(Bytes, self.mBuffer[self.mIndex:self.mIndex+8])
		bytebuffer := bytes.NewBuffer(Bytes)
		binary.Read(bytebuffer, binary.LittleEndian, (t.(*float64)))
		self.mIndex += 8
	case "*bool":
		Bytes := make([]byte, 1)
		copy(Bytes, self.mBuffer[self.mIndex:self.mIndex+1])
		bytebuffer := bytes.NewBuffer(Bytes)
		binary.Read(bytebuffer, binary.LittleEndian, (t.(*bool)))
		self.mIndex += 1
	case "*string":
		var len uint16
		self.DoSomething(&len)
		fmt.Println(len)
		Bytes := make([]byte, len)
		copy(Bytes, self.mBuffer[self.mIndex:self.mIndex+int(len)])
		*(t.(*string)) = string(Bytes)
		self.mIndex += int(len)
	}
}
