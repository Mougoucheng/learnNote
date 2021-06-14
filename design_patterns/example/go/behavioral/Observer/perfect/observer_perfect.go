package perfect

import (
	"container/list"
	"fmt"
	"log"
)

type IProgress interface {
	DoProgress(value float64)
}

type FileSplitter struct {
	txtFilePath   string
	txtFileNumber int

	//mProgress IProgress // 抽象通知机制
	mProgressList *list.List // 抽象通知机制，支持多个观察者
}

// NewFileSplitter 更进一步的FileSplitter可以接口化，
// 将AddIProgress、RemoveIProgress、onProgress提出成为接口
func NewFileSplitter(path string, num int) *FileSplitter {
	return &FileSplitter{
		txtFilePath:   path,
		txtFileNumber: num,
		mProgressList: list.New(),
	}
}
func (th *FileSplitter) AddIProgress(progress IProgress) *list.Element {
	return th.mProgressList.PushBack(progress)
}
func (th *FileSplitter) RemoveIProgress(progressElement *list.Element) {
	th.mProgressList.Remove(progressElement)
}
func (th *FileSplitter) Split() {
	// 1.读取大文件

	// 2.分批次向小文件中写入
	for i := 0; i < th.txtFileNumber; i++ {
		// ...
		log.Printf("do split[%d]", i)
		progressValue := float64((i + 1) / th.txtFileNumber)
		th.onProgress(progressValue)
	}
}
func (th *FileSplitter) onProgress(value float64) {
	// note:循环删除时，需要临时存一下next，因为remove后，e.next会赋值为nil
	for e := th.mProgressList.Front(); e != nil; e.Next() {
		progress := e.Value.(IProgress)
		if progress != nil {
			progress.DoProgress(value)
		}
	}
}

type ProgressBar struct {
	progress float64
}

func (th *ProgressBar) SetValue(progress float64) {
	th.progress = progress
}

// MainForm 文件切割界面
type MainForm struct {
	txtFilePath   string // 要切割的文件路径
	txtFileNumber int    // 要切割的个数

	progressBar *ProgressBar  // MainForm界面和ProgressBar本身是一个整体
	e           *list.Element // 用于将自己从观察者列表移除
}

func NewMainForm(path string, num int) *MainForm {
	return &MainForm{
		txtFilePath:   path,
		txtFileNumber: num,
	}
}

func (th *MainForm) ButtonClickToDoSplit() {
	splitter := NewFileSplitter(th.txtFilePath, th.txtFileNumber)
	th.e = splitter.AddIProgress(th) // 第一个观察者

	cn := &ConsoleNotifier{}
	cn.e = splitter.AddIProgress(cn) // 第二个观察者

	splitter.Split()

	// 移除观察者
	splitter.RemoveIProgress(th.e)
	splitter.RemoveIProgress(cn.e)
}

func (th *MainForm) DoProgress(value float64) {
	// 界面用进度条输出进度
	th.progressBar.SetValue(value)
}

type ConsoleNotifier struct {
	e *list.Element // 用于将自己从观察者列表移除
}

func (th *ConsoleNotifier) DoProgress(value float64) {
	// 控制台用.输出进度
	fmt.Print(".")
}
