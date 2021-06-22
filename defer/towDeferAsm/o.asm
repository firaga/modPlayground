TEXT main.main(SB) /Users/firaga/Gode/modPlayground/defer/twoDefer.go
  twoDefer.go:7         0x10a2ea0               65488b0c2530000000              MOVQ GS:0x30, CX
  twoDefer.go:7         0x10a2ea9               488d4424d0                      LEAQ -0x30(SP), AX
  twoDefer.go:7         0x10a2eae               483b4110                        CMPQ 0x10(CX), AX
  twoDefer.go:7         0x10a2eb2               0f8689010000                    JBE 0x10a3041
  twoDefer.go:7         0x10a2eb8               4881ecb0000000                  SUBQ $0xb0, SP
  twoDefer.go:7         0x10a2ebf               4889ac24a8000000                MOVQ BP, 0xa8(SP)
  twoDefer.go:7         0x10a2ec7               488dac24a8000000                LEAQ 0xa8(SP), BP
  twoDefer.go:7         0x10a2ecf               0f57c0                          XORPS X0, X0
  twoDefer.go:7         0x10a2ed2               0f11442468                      MOVUPS X0, 0x68(SP)
  twoDefer.go:7         0x10a2ed7               0f11442478                      MOVUPS X0, 0x78(SP)
  twoDefer.go:7         0x10a2edc               0f11842488000000                MOVUPS X0, 0x88(SP)
  twoDefer.go:7         0x10a2ee4               0f11842498000000                MOVUPS X0, 0x98(SP)
  twoDefer.go:7         0x10a2eec               c644243700                      MOVB $0x0, 0x37(SP)
  twoDefer.go:8         0x10a2ef1               0f57c0                          XORPS X0, X0
  twoDefer.go:8         0x10a2ef4               0f11442448                      MOVUPS X0, 0x48(SP)
  twoDefer.go:8         0x10a2ef9               488d05c0b00000                  LEAQ runtime.types+44608(SB), AX
  twoDefer.go:8         0x10a2f00               4889442448                      MOVQ AX, 0x48(SP)
  twoDefer.go:8         0x10a2f05               488d0de4370400                  LEAQ sync/atomic.CompareAndSwapUintptr.args_stackmap+192(SB), CX
  twoDefer.go:8         0x10a2f0c               48894c2450                      MOVQ CX, 0x50(SP)
  twoDefer.go:8         0x10a2f11               488d0df0c00200                  LEAQ go.func.*+16(SB), CX
  twoDefer.go:8         0x10a2f18               48894c2470                      MOVQ CX, 0x70(SP)
  twoDefer.go:8         0x10a2f1d               488d542448                      LEAQ 0x48(SP), DX
  twoDefer.go:8         0x10a2f22               4889942490000000                MOVQ DX, 0x90(SP)
  twoDefer.go:8         0x10a2f2a               48c784249800000001000000        MOVQ $0x1, 0x98(SP)
  twoDefer.go:8         0x10a2f36               48c78424a000000001000000        MOVQ $0x1, 0xa0(SP)
  twoDefer.go:8         0x10a2f42               c644243701                      MOVB $0x1, 0x37(SP)
  twoDefer.go:9         0x10a2f47               0f11442438                      MOVUPS X0, 0x38(SP)
  twoDefer.go:9         0x10a2f4c               4889442438                      MOVQ AX, 0x38(SP)
  twoDefer.go:9         0x10a2f51               488d15a8370400                  LEAQ sync/atomic.CompareAndSwapUintptr.args_stackmap+208(SB), DX
  twoDefer.go:9         0x10a2f58               4889542440                      MOVQ DX, 0x40(SP)
  twoDefer.go:9         0x10a2f5d               48894c2468                      MOVQ CX, 0x68(SP)
  twoDefer.go:9         0x10a2f62               488d4c2438                      LEAQ 0x38(SP), CX
  twoDefer.go:9         0x10a2f67               48894c2478                      MOVQ CX, 0x78(SP)
  twoDefer.go:9         0x10a2f6c               48c784248000000001000000        MOVQ $0x1, 0x80(SP)
  twoDefer.go:9         0x10a2f78               48c784248800000001000000        MOVQ $0x1, 0x88(SP)
  twoDefer.go:9         0x10a2f84               c644243703                      MOVB $0x3, 0x37(SP)
  twoDefer.go:11        0x10a2f89               0f11442458                      MOVUPS X0, 0x58(SP)
  twoDefer.go:11        0x10a2f8e               4889442458                      MOVQ AX, 0x58(SP)
  twoDefer.go:11        0x10a2f93               488d0576370400                  LEAQ sync/atomic.CompareAndSwapUintptr.args_stackmap+224(SB), AX
  twoDefer.go:11        0x10a2f9a               4889442460                      MOVQ AX, 0x60(SP)
  twoDefer.go:11        0x10a2f9f               488d442458                      LEAQ 0x58(SP), AX
  twoDefer.go:11        0x10a2fa4               48890424                        MOVQ AX, 0(SP)
  twoDefer.go:11        0x10a2fa8               48c744240801000000              MOVQ $0x1, 0x8(SP)
  twoDefer.go:11        0x10a2fb1               48c744241001000000              MOVQ $0x1, 0x10(SP)
  twoDefer.go:11        0x10a2fba               e8419affff                      CALL fmt.Println(SB)
  twoDefer.go:12        0x10a2fbf               c644243701                      MOVB $0x1, 0x37(SP)
  twoDefer.go:12        0x10a2fc4               488b442478                      MOVQ 0x78(SP), AX
  twoDefer.go:12        0x10a2fc9               488b8c2480000000                MOVQ 0x80(SP), CX
  twoDefer.go:12        0x10a2fd1               488b942488000000                MOVQ 0x88(SP), DX
  twoDefer.go:12        0x10a2fd9               48890424                        MOVQ AX, 0(SP)
  twoDefer.go:12        0x10a2fdd               48894c2408                      MOVQ CX, 0x8(SP)
  twoDefer.go:12        0x10a2fe2               4889542410                      MOVQ DX, 0x10(SP)
  twoDefer.go:12        0x10a2fe7               e8149affff                      CALL fmt.Println(SB)
  twoDefer.go:12        0x10a2fec               c644243700                      MOVB $0x0, 0x37(SP)
  twoDefer.go:12        0x10a2ff1               488b842490000000                MOVQ 0x90(SP), AX
  twoDefer.go:12        0x10a2ff9               488b8c2498000000                MOVQ 0x98(SP), CX
  twoDefer.go:12        0x10a3001               488b9424a0000000                MOVQ 0xa0(SP), DX
  twoDefer.go:12        0x10a3009               48890424                        MOVQ AX, 0(SP)
  twoDefer.go:12        0x10a300d               48894c2408                      MOVQ CX, 0x8(SP)
  twoDefer.go:12        0x10a3012               4889542410                      MOVQ DX, 0x10(SP)
  twoDefer.go:12        0x10a3017               e8e499ffff                      CALL fmt.Println(SB)
  twoDefer.go:12        0x10a301c               488bac24a8000000                MOVQ 0xa8(SP), BP
  twoDefer.go:12        0x10a3024               4881c4b0000000                  ADDQ $0xb0, SP
  twoDefer.go:12        0x10a302b               c3                              RET
  twoDefer.go:12        0x10a302c               e82fdbf8ff                      CALL runtime.deferreturn(SB)
  twoDefer.go:12        0x10a3031               488bac24a8000000                MOVQ 0xa8(SP), BP
  twoDefer.go:12        0x10a3039               4881c4b0000000                  ADDQ $0xb0, SP
  twoDefer.go:12        0x10a3040               c3                              RET
  twoDefer.go:7         0x10a3041               e83af4fbff                      CALL runtime.morestack_noctxt(SB)
  twoDefer.go:7         0x10a3046               e955feffff                      JMP main.main(SB)