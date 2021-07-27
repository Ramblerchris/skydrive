package media

import (
	"fmt"
	"github.com/skydrive/config"
	"github.com/skydrive/utils"
	"reflect"
	"strconv"
	"time"
)

type SCTask struct {
	Sha1         string
	Sctype       int
	Locationpath string
	FileName     string
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
	chTaskList = make(chan SCTask, taskNum)
	go working(chTaskList)
	tick := time.Tick(time.Second * 10)
	for {
		var chTaskListtemp chan SCTask
		var taskinit = SCTask{}
		if !taskueue.IsEmpty() {
			_, taskinit = taskueue.Pop()
			chTaskListtemp = chTaskList
		}
		select {
		case finishindex := <-finishOneTask:
			fmt.Printf("finishindex index: %d   ,tick 还剩size:%d  \n", finishindex, taskueue.GetQueueLength())
			if !taskinit.IsStructureEmpty() {
				chTaskList <- taskinit
			}
		case <-tick:
			if !taskueue.IsEmpty() {
				fmt.Printf(" tick 还剩size:%d   \n", taskueue.GetQueueLength())
			}
		case chTaskListtemp <- taskinit:
			{
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
		scImageAndVideo(task)

	}
}

func scImageAndVideo(task SCTask) {

	fmt.Printf("处理线程: %s   ,当前线程还剩size:%d \n", task.Sha1, taskueue.GetQueueLength())
	if task.Sctype == 1 {
		//video
		err, target := utils.CreateThumbDir(config.ThumbnailRoot, task.Sha1, strconv.FormatInt(config.Thumbnail_index, 10), task.FileName+".jpg")
		if err == nil {
			exists, _, info := utils.PathExistsInfo(target)
			if !exists || info.Size() < 1000 {
				VideoThumbnail(task.Locationpath, target)
			}
		}
	} else if task.Sctype == 0 {
		//image
		exists, _, _ := utils.PathExistsInfo(task.Locationpath)
		//if exists && info.Size() > 102400 {
		if exists {
			//图片压缩
			err, target := utils.CreateThumbDir(config.ThumbnailRoot, task.Sha1, strconv.FormatInt(config.Thumbnail_index, 10), task.FileName)
			if err == nil {
				//_, error := os.Open(data.FileLocation)
				exists, _ := utils.PathExists(target)
				if !exists {
					ScaleImageByWidthAndQuity(task.Locationpath, config.Thumbnail_width, config.Thumbnail_widthf, config.Thumbnail_Quality, target)
				}
			}
		}
	}
	finishOneTask <- 1
}
func AddSCTask(url SCTask) {
	taskueue.Push(url)
}
