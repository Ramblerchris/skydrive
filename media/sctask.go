package media

import (
	"fmt"
	"github.com/skydrive/config"
	"github.com/skydrive/utils"
	"image"
	"os"
	"reflect"
	"strconv"
	"time"
)

type SCTask struct {
	Sha1         string
	Sctype       int
	Minetype       string
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
	tick := time.Tick(time.Second * 1)
	for {
		var chTaskListtemp chan SCTask
		var taskinit = SCTask{}
		if !taskueue.IsEmpty() {
			_, taskinit = taskueue.Pop()
			chTaskListtemp = chTaskList
		}
		select {
		case <-finishOneTask:
			if !taskinit.IsStructureEmpty() {
				chTaskList <- taskinit
			}
			fmt.Printf("finishOneTask current size:%d  \n", taskueue.GetQueueLength())
		case <-tick:
			if !taskueue.IsEmpty() {
				fmt.Printf("ticktimer task current size:%d   \n", taskueue.GetQueueLength())
			}
			if !taskinit.IsStructureEmpty() {
				taskueue.Push(taskinit)
			}
		case chTaskListtemp <- taskinit:
			{
				fmt.Printf("send task success current size:%d   \n", taskueue.GetQueueLength())
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

var count=0

func scImageAndVideo(task SCTask) {
	count++
	fmt.Printf("Thumbnail count:%d sha: %s  \n", count ,task.Sha1)
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
					//todo 可以通过minitype 获取类型
					efile, err := os.Open(task.Locationpath)
					if err != nil {
						fmt.Printf("could not open file for : %s", task.Locationpath)
						return
					}
					_, format, err := image.Decode(efile)
					if format == "gif" {
						ImageThumbnailGif(config.Thumbnail_fuzz_gif, task.Locationpath, target)
					} else {
						ImageThumbnailJPG(config.Thumbnail_Quality, config.Thumbnail_widthf, task.Locationpath, target)
					}
				}
			}
		}else{
			fmt.Printf("file not exist : %s", task.Locationpath)
		}
	}
	finishOneTask <- 1
}
func AddSCTask(url SCTask) {
	taskueue.Push(url)
}
