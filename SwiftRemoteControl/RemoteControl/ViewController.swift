//
//  ViewController.swift
//  RemoteControl
//
//  Created by Lancelot HARDEL on 10/09/17.
//  Copyright © 2017 Lancelot HARDEL. All rights reserved.
//

import Cocoa
import SwiftSocket
import Foundation
import SwiftyJSON

struct KeyboardEvent {
    var KeyCode = 0
    var Status = 0
}

struct CursorPosition {
    var PosX = 0.0
    var PosY = 0.0
    var MouseLeftStatus = 0
    var Keys = [KeyboardEvent]()
}

extension String {
    var asciiArray: [UInt32] {
        return unicodeScalars.filter{$0.isASCII}.map{$0.value}
    }
}
extension Character {
    var asciiValue: UInt32? {
        return String(self).unicodeScalars.filter{$0.isASCII}.first?.value
    }
}

class ViewController: NSViewController {

    var mouseLocation: CGPoint = .zero
    var client: TCPClient = TCPClient(address: "192.168.0.33", port: 8080);
    var state: CursorPosition = CursorPosition(PosX: 0.0, PosY: 0.0, MouseLeftStatus: 0, Keys:[])
    var timeout: Int = 0
    
    override func viewDidLoad() {
        super.viewDidLoad()
        
        self.connect()

        // Do any additional setup after loading the view.
        NSEvent.addLocalMonitorForEvents(matching: .flagsChanged) {
            self.flagsChanged(with: $0)
            return $0
        }
        
        NSEvent.addGlobalMonitorForEvents(matching: .flagsChanged) {
            self.flagsChanged(with: $0)
        }
        
        NSEvent.addLocalMonitorForEvents(matching: .keyDown) {
            self.keyDown(with: $0)
            return $0
        }
        
        NSEvent.addLocalMonitorForEvents(matching: .keyUp) {
            self.keyUp(with: $0)
            return $0
        }
        
        
        NSEvent.addLocalMonitorForEvents(matching: [.mouseMoved]) {
            self.mouseLocation = NSEvent.mouseLocation()
            print(String(format: "l%.0f, %.0f", self.mouseLocation.x, self.mouseLocation.y))
            self.state.PosX = Double(self.mouseLocation.x)
            self.state.PosY = 900 - Double(self.mouseLocation.y)
            self.sendData()
            return $0
        }
        NSEvent.addGlobalMonitorForEvents(matching: [.mouseMoved]) { _ in
            self.mouseLocation = NSEvent.mouseLocation()
            print(String(format: "g%.0f, %.0f", self.mouseLocation.x, self.mouseLocation.y))
            self.state.PosX = Double(self.mouseLocation.x)
            self.state.PosY = 900 - Double(self.mouseLocation.y)
            self.sendData()
        }
        
        NSEvent.addGlobalMonitorForEvents(matching: [.rightMouseDown]) { (_) in
            print(String(format:"Right Mouse Down"));
            self.state.MouseLeftStatus = 1
        }
        
        NSEvent.addGlobalMonitorForEvents(matching: [.leftMouseDown]) { (_) in
            print(String(format:"Left Mouse Down"));
            self.state.MouseLeftStatus = 1
            self.sendData()
        }
        
        NSEvent.addGlobalMonitorForEvents(matching: [.leftMouseDragged]) { (_) in
            print(String(format:"Left Mouse Dragged"));
            if(self.state.MouseLeftStatus == 0) {
                self.state.MouseLeftStatus = 1
            }
            
            if(self.state.MouseLeftStatus == 1) {
                self.state.MouseLeftStatus = 3;
            }
        
            self.mouseLocation = NSEvent.mouseLocation()
            self.state.PosX = Double(self.mouseLocation.x)
            self.state.PosY = 900 - Double(self.mouseLocation.y)
            self.sendData()
        }
        
        NSEvent.addGlobalMonitorForEvents(matching: [.leftMouseUp]) { (_) in
            print(String(format:"Left Mouse Up"));
            self.state.MouseLeftStatus = 2
            self.sendData()
        }
    }

