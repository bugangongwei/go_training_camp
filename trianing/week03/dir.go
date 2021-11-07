package week03

import (
	"errors"
	"fmt"
	"strconv"
)

const (
	FileNum = 1000000
	// targetFileName = 977777
	errFileIdx = 977777
)

// listDir list dir and return slice
func listDir() []int {
	var result = make([]int, FileNum)
	for i := 0; i< FileNum;i++ {
		result[i] = i
	}
	return result
}

// streamDir read dir and send it to channel one by one
/*
a leak of goroutine
when caller return in targetFileName, there is no receiver left
and the goroutine will block to wait for receiver
*/
func streamDir(targetFileIdx int) chan string {
	var c = make(chan string)
	go func(targetFileIdx int) {
		defer close(c)

		for i := 0; i< FileNum;i++ {
			// close channel while err
			if i == errFileIdx {
				return
			}

			// normal send
			c <- strconv.Itoa(i)

			// send block
			if i > targetFileIdx {
				fmt.Println("targetFile 以后的 file： ", i)
			}
		}
	}(targetFileIdx)
	return c
}

// callbackDir scan dir and call callback method when find the target file
func callbackDir(targetFileIdx int, print func(fileIdx int)) error {
	for i := 0; i< FileNum;i++ {
		if i == errFileIdx {
			return errors.New("we encounter an error")
		}

		if i == targetFileIdx {
			print(i)
		}
	}
	return nil
}

// FindFileFromSlice iterator slice and print target file
func FindFileFromSlice(targetFileIdx int) {
	files := listDir()
	for _, file := range files {
		if file == targetFileIdx {
			fmt.Println(file)
		}
	}
}

// FindFileFromChan read dir from channel and return when find target one
func FindFileFromChan(targetFileIdx int) {
	defer fmt.Println("剩下的 channel 大小： ", len(streamDir(targetFileIdx)))

	for file := range streamDir(targetFileIdx) {
		if file == strconv.Itoa(targetFileIdx) {
			fmt.Println(file)
			return
		}
	}
}

// FindFileAndCallBack find file by callback
func FindFileAndCallBack(targetFileIdx int) error {
	return callbackDir(targetFileIdx, func(fileIdx int) {
		fmt.Println(fileIdx)
	})
}



