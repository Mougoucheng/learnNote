package good

import "log"

type IProgress interface {
	DoProgress(value float64)
}

type FileSplitter struct {
	txtFilePath   string
	txtFileNumber int

	//mProgressBar *ProgressBar // 具体通知控件
	mProgress IProgress // 抽象通知机制
}

func NewFileSplitter(path string, num int, progress IProgress) *FileSplitter {
	return &FileSplitter{
		txtFilePath:   path,
		txtFileNumber: num,
		mProgress:     progress,
	}
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
	if th.mProgress != nil {
		th.mProgress.DoProgress(value)
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

	progressBar *ProgressBar // MainForm界面和ProgressBar本身是一个整体
}

func NewMainForm(path string, num int) *MainForm {
	return &MainForm{
		txtFilePath:   path,
		txtFileNumber: num,
	}
}

func (th *MainForm) ButtonClickToDoSplit() {
	splitter := NewFileSplitter(th.txtFilePath, th.txtFileNumber, th)
	splitter.Split()
}

func (th *MainForm) DoProgress(value float64) {
	th.progressBar.SetValue(value)
}
