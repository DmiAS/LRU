package capacity

// тип для задания размера в количестве элементов
type ElemCap uint16

func NewElemCap(size uint16) ElemCap {
	return ElemCap(size)
}

func (e ElemCap) GetSize() uint16 {
	if e > maxCacheSize {
		return maxCacheSize
	}
	return uint16(e)
}
