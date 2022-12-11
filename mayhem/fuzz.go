package fuzz

import "strconv"
import "unsafe"
import "github.com/dicedb/dice/core"
import "github.com/dicedb/dice/core/dencoding"

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
        var pointer = unsafe.Pointer(&content)
        var value = uint32(num)
        var test core.EvictionPool
        test.Push(pointer, value)
        return 0

    case 2:
        var n = uint64(num)
        dencoding.EncodeUInt(n)
        return 0

    case 3:
        dencoding.DecodeUInt(bytes)
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