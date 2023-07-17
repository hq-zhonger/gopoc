package main

import (
	"changeme/config"
	"context"
	"fmt"
	"github.com/imroc/req/v3"
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
				a.diag(config.Description, false)
			}),

			menu.Text("检查更新", nil, func(_ *menu.CallbackData) {
				resp, err := req.C().SetUserAgent("Chrome Win7: Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/535.1 (KHTML, like Gecko) Chrome/14.0.835.163 Safari/535.1").R().Get("https://github.com/hq-zhonger/gopoc/version")
				if err != nil {
					a.diag("检查更新出错\n"+err.Error(), true)
					return
				}
				fmt.Println(resp.String())

				needUpdate := config.Version < resp.String()
				msg := config.VersionNewMsg
				btns := []string{config.BtnConfirmText}
				if needUpdate {
					msg = fmt.Sprintf(config.VersionOldMsg, config.Version)
					btns = []string{"确定", "取消"}
				}
				selection, err := a.diag(msg, false, btns...)
				if err != nil {
					return
				}

				if needUpdate && selection == config.BtnConfirmText {
					url := fmt.Sprintf("https://github.com/yhy0/ChYing/releases/tag/%s", config.Version)
					runtime.BrowserOpenURL(a.ctx, url)
				}
			}),

			menu.Text(
				"主页",
				keys.Combo("H", keys.CmdOrCtrlKey, keys.ShiftKey),
				func(_ *menu.CallbackData) {
					runtime.BrowserOpenURL(a.ctx, "https://github.com/yhy0/ChYing/")
				},
			),

			menu.Text(
				"背景",
				keys.Combo("H", keys.CmdOrCtrlKey, keys.ShiftKey),
				func(_ *menu.CallbackData) {

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

// diag ...
func (a *App) diag(message string, error bool, buttons ...string) (string, error) {
	if len(buttons) == 0 {
		buttons = []string{
			config.BtnConfirmText,
		}
	}

	var t runtime.DialogType

	if error {
		t = runtime.ErrorDialog
	} else {
		t = runtime.InfoDialog
	}

	return runtime.MessageDialog(a.ctx, runtime.MessageDialogOptions{
		Type:          t,
		Title:         config.Title,
		Message:       message,
		CancelButton:  config.BtnCancelText,
		DefaultButton: config.BtnConfirmText,
		Buttons:       buttons,
	})
}
