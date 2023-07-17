package main

import (
	"context"
	"fmt"
	"github.com/wailsapp/wails/v2/pkg/menu"
	"github.com/wailsapp/wails/v2/pkg/menu/keys"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) Menu() *menu.Menu {
	return menu.NewMenuFromItems(
		menu.SubMenu("设置", menu.NewMenuFromItems(
			menu.Text("关于", nil, func(_ *menu.CallbackData) {
				runtime.MessageDialog(a.ctx, runtime.MessageDialogOptions{
					Type:          "",
					Title:         "",
					Message:       "message",
					CancelButton:  "",
					DefaultButton: "",
					Buttons:       nil,
				})
			}),
			menu.Text("检查更新", nil, func(_ *menu.CallbackData) {

			}),
			menu.Text(
				"主页",
				keys.Combo("H", keys.CmdOrCtrlKey, keys.ShiftKey),
				func(_ *menu.CallbackData) {
					runtime.BrowserOpenURL(a.ctx, "https://github.com/yhy0/ChYing/")
				},
			),
			menu.Separator(),
			menu.Text("退出", keys.CmdOrCtrl("Q"), func(_ *menu.CallbackData) {
				runtime.Quit(a.ctx)
			}),
		)),

		menu.EditMenu(),
		menu.SubMenu("Help", menu.NewMenuFromItems(
			menu.Text(
				"打开配置文件夹",
				keys.Combo("C", keys.CmdOrCtrlKey, keys.ShiftKey),
				func(_ *menu.CallbackData) {

				},
			),
		)),
	)
}
