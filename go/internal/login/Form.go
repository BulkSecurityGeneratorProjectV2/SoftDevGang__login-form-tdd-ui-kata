package login

import (
	"github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type FormUI interface {
	Button(id string, bounds rl.Rectangle, text string) bool
	TextBox(id string, bounds rl.Rectangle, text string) string
	Label(id string, bounds rl.Rectangle, text string)
}

type RaylibFormUI struct{}

func (ui *RaylibFormUI) Button(id string, bounds rl.Rectangle, text string) bool {
	return raygui.Button(rl.Rectangle{}, text)
}

func (ui *RaylibFormUI) TextBox(id string, bounds rl.Rectangle, text string) string {
	return raygui.TextBox(bounds, text)
}

func (ui *RaylibFormUI) Label(id string, bounds rl.Rectangle, text string) {
	raygui.Label(bounds, text)
}

const UserNameLimit = 20

// Form renders the controls for providing the necessary login credentials.
type Form struct {
	UserName string // Go style export more than we usually do.
	Password string
}

func (form *Form) Render(ui FormUI) {
	ui.Label("username", rl.Rectangle{}, "Phone, email or username")
	form.UserName = ui.TextBox("username", rl.Rectangle{
		X:      0,
		Y:      40,
		Width:  100,
		Height: 30,
	}, form.UserName)
	if len(form.UserName) > UserNameLimit {
		form.UserName = form.UserName[:UserNameLimit]
	}

	form.Password = ui.TextBox("password", rl.Rectangle{
		X:      0,
		Y:      80,
		Width:  100,
		Height: 30,
	}, form.Password)

	ui.Button("login", rl.Rectangle{235, 165, 345, 195}, "Log in")
}
