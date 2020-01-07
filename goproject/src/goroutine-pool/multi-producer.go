package main

import (
	"fmt"
	"github.com/panjf2000/ants"
	//"runtime/debug"
	"sync"
	"time"
)

type ChannelData struct {
	Data []int
}

func (cd *ChannelData) Extend(extendCd *ChannelData) {

	cd.Data = append(cd.Data, extendCd.Data...)

}

func (cd *ChannelData) Deal() (err error) {

	fmt.Printf(cd.ToString())

	cd.Data = nil
	return

}
func (cd *ChannelData) ToString() string {
	return fmt.Sprintf("%v", cd.Data)

}

var MultiProducerPool *ants.PoolWithFunc

type MultiProducerParams struct {
	Wg    *sync.WaitGroup //must be pointer or can not be transfer into here
	Mutex sync.Mutex

	Begin int
	End   int

	C chan *ChannelData
	//SouceData *[]map[string]string
}

var testChannel = make(chan *ChannelData, 65535)

func (mpp *MultiProducerParams) Work() {
	tmpBegin := mpp.Begin
	for tmpBegin < mpp.End {

		var cd ChannelData

		cd.Data = []int{tmpBegin}
		mpp.C <- &cd

		tmpBegin = tmpBegin + 1
	}

}
func (mpp *MultiProducerParams) Done() {
	mpp.Wg.Done()
}

func init() {
	var err error
	MultiProducerPool, err = ants.NewPoolWithFunc(1000, func(i interface{}) {

		defer func() {
			if r := recover(); r != nil {

				panic(r)
			}
		}()

		fqp, _ := i.(MultiProducerParams)
		defer fqp.Done()

		fqp.Work()

	})

	if err != nil {
		panic("ants FlagQueryPool init fail err: " + err.Error())
	}
}
func main() {
	fmt.Println("good")
	var wg sync.WaitGroup
	//var mutex sync.Mutex
	delta := 2
	for i := 0; i <= 100; i = i + 1 {
		begin := i * delta
		end := i*delta + delta

		wg.Add(1)

		var mpp MultiProducerParams

		mpp.Wg = &wg

		mpp.Begin = begin
		mpp.End = end

		mpp.C = testChannel
		MultiProducerPool.Invoke(mpp)

	}

	//time.Sleep(time.Second * 1)
	wg.Wait()

	fmt.Println("close channel")

	close(testChannel)

	wg.Add(1)
	go func() {
		var cd *ChannelData

		cd = &ChannelData{}
		nowNum := 0

		bulkNum := 10

		for {
			time.Sleep(1)

			data, ok := <-testChannel
			if !ok {
				fmt.Println("not ok", data == nil, " len:", len(testChannel))
				break
			}

			if nowNum < bulkNum {

				cd.Extend(data)
				nowNum = nowNum + 1

			} else {

				cd.Deal()

				fmt.Println(" len:", len(testChannel))
				nowNum = 0

			}

		}

		cd.Deal()

		wg.Done()

	}()

	wg.Wait()
}
