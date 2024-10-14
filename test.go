package test

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/config"
	"github.com/GoAdminGroup/go-admin/modules/language"
	"github.com/GoAdminGroup/go-admin/modules/service"
	"github.com/GoAdminGroup/go-admin/modules/utils"
	"github.com/GoAdminGroup/go-admin/plugins"
	form2 "github.com/GoAdminGroup/go-admin/plugins/admin/modules/form"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/beesoft/goadmin-plugin-test/controller"
	"github.com/beesoft/goadmin-plugin-test/guard"
	language2 "github.com/beesoft/goadmin-plugin-test/modules/language"
)

type Test struct {
	*plugins.Base

	handler *controller.Handler
	guard   *guard.Guardian

	// ...
}

func init() {
	plugins.Add(&Test{
		Base: &plugins.Base{PlugName: "test", URLPrefix: "test"},
		// ....
	})
}

func NewTest( /*...*/ ) *Test {
	return &Test{
		Base: &plugins.Base{PlugName: "test", URLPrefix: "test"},
		// ...
	}
}

func (plug *Test) IsInstalled() bool {
	// ... implement it
	return true
}

func (plug *Test) GetIndexURL() string {
	return config.Url("/test/example?param=helloworld")
}

func (plug *Test) InitPlugin(srv service.List) {

	// DO NOT DELETE
	plug.InitBase(srv, "test")

	plug.handler = controller.NewHandler( /*...*/ )
	plug.guard = guard.New( /*...*/ )
	plug.App = plug.initRouter(srv)
	plug.handler.HTML = plug.HTML
	plug.handler.HTMLMenu = plug.HTMLMenu

	language.Lang[language.CN].Combine(language2.CN)
	language.Lang[language.EN].Combine(language2.EN)

	plug.SetInfo(info)
}

var info = plugins.Info{
	Website:     "",
	Title:       "Test",
	Description: "",
	Version:     "",
	Author:      "",
	Url:         "",
	Cover:       "",
	Agreement:   "",
	Uuid:        "",
	Name:        "",
	ModulePath:  "",
	CreateDate:  utils.ParseTime("2000-01-11"),
	UpdateDate:  utils.ParseTime("2000-01-11"),
}

func (plug *Test) GetSettingPage() table.Generator {
	return func(ctx *context.Context) (testConfiguration table.Table) {

		cfg := table.DefaultConfigWithDriver(config.GetDatabases().GetDefault().Driver)

		if !plug.IsInstalled() {
			cfg = cfg.SetOnlyNewForm()
		} else {
			cfg = cfg.SetOnlyUpdateForm()
		}

		testConfiguration = table.NewDefaultTable(cfg)

		formList := testConfiguration.GetForm().
			AddXssJsFilter().
			HideBackButton().
			HideContinueNewCheckBox().
			HideResetButton()

		// formList.AddField(...)

		formList.SetInsertFn(func(values form2.Values) error {
			// TODO: finish the installation
			return nil
		})

		formList.EnableAjaxData(types.AjaxData{
			SuccessTitle:   language2.Get("install success"),
			ErrorTitle:     language2.Get("install fail"),
			SuccessJumpURL: "...",
		}).SetFormNewTitle(language2.GetHTML("test installation")).
			SetTitle(language2.Get("test installation")).
			SetFormNewBtnWord(language2.GetHTML("install"))

		return
	}
}
