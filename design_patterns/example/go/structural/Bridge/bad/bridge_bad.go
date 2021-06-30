package bad

import "log"

type Image struct {
}

type IMessager interface {
	Login(username, password string)
	SendMessage(message string)
	SendPicture(image Image)

	PlayerSound()
	DrawShape()
	WriteText()
	Connect()
}

// PCMessagerBase 平台实现
type PCMessagerBase struct {
}

func (th *PCMessagerBase) Login(username, password string) {
	// do nothing
}
func (th *PCMessagerBase) SendMessage(message string) {
	// do nothing
}
func (th *PCMessagerBase) SendPicture(image Image) {
	// do nothing
}

func (th *PCMessagerBase) PlayerSound() {
	log.Println("PCMessagerBase.PlayerSound")
}

func (th *PCMessagerBase) DrawShape() {
	log.Println("PCMessagerBase.DrawShape")
}

func (th *PCMessagerBase) WriteText() {
	log.Println("PCMessagerBase.WriteText")
}

func (th *PCMessagerBase) Connect() {
	log.Println("PCMessagerBase.Connect")
}

// MobileMessagerBase 平台实现
type MobileMessagerBase struct {
}

func (th *MobileMessagerBase) Login(username, password string) {
	// do nothing
}
func (th *MobileMessagerBase) SendMessage(message string) {
	// do nothing
}
func (th *MobileMessagerBase) SendPicture(image Image) {
	// do nothing
}

func (th *MobileMessagerBase) PlayerSound() {
	log.Println("MobileMessagerBase.PlayerSound")
}

func (th *MobileMessagerBase) DrawShape() {
	log.Println("MobileMessagerBase.DrawShape")
}

func (th *MobileMessagerBase) WriteText() {
	log.Println("MobileMessagerBase.WriteText")
}

func (th *MobileMessagerBase) Connect() {
	log.Println("MobileMessagerBase.Connect")
}

// PCMessagerLite 业务抽象：PC精简版
type PCMessagerLite struct {
	PCBase PCMessagerBase
}

func (th *PCMessagerLite) Login(username, password string) {
	th.PCBase.Connect()
	// ====================
}

func (th *PCMessagerLite) SendMessage(message string) {
	th.PCBase.WriteText()
	// ====================
}

func (th *PCMessagerLite) SendPicture(image Image) {
	th.PCBase.DrawShape()
	// ====================
}
func (th *PCMessagerLite) PlayerSound() {
	// do nothing
}

func (th *PCMessagerLite) DrawShape() {
	// do nothing
}

func (th *PCMessagerLite) WriteText() {
	// do nothing
}

func (th *PCMessagerLite) Connect() {
	// do nothing
}

// PCMessagerPerfect 业务抽象：PC完美版
type PCMessagerPerfect struct {
	PCBase PCMessagerBase
}

func (th *PCMessagerPerfect) Login(username, password string) {
	th.PCBase.PlayerSound()
	// ********************
	th.PCBase.Connect()
	// ====================
}

func (th *PCMessagerPerfect) SendMessage(message string) {
	th.PCBase.PlayerSound()
	// **********************
	th.PCBase.WriteText()
	// ====================
}

func (th *PCMessagerPerfect) SendPicture(image Image) {
	th.PCBase.PlayerSound()
	// ********************
	th.PCBase.DrawShape()
	// ====================
}

func (th *PCMessagerPerfect) PlayerSound() {
	// do nothing
}

func (th *PCMessagerPerfect) DrawShape() {
	// do nothing
}

func (th *PCMessagerPerfect) WriteText() {
	// do nothing
}

func (th *PCMessagerPerfect) Connect() {
	// do nothing
}

// Mobile--------------------------------------------------------------------

// MobileMessagerLite 业务抽象：PC精简版
type MobileMessagerLite struct {
	MobileBase MobileMessagerBase
}

func (th *MobileMessagerLite) Login(username, password string) {
	th.MobileBase.Connect()
	// ====================
}

func (th *MobileMessagerLite) SendMessage(message string) {
	th.MobileBase.WriteText()
	// ====================
}

func (th *MobileMessagerLite) SendPicture(image Image) {
	th.MobileBase.DrawShape()
	// ====================
}

func (th *MobileMessagerLite) PlayerSound() {
	// do nothing
}

func (th *MobileMessagerLite) DrawShape() {
	// do nothing
}

func (th *MobileMessagerLite) WriteText() {
	// do nothing
}

func (th *MobileMessagerLite) Connect() {
	// do nothing
}

// MobileMessagerPerfect 业务抽象：PC完美版
type MobileMessagerPerfect struct {
	MobileBase MobileMessagerBase
}

func (th *MobileMessagerPerfect) Login(username, password string) {
	th.MobileBase.PlayerSound()
	// ********************
	th.MobileBase.Connect()
	// ====================
}

func (th *MobileMessagerPerfect) SendMessage(message string) {
	th.MobileBase.PlayerSound()
	// **********************
	th.MobileBase.WriteText()
	// ====================
}

func (th *MobileMessagerPerfect) SendPicture(image Image) {
	th.MobileBase.PlayerSound()
	// ********************
	th.MobileBase.DrawShape()
	// ====================
}

func (th *MobileMessagerPerfect) PlayerSound() {
	// do nothing
}

func (th *MobileMessagerPerfect) DrawShape() {
	// do nothing
}

func (th *MobileMessagerPerfect) WriteText() {
	// do nothing
}

func (th *MobileMessagerPerfect) Connect() {
	// do nothing
}
