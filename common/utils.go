package common

import (
	"math/rand"
	"os"
	"os/signal"
	"reflect"
	"runtime"
	"syscall"
	"time"
)

//////////////////////////////////////////////////////////////////////
// check param invalid, if not assert
//////////////////////////////////////////////////////////////////////
func Assert(expr bool, message string) {
	if !expr {
		panic(message)
	}
}

func CheckParam(expr bool) {
	Assert(expr, "invalid param")
}

func WaitKill() {
	// Wait for terminating signal
	sc := make(chan os.Signal, 2)
	signal.Notify(sc, syscall.SIGTERM, syscall.SIGINT)
	<-sc
}

func GenerateRandomKey(length int) []byte {
	rand.Seed(time.Now().UTC().UnixNano())
	key := make([]byte, length)
	for i := 0; i < length; i++ {
		key[i] = byte(rand.Intn(256))
	}
	return key
}

// get function name
func GetFuncName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}
