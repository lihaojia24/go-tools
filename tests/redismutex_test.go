package tests

import (
	"fmt"
	"testing"
	"time"

	"github.com/lihaojia24/go-xtools"
)

func Test_redisMutex(t *testing.T) {
	m1 := xtools.NewRedisMutex("m1", "lockname")
	m2 := xtools.NewRedisMutex("m2", "lockname")
	m1.Lock()
	m1.Unlock()
	// fmt.Printf("m1.TryLock(): %v\n", m1.TryLock())
	// time.Sleep(time.Second * 5)
	m2.Lock()
	// fmt.Printf("m2.TryLock(): %v\n", m2.TryLock())
	fmt.Printf("m1.Unlock(): %v\n", m1.Unlock())
	fmt.Printf("m2.Unlock(): %v\n", m2.Unlock())
	time.Sleep(time.Second)
}
