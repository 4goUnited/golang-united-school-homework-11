package batch

import (
	"time"
	"sync"
)

type user struct {
	ID int64
}

func getOne(id int64) user {
	time.Sleep(time.Millisecond * 100)
	return user{ID: id}
}

//func getBatch(n int64, pool int64) (res []user) {
//	var id int64 = -1
//	var mx sync.Mutex
//	var wg sync.WaitGroup
//	for int64(len(res)) < n {
//		for i := 0; int64(i) < pool; i++ {
//			id += 1
//			wg.Add(1)
//			go func(wg *sync.WaitGroup, j int64) {
//	                        defer wg.Done()
//				u :=  getOne(j)
//				mx.Lock()
//				res = append(res, u)
//				mx.Unlock()
//			}(&wg, id)
//		}
//		wg.Wait()
//	}
//	return
//}

func getBatch(n int64, pool int64) (res []user) {
	var i int64
	var mx sync.Mutex
	var wg sync.WaitGroup
	ch := make(chan struct{}, pool)
	for ; i < n; i++ {
		wg.Add(1)
		ch <- struct{}{}
		go func(i int64) {
			u := getOne(i)
			<-ch
			mx.Lock()
			res = append(res, u)
			mx.Unlock()
			wg.Done()
		}(i)
	}
	wg.Wait()
	return
}
