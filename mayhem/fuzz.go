package fuzz

import "strconv"
import "github.com/dicedb/dice/core"

func mayhemit(bytes []byte) int {

    var num int
    if len(bytes) < 1 {
        num = 0
    } else {
        num, _ = strconv.Atoi(string(bytes[0]))
    }

    switch num {
    
    case 0:
        var test core.Client
        test.Write(bytes)
        return 0

    case 1:
        content := string(bytes)
        var value uint32 = uint32(num)
        var test core.EvictionPool
        test.Push(content, value)
        return 0

    case 2:
        core.DecodeOne(bytes)
        return 0

    case 3:
        core.Decode(bytes)
        return 0

    case 4:
        content := string(bytes)
        core.Get(content)
        return 0

    case 5:
        var test core.Client
        test.Read(bytes)
        return 0

    default:
        content := string(bytes)
        core.Del(content)
        return 0
    }
}

func Fuzz(data []byte) int {
    _ = mayhemit(data)
    return 0
}