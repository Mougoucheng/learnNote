package good

type Image struct {
}

type IMessager interface {
	Login(username, password string)
	SendMessage(message string)
	SendPicture(image Image)
}

type IMessagerImp interface {
	PlayerSound()
	DrawShape()
	WriteText()
	Connect()
}

// PCMessagerBase 平台实现
type PCMessagerBase struct {
}

func (th *PCMessagerBase) PlayerSound() {
	// ====================
}

func (th *PCMessagerBase) DrawShape() {
	// ====================
}

func (th *PCMessagerBase) WriteText() {
	// ====================
}

func (th *PCMessagerBase) Connect() {
	// ====================
}

// MobileMessagerBase 平台实现
type MobileMessagerBase struct {
}

func (th *MobileMessagerBase) PlayerSound() {
	// ====================
}

func (th *MobileMessagerBase) DrawShape() {
	// ====================
}

func (th *MobileMessagerBase) WriteText() {
	// ====================
}

func (th *MobileMessagerBase) Connect() {
	// ====================
}

// MessagerLite 业务抽象：PC精简版
type MessagerLite struct {
	imp IMessagerImp
}

func (th *MessagerLite) Login(username, password string) {
	th.imp.Connect()
	// ====================
}

func (th *MessagerLite) SendMessage(message string) {
	th.imp.WriteText()
	// ====================
}

func (th *MessagerLite) SendPicture(image Image) {
	th.imp.DrawShape()
	// ====================
}


// MessagerPerfect 业务抽象：PC/Mobile完美版
type MessagerPerfect struct {
	imp IMessagerImp
}

func (th *MessagerPerfect) Login(username, password string) {
	th.imp.PlayerSound()
	// ********************
	th.imp.Connect()
	// ====================
}

func (th *MessagerPerfect) SendMessage(message string) {
	th.imp.PlayerSound()
	// **********************
	th.imp.WriteText()
	// ====================
}

func (th *MessagerPerfect) SendPicture(image Image) {
	th.imp.PlayerSound()
	// ********************
	th.imp.DrawShape()
	// ====================
}
