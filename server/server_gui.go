package main

import (
	"fmt"
	"os"

	"github.com/therecipe/qt/widgets"
)

var serverLog *widgets.QTextBrowser

func startGUI() {
	widgets.NewQApplication(len(os.Args), os.Args)

	var (
		serverSettings    = widgets.NewQGroupBox2("Server Settings", nil)
		serverAddress     = widgets.NewQLabel2("Address:", nil, 0)
		serverAddressEdit = widgets.NewQLineEdit(nil)

		serverPort     = widgets.NewQLabel2("Port:", nil, 0)
		serverPortEdit = widgets.NewQLineEdit(nil)

		serverStartBtn = widgets.NewQPushButton(nil)
		serverStop     = widgets.NewQPushButton(nil)
	)

	serverLog = widgets.NewQTextBrowser(nil)

	serverAddressEdit.SetPlaceholderText(DEFAULT_URL)
	serverPortEdit.SetPlaceholderText(DEFAULT_PORT)

	serverStartBtn.SetText("Start Server")
	serverStop.SetText("Stop Server")

	var serverSettingsLayout = widgets.NewQGridLayout2()
	serverSettingsLayout.AddWidget(serverAddress, 0, 0, 0)
	serverSettingsLayout.AddWidget(serverAddressEdit, 0, 1, 0)

	serverSettingsLayout.AddWidget(serverPort, 2, 0, 0)
	serverSettingsLayout.AddWidget(serverPortEdit, 2, 1, 0)

	serverSettingsLayout.AddWidget(serverStartBtn, 4, 0, 0)
	serverSettingsLayout.AddWidget(serverStop, 4, 1, 1)

	serverSettingsLayout.AddWidget3(serverLog, 5, 0, 1, 2, 0)

	serverSettings.SetLayout(serverSettingsLayout)

	serverStartBtn.ConnectClicked(func(checked bool) {
		addr := serverAddressEdit.Text()
		port := serverPortEdit.Text()

		if len(addr) == 0 {
			addr = DEFAULT_URL
		}

		if len(port) == 0 {
			port = DEFAULT_PORT
		}

		fullAddr := addr + ":" + port
		go startServer(fullAddr)
	})

	serverStop.ConnectClicked(func(checked bool) {
		fmt.Println("Bout to close the server! (not done yet)")
		stopServer()
	})

	var layout = widgets.NewQGridLayout2()
	layout.AddWidget(serverSettings, 0, 0, 0)

	var window = widgets.NewQMainWindow(nil, 0)
	window.SetWindowTitle("Server")

	var centralWidget = widgets.NewQWidget(window, 0)
	centralWidget.SetLayout(layout)
	window.SetCentralWidget(centralWidget)

	window.Show()

	widgets.QApplication_Exec()
}

func appendToLogger(msg string) {
	fmt.Printf("GUI: %s\n", msg)
	serverLog.InsertPlainText(msg)
}

func cleanLogger() {
	serverLog.Clear()
}
