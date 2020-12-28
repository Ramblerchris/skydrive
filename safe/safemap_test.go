package safe

import (
	"fmt"
	"testing"
	"time"
)

func Test_SafeMap(t *testing.T) {
	//var saveMap =SafeMap{
	//	sy:sync.RWMutex{},
	//	Map:make(map[string]interface{})}
	 saveMap:=new (SafeMap)
	 saveMap.Map= make(map[string]int)
	for i := 1; i < 10; i++ {
		go saveMap.WriteMap(fmt.Sprintf("%d test", i), i)
	}
	for i := 1; i < 10; i++ {
		go func() {
			println(saveMap.ReadMap(fmt.Sprintf("%d test", i)))
		}()
	}
	time.Sleep(time.Second)
}
