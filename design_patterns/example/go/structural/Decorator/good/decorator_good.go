package good

// IStream 业务操作
type IStream interface {
	Read(number int) []byte
	Seek(position int)
	Write(data []byte)
}

// FileStream 主体类：文件流
type FileStream struct {
}

func NewFileStream() *FileStream {
	return &FileStream{}
}

func (th *FileStream) Read(number int) []byte {
	// 读文件流
	return nil
}

func (th *FileStream) Seek(position int) {
	// 定位文件流
}

func (th *FileStream) Write(data []byte) {
	// 写文件流
}

// NetworkStream 主体类：网络流
type NetworkStream struct {
}

func NewNetworkStream() *NetworkStream {
	return &NetworkStream{}
}

func (th *NetworkStream) Read(number int) []byte {
	// 读网络流
	return nil
}

func (th *NetworkStream) Seek(position int) {
	// 定位网络流
}

func (th *NetworkStream) Write(data []byte) {
	// 写网络流
}

// MemoryStream 主体类：内存流
type MemoryStream struct {
}

func NewMemoryStream() *MemoryStream {
	return &MemoryStream{}
}

func (th *MemoryStream) Read(number int) []byte {
	// 读内存流
	return nil
}

func (th *MemoryStream) Seek(position int) {
	// 定位内存流
}

func (th *MemoryStream) Write(data []byte) {
	// 写内存流
}

// 扩展操作：加密------------------------------------------

// CryptoStream 扩展操作(实现IStream，为了规范流接口定义)
type CryptoStream struct {
	s IStream
}

func NewCryptoStream(stream IStream) *CryptoStream {
	return &CryptoStream{
		s: stream,
	}
}
func (th *CryptoStream) Read(number int) []byte {
	// 额外的加密操作...
	th.s.Read(number) // 读文件流
	return nil
}

func (th *CryptoStream) Seek(position int) {
	// 额外的加密操作...
	th.s.Seek(position) // 定位文件流
	// 额外的加密操作...

}

func (th *CryptoStream) Write(data []byte) {
	// 额外的加密操作...
	th.s.Write(data) // 写文件流
	// 额外的加密操作...

}

// 扩展操作：缓存------------------------------------------

type BufferStream struct {
	s IStream
}

func NewBufferStream(stream IStream) *BufferStream {
	return &BufferStream{
		s: stream,
	}
}
func (th *BufferStream) Read(number int) []byte {
	// 额外的加密操作...
	th.s.Read(number) // 读文件流
	return nil
}

func (th *BufferStream) Seek(position int) {
	// 额外的缓存操作...
	th.s.Seek(position) // 定位文件流
	// 额外的缓存操作...

}

func (th *BufferStream) Write(data []byte) {
	// 额外的缓存操作...
	th.s.Write(data) // 写文件流
	// 额外的缓存操作...

}

// 说明：由于GO中没有继承，对于公共成员IStream，不再提取作为基类。
// 多个子类具有公共字段，可以将公共字段提取到父类中：
// 如果是C++，对于CryptoStream、BufferStream可以提取一个DecoratorStream，
// 继承IStream且有一个成员IStream，CryptoStream、BufferStream继承DecoratorStream