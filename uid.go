package uid

type uid interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | 
		~string
}

type Identifier[T uid] interface {
	Value() any
}
