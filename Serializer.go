package goprotobuffer

type Serializer struct {
	mSerializeArchive   SerializeArchive
	mDeserializeArchive DeserializeArchive
}

func (self *Serializer) Serialize(packet IExtensible, buffer []byte, offset int, length *int) {
	self.mSerializeArchive.SetInit(buffer, offset)
	packet.Serialize(&self.mSerializeArchive)
	*length = self.mSerializeArchive.Length()
}

func (self *Serializer) Deserialize(buffer []byte, offset int, packet IExtensible) {
	self.mDeserializeArchive.SetInit(buffer, offset)
	packet.Serialize(&self.mDeserializeArchive)
}
