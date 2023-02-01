// +build go1.15,!go1.16

/*
 * Copyright 2022 ByteDance Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

// This is a workaround agains the bug of "go:linkname" in Go 1.15.

package rtx

import (
    `sync/atomic`
    `unsafe`

    `github.com/cloudwego/frugal/internal/rt`
    `golang.org/x/arch/x86/x86asm`
)

const (
    _MaxInstr = 15
)

func isvar(arg x86asm.Arg) bool {
    v, ok := arg.(x86asm.Mem)
    return ok && v.Base == x86asm.RIP
}

func iszero(arg x86asm.Arg) bool {
    v, ok := arg.(x86asm.Imm)
    return ok && v == 0
}

func gcwbaddr() unsafe.Pointer {
    var err error
    var off uintptr
    var ins x86asm.Inst

    /* get the function address */
    pc := uintptr(0)
    fp := rt.FuncAddr(atomic.StorePointer)

    /* search within the first 16 instructions */
    for i := 0; i < 16; i++ {
        mem := unsafe.Pointer(uintptr(fp) + pc)
        buf := rt.BytesFrom(mem, _MaxInstr, _MaxInstr)

        /* disassemble the instruction */
        if ins, err = x86asm.Decode(buf, 64); err != nil {
            panic("gcwbaddr: " + err.Error())
        }

        /* check for a byte comparison with zero */
        if ins.Op == x86asm.CMP && ins.MemBytes == 1 && isvar(ins.Args[0]) && iszero(ins.Args[1]) {
            off = pc + uintptr(ins.Len) + uintptr(ins.Args[0].(x86asm.Mem).Disp)
            break
        }

        /* move to next instruction */
        nb := ins.Len
        pc += uintptr(nb)
    }

    /* check for address */
    if off == 0 {
        panic("gcwbaddr: could not locate the variable `writeBarrier`")
    } else {
        return unsafe.Pointer(uintptr(fp) + off)
    }
}
