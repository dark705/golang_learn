package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

type Bank struct {
	bill int64
}

func (b *Bank) AddMoney(sum int) {
	atomic.AddInt64(&b.bill, int64(sum))
}

func (b *Bank) ShowBill() int64 {
	return atomic.LoadInt64(&b.bill)
}

type User struct {
	name string
	bill int
}

func (u User) SendMoney(b *Bank, sum int) {
	time.Sleep(time.Millisecond * 1) //transaction wait
	b.AddMoney(sum)
	u.bill -= sum
}

func (b *Bank) MonthCommissionConsistently(users []User) {
	for _, u := range users {
		u.SendMoney(b, 1)
	}
}

func (b *Bank) MonthCommissionConcurrency(users []User) {
	sync := make(chan struct{})
	for _, u := range users {
		go func() {
			<-sync

			u.SendMoney(b, 1)
		}()
	}
	close(sync)
}

func main() {
	b := Bank{}
	users := make([]User, 0, 8000)
	for i := 0; i < 8000; i++ {
		users = append(users, User{"Some", 100})
	}
	//	b.MonthCommissionConsistently(users) //long
	fmt.Println(b.ShowBill()) //8000

	b.MonthCommissionConcurrency(users) //fast
	time.Sleep(time.Second)
	fmt.Println(b.ShowBill()) //16000

}
