package media

import (
	"fmt"
	"reflect"
	"time"
)

type SCTask struct {
	id  int
	url string
}
func (x SCTask) IsStructureEmpty() bool {
	return reflect.DeepEqual(x, SCTask{})
}

type Queue struct {
	element []SCTask
}

//创建一个新队列
func NewQueue() *Queue {
	return &Queue{}
}

//判断队列是否为空
func (s *Queue) IsEmpty() bool {
	if len(s.element) == 0 {
		return true
	} else {
		return false
	}
}

//求队列的长度
func (s *Queue) GetQueueLength() int {
	return len(s.element)
}

//进队操作
func (s *Queue) Push(value SCTask) {
	s.element = append(s.element, value)
}

//出队操作
func (s *Queue) Pop() (bool, SCTask) {
	var sctask SCTask
	if s.IsEmpty() {
		return false, sctask
	} else {
		sctask = s.element[0]
		s.element = s.element[1:]
	}
	return true, sctask
}

var chTaskList chan SCTask

var finishOneTask = make(chan int)
var CLOSE = make(chan int)
var taskueue = NewQueue()

func StartScWork(taskNum int) {
	if taskNum < 0 {
		taskNum = 1
	}
	go dealScWork(taskNum)
}

func CloseScWork() {
	close(CLOSE)
}

func dealScWork(taskNum int) {
	chTaskList = make(chan SCTask,taskNum)
	go working(chTaskList)
	tick := time.Tick(time.Second * 10)
	for {
		var chTaskListtemp chan SCTask
		var taskinit SCTask=SCTask{}
		if !taskueue.IsEmpty() {
			_, taskinit= taskueue.Pop()
			chTaskListtemp=chTaskList
		}
		select {
		case finishindex := <-finishOneTask:
			fmt.Printf("finishindex index: %d   ,tick 还剩size:%d  \n", finishindex, taskueue.GetQueueLength())
			if !taskinit .IsStructureEmpty(){
				chTaskList <- taskinit
			}
		case <-tick:
			if !taskueue.IsEmpty() {
				fmt.Printf(" tick 还剩size:%d   \n", taskueue.GetQueueLength())
			}
		case chTaskListtemp <- taskinit:{
			fmt.Printf(" 插入ok 还剩size:%d   \n", taskueue.GetQueueLength())
		}
		case <-CLOSE:
			//关闭任务列队
			fmt.Println("close")
			return
		}
	}
	//wg.Wait()
}

func working(chTasks chan SCTask) {
	for task := range chTasks {
		fmt.Printf("处理线程: %d  url:%s ,id:%d  ,当前线程还剩size:%d \n", 1, task.url, task.id, taskueue.GetQueueLength())
		time.Sleep(time.Second * 3)
		finishOneTask <- 1

	}
}
var id=0;
func AddTask( url string) {
	taskueue.Push(SCTask{
		id:  id,
		url: url,
	})
	id++;
}
