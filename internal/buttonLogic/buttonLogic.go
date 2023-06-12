package buttonLogic

import (
	"fyne.io/fyne/v2/widget"
	"log"
	"os"
)

func SignIn(l, p *widget.Entry) {
	login := l.Text
	password := p.Text

	send_login_sign(login)
	send_password_sign(password)
}

func Registration(l, p *widget.Entry) {
	login := l.Text
	password := p.Text
	_, err := os.Create("log.txt")
	if err != nil {
		log.Fatal(err)
	}

	send_login_registation(login)
	send_password_registration(password)
}

func NewLayout() {

}
