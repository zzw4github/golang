package main


type PipeDate struct {
	value int
	handler func(int) int
	next chan int
}

func handle(queue chan *PipeDate) {
	for data := range queue {
		data.next <- data.handler(data.value)
	}
}

