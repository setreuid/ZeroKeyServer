package main

import (
	"os"
	"log"
	"time"
    "runtime"
)


type Keyboard struct {
    number int
    isRunning bool
    tasks []StepKey
    modPressed map[int]byte // Pressed mod keys (ex: ctrl, shift...)
    pressed map[int]byte // Pressed keys
    runCount int
}


type StepKey struct {
    code int
    state int
    time uint64
}


func makeKeyboard() *Keyboard {
    return &Keyboard{
        number: 0,
        isRunning: false,
    }
}


func (self *Keyboard) Write(data []byte) {
	hid, err := os.Create("/dev/hidg0")
	defer hid.Close()
	if err != nil {
		log.Fatal(err)
	}

	hid.Write(data)
}


func (self *Keyboard) Run(num int) {
    if self.isRunning {
        _P("Already running on %d", num)
        return
    }

    self.isRunning = true
    self.number = num
    self.runCount = 0

    go self.InitStage(num)
}


func (self *Keyboard) Stop(num int) {
    self.isRunning = false
    self.Write(rep(KEY_NULL, 8))
}


func (self *Keyboard) InitStage(num int) {
    self.tasks = self.tasks[:0]
    self.modPressed = map[int]byte{}
    self.pressed = map[int]byte{}
    self.runCount++

    _P("Running count %d. go(%d)", self.runCount, runtime.NumGoroutine())

    self.Write(rep(KEY_NULL, 8))

    stTime := GetTime()
    for _, v := range Preset.Acts {
        self.tasks = append(self.tasks, StepKey{
            code: int(v.Code),
            state: int(v.State),
            time: stTime + uint64(v.Time),
        })
    }

    for self.isRunning {
        if !self.Step(num) {
            break
        } else {
            time.Sleep(50 * time.Millisecond)
        }
    }

    if self.isRunning {
        self.InitStage(num)
    }
}


func (self *Keyboard) Step(num int) bool {
    if len(self.tasks) > 0 {
        if self.tasks[0].time < GetTime() {
            task := self.tasks[0]
            self.tasks = self.tasks[1:]

            if task.state == 1 { // DOWN
                self.KeyDown(task.code)
            } else if task.state == 2 { // UP
                self.KeyUp(task.code)
            }
        }

        return true
    } else {
        return false
    }
}


func (self *Keyboard) KeyDown(code int) {
    if self.IsModKey(code) {
        self.modPressed[code] = keyMods[code]
    } else if val, ok := keyStore[code]; ok {
        self.pressed[code] = val
    }

    self.SendPressedKeys()
}


func (self *Keyboard) KeyUp(code int) {
    if self.IsModKey(code) {
        delete(self.modPressed, code)
    } else if _, ok := keyStore[code]; ok {
        delete(self.pressed, code)
    }

    self.SendPressedKeys()
}


func (self *Keyboard) IsModKey(code int) bool {
    if _, ok := keyMods[code]; ok {
        return true
    } else {
        return false
    }
}


func (self *Keyboard) SendPressedKeys() {
    var pKeys = make([]byte, 8, 8)
    var modKey byte = 0x00

    for _, v := range self.modPressed {
        modKey = modKey | v
    }

    pKeys[0] = modKey

    var i = 2
    for _, v := range self.pressed {
        if i > 7 {
            break
        }

        pKeys[i] = v
        i++
    }

    self.Write(pKeys)
}


func (self *Keyboard) Test() {
    _P("Test after 1s")
    time.Sleep(1 * time.Second)

	// Press a
	// self.Write([]byte{KEY_NULL, KEY_NULL, 0x04, KEY_NULL, KEY_NULL, KEY_NULL, KEY_NULL, KEY_NULL})

	// Release keys
	// self.Write(rep(KEY_NULL, 8))

	// Press SHIFT + a = A
	// self.Write(append([]byte{0x20, KEY_NULL, 0x04}, rep(KEY_NULL, 5)...))

    // Press Ctrl + 1
    // self.Write(append([]byte{KEY_MOD_LCTRL, KEY_NULL, KEY_1}, rep(KEY_NULL, 5)...))

    // Press Ctrl + Alt + Delete
    self.Write(append([]byte{KEY_MOD_LCTRL|KEY_MOD_LALT, KEY_NULL, KEY_DELETE}, rep(KEY_NULL, 5)...))

	// Release all keys
	self.Write(rep(KEY_NULL, 8))
}
