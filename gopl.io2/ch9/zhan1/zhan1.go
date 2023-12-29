package zhan1

// todo 使用通道构造一个把任意多个 goroutine 串联在一起的流水线序。在内存耗尽之前你能创建的最大流水线级数是多少?一个值穿过整个流水线需要多久?
type A struct {
	Message string
}

func (in A) GetMessage() string {
	return in.Message
}
		