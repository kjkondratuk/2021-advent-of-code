package lib

type IntCollection []int

func (c IntCollection) Comprehend(bufferSize int, interpreter func([]int) []int) []int {
	uDepth := len(c)
	res := make([]int, 0)
	var i int
	for i = 0; i < uDepth-bufferSize+1; i++ {
		buffer := make([]int, 0)
		var offset int
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
