package main

//import "runtime"

type Vector []float64

//runtime.GOMAXPROCS(runtime.NumCPU())
// 分配给每个CPU的计算任务
func (v Vector) DoSome(i, n int, u Vector, c chan int) {
	for ; i < n; i++ {
		//		v[i] += u.Op(v[i])
	}
	c <- 1 // 发信号告诉任务管理者我已经计算完成了
}

const NCPU = 4 // 假设总共有16核
func (v Vector) DoAll(u Vector) {
	c := make(chan int, NCPU) // 用于接收每个CPU的任务完成信号
	for i := 0; i < NCPU; i++ {
		go v.DoSome(i*len(v)/NCPU, (i+1)*len(v)/NCPU, u, c)
	}
	// 等待所有CPU的任务完成
	for i := 0; i < NCPU; i++ {
		<-c // 获取到一个数据，表示一个CPU计算完成了
	}
	// 到这里表示所有计算已经结束
}

func maind() {
	//	DoAll()
}
