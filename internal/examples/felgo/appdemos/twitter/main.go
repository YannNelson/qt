package main

import (
	"os"
	"strings"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/felgo"
	"github.com/therecipe/qt/qml"
	"github.com/therecipe/qt/widgets"
)

func main() {
	widgets.NewQApplication(len(os.Args), os.Args)
	felgoApp := felgo.NewFelgoApplication(nil)

	// Use platform-specific fonts instead of Felgo's default font
	felgoApp.SetPreservePlatformFonts(true)

	// QQmlApplicationEngine is the preferred way to start qml projects since Qt 5.2
	// if you have older projects using Qt App wizards from previous QtCreator versions than 3.1, please change them to QQmlApplicationEngine
	engine := qml.NewQQmlApplicationEngine(nil)
	felgoApp.Initialize(engine)

	// use this during development
	// for PUBLISHING, use the entry point below
	if pwd, _ := os.Getwd(); strings.Contains(pwd, "/deploy/") || core.QSysInfo_ProductType() == "ios" || core.QSysInfo_ProductType() == "android" {
		felgoApp.SetMainQmlFileName("qml/TwitterMain.qml") //to make qtdeploy runs work
	} else {
		felgoApp.SetMainQmlFileName(pwd + "/qml/TwitterMain.qml") //to make go run/build runs work
	}

	// use this instead of the above call to avoid deployment of the qml files and compile them into the binary with qt's resource system qrc
	// this is the preferred deployment option for publishing games to the app stores, because then your qml files and js files are protected
	// felgoApp.SetMainQmlFileName("qrc:/qml/TwitterMain.qml")

	engine.Load(core.NewQUrl3(felgoApp.MainQmlFileName(), 0))

	// to start your project as Live Client, comment (remove) the lines "felgoApp.SetMainQmlFileName ..." & "engine.Load ...",
	// and uncomment the line below
	//felgo.NewFelgoLiveClient(engine, nil)

	widgets.QApplication_Exec()
}
