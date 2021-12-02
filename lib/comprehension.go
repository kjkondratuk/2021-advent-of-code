package lib

type UInt32Comprehendable []uint32

func (c UInt32Comprehendable) Comprehend(bufferSize uint32, interpreter func([]uint32) []uint32) []uint32 {
	uDepth := uint32(len(c))
	res := make([]uint32, 0)
	var i uint32
	for i = 0; i < uDepth-bufferSize+1; i++ {
		buffer := make([]uint32, 0)
		var offset uint32
		for offset = 0; offset < bufferSize; offset++ {
			//log.Printf("comprehension index: %d", i+offset)
			if i+offset > uDepth-1 {
				break
			}
			buffer = append(buffer, c[i+offset])
		}
		//log.Printf("evaluating: %+v", buffer)
		res = append(res, interpreter(buffer)...)
	}
	return res
}
