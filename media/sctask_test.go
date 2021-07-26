package media

import (
	"testing"
)

func Test_sctask(t *testing.T) {
	StartScWork(1)
	AddTask(1,"url1")
	AddTask(2,"url2")
	AddTask(3,"url3")
	AddTask(4,"url4")
	AddTask(5,"url5")
}

