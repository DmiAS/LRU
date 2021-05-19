package capacity

// тип для задания размера в байтах
type BytesCap uint64

func NewBytesCap(size uint64) BytesCap {
	return BytesCap(size)
}

func (b BytesCap) GetSize() uint16 {
	size := b / (keySize + valSize) // 4 байта ключ + 128 байт строка
	if size > maxCacheSize {
		size = maxCacheSize
	}
	return uint16(size)
}
