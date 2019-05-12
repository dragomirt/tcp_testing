package main

import (
	"fmt"
	"os"

	"github.com/therecipe/qt/widgets"
)

var serverConn *widgets.QTextBrowser
var serverMsgEdit *widgets.QLineEdit

func startGUI() {
	widgets.NewQApplication(len(os.Args), os.Args)

	var (
		clientSettings    = widgets.NewQGroupBox2("Server Settings", nil)
		serverAddress     = widgets.NewQLabel2("Address:", nil, 0)
		serverAddressEdit = widgets.NewQLineEdit(nil)

		serverPort     = widgets.NewQLabel2("Port:", nil, 0)
		serverPortEdit = widgets.NewQLineEdit(nil)

		serverClientNameLabel = widgets.NewQLabel2("Client Name:", nil, 0)
		serverClientNameEdit  = widgets.NewQLineEdit(nil)

		clientMsgHistory = widgets.NewQLabel2("Message History:", nil, 0)

		clientConnect    = widgets.NewQPushButton(nil)
		clientDisconnect = widgets.NewQPushButton(nil)
		clientSendMsg    = widgets.NewQPushButton(nil)
	)

	serverConn = widgets.NewQTextBrowser(nil)
	serverMsgEdit = widgets.NewQLineEdit(nil)

	serverAddressEdit.SetPlaceholderText(DEFAULT_URL)
	serverPortEdit.SetPlaceholderText(DEFAULT_PORT)
	serverMsgEdit.SetPlaceholderText("Type your msg...")

	clientConnect.SetText("Connect")
	clientDisconnect.SetText("Disconnect")
	clientSendMsg.SetText("Send")

	var clientSettingsLayout = widgets.NewQGridLayout2()
	clientSettingsLayout.AddWidget(serverAddress, 0, 0, 0)
	clientSettingsLayout.AddWidget(serverAddressEdit, 0, 1, 0)

	clientSettingsLayout.AddWidget(serverPort, 2, 0, 0)
	clientSettingsLayout.AddWidget(serverPortEdit, 2, 1, 0)

	clientSettingsLayout.AddWidget(serverClientNameLabel, 3, 0, 0)
	clientSettingsLayout.AddWidget(serverClientNameEdit, 3, 1, 0)

	clientSettingsLayout.AddWidget3(clientConnect, 4, 0, 1, 2, 0)
	// clientSettingsLayout.AddWidget(clientDisconnect, 4, 1, 1)

	clientSettingsLayout.AddWidget3(clientMsgHistory, 5, 0, 1, 2, 0)

	clientSettingsLayout.AddWidget3(serverConn, 6, 0, 1, 2, 0)

	clientSettingsLayout.AddWidget3(serverMsgEdit, 7, 0, 1, 2, 0)
	clientSettingsLayout.AddWidget3(clientSendMsg, 8, 0, 1, 2, 0)

	clientSettings.SetLayout(clientSettingsLayout)

	clientConnect.ConnectClicked(func(checked bool) {
		addr := serverAddressEdit.Text()
		port := serverPortEdit.Text()

		if len(addr) == 0 {
			addr = DEFAULT_URL
		}

		if len(port) == 0 {
			port = DEFAULT_PORT
		}

		fullAddr := addr + ":" + port
		go startClient(fullAddr)
	})

	clientDisconnect.ConnectClicked(func(checked bool) {
		stopClient()
	})

	clientSendMsg.ConnectClicked(func(checked bool) {
		// fmt.Println(serverMsgEdit.Text())
		writeToServer(fmt.Sprintf(("%s -> %s"), serverClientNameEdit.Text(), serverMsgEdit.Text()))
		serverMsgEdit.Clear()
	})

	var layout = widgets.NewQGridLayout2()
	layout.AddWidget(clientSettings, 0, 0, 0)

	var window = widgets.NewQMainWindow(nil, 0)
	window.SetWindowTitle("Client")

	var centralWidget = widgets.NewQWidget(window, 0)
	centralWidget.SetLayout(layout)
	window.SetCentralWidget(centralWidget)

	window.Show()

	widgets.QApplication_Exec()
}

func appendToLogger(msg string) {
	fmt.Printf("GUI: %s\n", msg)
	serverConn.InsertPlainText(msg)
}

func cleanLogger() {
	serverConn.Clear()
}
