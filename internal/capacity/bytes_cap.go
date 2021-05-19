package capacity

type BytesCap uint64

func NewBytesCap(size uint64) BytesCap {
	return BytesCap(size)
}

func (b BytesCap) GetSize() uint16 {
	size := b / (keySize + valSize)
	if size > maxCacheSize {
		size = maxCacheSize
	}
	return uint16(size)
}
