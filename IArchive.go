package goproto

type IArchive interface {
	DoSomething(interface{})
	DoSerialize(IExtensible)
}

type IExtensible interface {
	Serialize(ar IArchive)
}
