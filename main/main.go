package main

import (
	"github.com/bugangongwei/go_training_camp/trianing/week03"
)

func main() {
	// // record not found
	// fmt.Println(week02.GetAuthor(5))
	//
	// // record exist
	// fmt.Println(week02.GetAuthor(1))

	// iter slice
	// week03.FindFileFromSlice(1000)

	// go routine leak: chan send block because no receiver
	// week03.FindFileFromChan(9000)
	// err: do not know if it is error or no more data
	// week03.FindFileFromChan(977778)

	// call back: return with no block
	// week03.FindFileAndCallBack(9000)
	// call back: err
	// fmt.Println(week03.FindFileAndCallBack(977778))

	week03.LifeCycleManage()

	// quit := make(chan os.Signal, 1)
	// signal.Reset(os.Interrupt, syscall.SIGTERM, syscall.SIGKILL)
	// signal.Notify(quit, os.Interrupt, syscall.SIGTERM, syscall.SIGKILL)
	// <-quit
}
