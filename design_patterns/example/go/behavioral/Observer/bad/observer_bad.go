package bad

import "log"

type FileSplitter struct {
	txtFilePath   string
	txtFileNumber int

	// 新增进度条需求
	// 违背依赖倒置原则：依赖了实现细节（可能会变化，比如不用progressBar，变成百分比显示）
	mProgressBar *ProgressBar // 具体通知控件
}

func NewFileSplitter(path string, num int, progressBar *ProgressBar) *FileSplitter {
	return &FileSplitter{
		txtFilePath:   path,
		txtFileNumber: num,
		mProgressBar:  progressBar,
	}
}

func (th *FileSplitter) split() {
	// 1.读取大文件

	// 2.分批次向小文件中写入
	for i := 0; i < th.txtFileNumber; i++ {
		// ...
		log.Printf("do split[%d]", i)
		if th.mProgressBar != nil {
			th.mProgressBar.SetValue(float64((i + 1) / th.txtFileNumber))
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

	progressBar *ProgressBar
}

func NewMainForm(path string, num int) *MainForm {
	return &MainForm{
		txtFilePath:   path,
		txtFileNumber: num,
	}
}

func (th *MainForm) ButtonClickToDoSplit() {
	splitter := NewFileSplitter(th.txtFilePath, th.txtFileNumber, th.progressBar)
	splitter.split()
}
