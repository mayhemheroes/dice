package fuzzCore

import "strconv"
import "unsafe"
import "github.com/dicedb/dice/core"
import "github.com/dicedb/dice/core/dencoding"
import "github.com/dicedb/dice/testutils"
import fuzz "github.com/AdaLogics/go-fuzz-headers"

func mayhemit(bytes []byte) int {

    var num int
    if len(bytes) > 2 {
        num, _ = strconv.Atoi(string(bytes[0]))
        bytes = bytes[1:]

        switch num {

        case 0:
            dencoding.DecodeUInt(bytes)
            return 0

        case 5:
            fuzzConsumer := fuzz.NewConsumer(bytes)
            var test uint64
            err := fuzzConsumer.CreateSlice(&test)
            if err != nil {
                return 0
            }
            dencoding.EncodeUInt(test)
            return 0

        case 1:
            fuzzConsumer := fuzz.NewConsumer(bytes)
            var content string
            err := fuzzConsumer.CreateSlice(&content)
            if err != nil {
                return 0
            }

            core.Get(content)
            return 0

        case 2:
            fuzzConsumer := fuzz.NewConsumer(bytes)
            var content string
            err := fuzzConsumer.CreateSlice(&content)
            if err != nil {
                return 0
            }

            core.Del(content)
            return 0

        case 3:
            fuzzConsumer := fuzz.NewConsumer(bytes)
            var content string
            err := fuzzConsumer.CreateSlice(&content)
            if err != nil {
                return 0
            }

            var pointer = unsafe.Pointer(&content)
            core.DelByPtr(pointer)
            return 0

        case 4:
            fuzzConsumer := fuzz.NewConsumer(bytes)
            var content string
            err := fuzzConsumer.CreateSlice(&content)
            if err != nil {
                return 0
            }
            
            var test core.StackRef
            test.Push(content)
            return 0

        default:
            split := num
            if split >= len(bytes) {

                split = len(bytes) - 1
            }

            bytes1 := bytes[0:split]
            bytes2 := bytes[split:len(bytes)]

            testutils.EqualByteSlice(bytes1, bytes2)
            return 0
        }
    }
    return 0
}

func Fuzz(data []byte) int {
    _ = mayhemit(data)
    return 0
}