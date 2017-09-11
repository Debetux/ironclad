## What?
Ironclad is a client / server solution for using your mouse and keyboard in order to control another computer. 
e.g. : you're lying on the couch with your Macbook. Now, you want to control the computer connected to your TV without using another keyboard ; for playing a game or starting a movie. 

- **SwiftRemoteControl:** Swift application for macOS in order to send yourmouse and keyboard events to the server.
- **GoControlServer:** Server receive calls from one or multiple clients, and send them to windows with the win32 API.


### GoControlServer:

Requirements:
- Go compiler: http://golang.org/
- MinGW64 – gcc compiler for windows: http://mingw-w64.sourceforge.net/ (yes, mingw64 and not mingw)

```
go get -u github.com/firstrow/tcp_server
go get -u github.com/AllenDang/w32
go install github.com/AllenDang/w32
```


### SwiftRemoteControl:
`⌥` for disconnecting.
`ctrl` for connecting.

Server IP is hardcoded for now in `RemoteControl/ViewController.swift`.
Open `RemoteControl.xcworkspace`. Compile it.


### TODO
Lots of things needs to be done. Better keyboard handling, right click, refactoring Swift app, etc...
I'd like to add a streaming functionnality in order to play my games locally on my mac (without Steam).

Contributions welcome!