    override var representedObject: Any? {
        didSet {
        // Update the view, if already loaded.
        }
    }

    override func keyDown(with event: NSEvent) {
        let scalars = String(event.characters!)?.unicodeScalars
        self.state.Keys.append(KeyboardEvent(KeyCode: Int((scalars?[(scalars?.startIndex)!].value)!), Status:1))
        self.sendData()
        switch event.modifierFlags.intersection(.deviceIndependentFlagsMask) {
        case [.command] where event.characters == "l",
             [.command, .shift] where event.characters == "l":
            print("command-l or command-shift-l")
        default:
            break
        }
    }
    
    override func keyUp(with event: NSEvent) {
        //let scalars = String(event.characters!)?.unicodeScalars
        //self.state.Keys.append(KeyboardEvent(KeyCode: Int((scalars?[(scalars?.startIndex)!].value)!), Status:2))
        print(event.characters!)
        //self.sendData()
        switch event.modifierFlags.intersection(.deviceIndependentFlagsMask) {
        case [.command] where event.characters == "l",
             [.command, .shift] where event.characters == "l":
            print("command-l or command-shift-l")
        default:
            break
        }
    }

    override func flagsChanged(with event: NSEvent) {
        switch event.modifierFlags.intersection(.deviceIndependentFlagsMask) {
        case [.shift]:
            print("shift key is pressed")
        case [.control]:
            print("control key is pressed")
            self.connect()
        case [.option] :
            print("option key is pressed")
            self.disconnect()
        case [.command]:
            print("Command key is pressed")
        case [.control, .shift]:
            print("control-shift keys are pressed")
        case [.option, .shift]:
            print("option-shift keys are pressed")
        case [.command, .shift]:
            print("command-shift keys are pressed")
        case [.control, .option]:
            print("control-option keys are pressed")
        case [.control, .command]:
            print("control-command keys are pressed")
        case [.option, .command]:
            print("option-command keys are pressed")
        case [.shift, .control, .option]:
            print("shift-control-option keys are pressed")
        case [.shift, .control, .command]:
            print("shift-control-command keys are pressed")
        case [.control, .option, .command]:
            print("control-option-command keys are pressed")
        case [.shift, .command, .option]:
            print("shift-command-option keys are pressed")
        case [.shift, .control, .option, .command]:
            print("shift-control-option-command keys are pressed")
        default:
            print("no modifier keys are pressed")
        }
    }
    
    func connect() -> Void {
        let when = DispatchTime.now() + .seconds(self.timeout)
        DispatchQueue.main.asyncAfter(deadline: when) {
            self.client.close()
            switch self.client.connect(timeout: 10) {
            case .success:
                print("CONNECTED");
            case .failure(_):
                print("CONNECT FAIL")
                self.timeout += 10
            }
        }
    }
    
    func disconnect() -> Void {
        self.client.close()
    }
    
    func sendData() -> Int {
        let keysArray = "[" + state.Keys.map {String(format:"{\"KeyCode\": \"0x%x\", \"Status\": %x}", Int($0.KeyCode), $0.Status)}.joined(separator: ",") + "]";
        let json = String(format: "{\"PosX\": %.0f, \"PosY\": %.0f, \"MouseLeftStatus\":%x, \"Keys\": " + keysArray + " }\n", self.state.PosX * 1.375, self.state.PosY * 1.2, self.state.MouseLeftStatus)
        print(json)
        self.state.Keys = []
        if(self.state.MouseLeftStatus == 2) {
            self.state.MouseLeftStatus = 0;
        }
        switch self.client.send(string: json) {
        case .success:
            return 1;
        case .failure(_):
            
            return 0;
            
        }
        return 0
    }

}

