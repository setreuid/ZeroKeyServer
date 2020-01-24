package main

import (
    pb "ZeroKey/data"
    "encoding/json"
    "io/ioutil"
)

var (
    Preset = pb.KeysRequest{}
    MainKeyboard = makeKeyboard()
)


func main() {
    var server Server

    if file, err := ioutil.ReadFile("preset.json"); err == nil {
        if err := json.Unmarshal([]byte(file), &Preset); err == nil {
            _P("Load preset.json")
        } else {
            _P("Load json err: %v", err)
        }
    }

    _P("Open server at port %s", PORT)
    server.Start()

    // MainKeyboard.Test()
}
