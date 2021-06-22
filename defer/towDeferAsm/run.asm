
"".main STEXT size=468 args=0x0 locals=0xc0 funcid=0x0
        0x0000 00000 (/Users/firaga/Gode/modPlayground/defer/twoDefer.go:7)     TEXT    "".main(SB), ABIInternal, $192-0
        0x0000 00000 (/Users/firaga/Gode/modPlayground/defer/twoDefer.go:7)     MOVQ    (TLS), CX
        0x0009 00009 (/Users/firaga/Gode/modPlayground/defer/twoDefer.go:7)     LEAQ    -64(SP), AX
        0x000e 00014 (/Users/firaga/Gode/modPlayground/defer/twoDefer.go:7)     CMPQ    AX, 16(CX)
        0x0012 00018 (/Users/firaga/Gode/modPlayground/defer/twoDefer.go:7)     PCDATA  $0, $-2
        0x0012 00018 (/Users/firaga/Gode/modPlayground/defer/twoDefer.go:7)     JLS     458
        0x0018 00024 (/Users/firaga/Gode/modPlayground/defer/twoDefer.go:7)     PCDATA  $0, $-1
        0x0018 00024 (/Users/firaga/Gode/modPlayground/defer/twoDefer.go:7)     SUBQ    $192, SP
        0x001f 00031 (/Users/firaga/Gode/modPlayground/defer/twoDefer.go:7)     MOVQ    BP, 184(SP)
        0x0027 00039 (/Users/firaga/Gode/modPlayground/defer/twoDefer.go:7)     LEAQ    184(SP), BP
        0x002f 00047 (/Users/firaga/Gode/modPlayground/defer/twoDefer.go:7)     XORPS   X0, X0
        0x0032 00050 (/Users/firaga/Gode/modPlayground/defer/twoDefer.go:7)     MOVUPS  X0, 120(SP)
        0x0037 00055 (/Users/firaga/Gode/modPlayground/defer/twoDefer.go:7)     MOVUPS  X0, 136(SP)
        0x003f 00063 (/Users/firaga/Gode/modPlayground/defer/twoDefer.go:7)     MOVUPS  X0, 152(SP)
        0x0047 00071 (/Users/firaga/Gode/modPlayground/defer/twoDefer.go:7)     MOVUPS  X0, 168(SP)
        0x004f 00079 (/Users/firaga/Gode/modPlayground/defer/twoDefer.go:7)     FUNCDATA        $0, gclocals·69c1753bd5f81501d95132d08af04464(SB)
        0x004f 00079 (/Users/firaga/Gode/modPlayground/defer/twoDefer.go:7)     FUNCDATA        $1, gclocals·1a3f980b14e0b4569583e99692361894(SB)
        0x004f 00079 (/Users/firaga/Gode/modPlayground/defer/twoDefer.go:7)     FUNCDATA        $2, "".main.stkobj(SB)
        0x004f 00079 (/Users/firaga/Gode/modPlayground/defer/twoDefer.go:7)     FUNCDATA        $4, "".main.opendefer(SB)
        0x004f 00079 (/Users/firaga/Gode/modPlayground/defer/twoDefer.go:7)     MOVB    $0, ""..autotmp_20+71(SP)
        0x0054 00084 (/Users/firaga/Gode/modPlayground/defer/twoDefer.go:8)     XORPS   X0, X0
        0x0057 00087 (/Users/firaga/Gode/modPlayground/defer/twoDefer.go:8)     MOVUPS  X0, ""..autotmp_11+104(SP)
        0x005c 00092 (/Users/firaga/Gode/modPlayground/defer/twoDefer.go:8)     LEAQ    type.string(SB), AX
        0x0063 00099 (/Users/firaga/Gode/modPlayground/defer/twoDefer.go:8)     MOVQ    AX, ""..autotmp_11+104(SP)
        0x0068 00104 (/Users/firaga/Gode/modPlayground/defer/twoDefer.go:8)     LEAQ    ""..stmp_0(SB), CX
        0x006f 00111 (/Users/firaga/Gode/modPlayground/defer/twoDefer.go:8)     MOVQ    CX, ""..autotmp_11+112(SP)
        0x0074 00116 (/Users/firaga/Gode/modPlayground/defer/twoDefer.go:8)     LEAQ    fmt.Println·f(SB), CX
        0x007b 00123 (/Users/firaga/Gode/modPlayground/defer/twoDefer.go:8)     MOVQ    CX, ""..autotmp_21+128(SP)
        0x0083 00131 (/Users/firaga/Gode/modPlayground/defer/twoDefer.go:8)     LEAQ    ""..autotmp_11+104(SP), DX
        0x0088 00136 (/Users/firaga/Gode/modPlayground/defer/twoDefer.go:8)     MOVQ    DX, ""..autotmp_22+160(SP)
        0x0090 00144 (/Users/firaga/Gode/modPlayground/defer/twoDefer.go:8)     MOVQ    $1, ""..autotmp_22+168(SP)
        0x009c 00156 (/Users/firaga/Gode/modPlayground/defer/twoDefer.go:8)     MOVQ    $1, ""..autotmp_22+176(SP)
        0x00a8 00168 (/Users/firaga/Gode/modPlayground/defer/twoDefer.go:8)     MOVB    $1, ""..autotmp_20+71(SP)
        0x00ad 00173 (/Users/firaga/Gode/modPlayground/defer/twoDefer.go:9)     MOVUPS  X0, ""..autotmp_15+88(SP)
        0x00b2 00178 (/Users/firaga/Gode/modPlayground/defer/twoDefer.go:9)     MOVQ    AX, ""..autotmp_15+88(SP)
        0x00b7 00183 (/Users/firaga/Gode/modPlayground/defer/twoDefer.go:9)     LEAQ    ""..stmp_1(SB), DX
        0x00be 00190 (/Users/firaga/Gode/modPlayground/defer/twoDefer.go:9)     MOVQ    DX, ""..autotmp_15+96(SP)
        0x00c3 00195 (/Users/firaga/Gode/modPlayground/defer/twoDefer.go:9)     MOVQ    CX, ""..autotmp_23+120(SP)
        0x00c8 00200 (/Users/firaga/Gode/modPlayground/defer/twoDefer.go:9)     LEAQ    ""..autotmp_15+88(SP), CX
        0x00cd 00205 (/Users/firaga/Gode/modPlayground/defer/twoDefer.go:9)     MOVQ    CX, ""..autotmp_24+136(SP)
        0x00d5 00213 (/Users/firaga/Gode/modPlayground/defer/twoDefer.go:9)     MOVQ    $1, ""..autotmp_24+144(SP)
        0x00e1 00225 (/Users/firaga/Gode/modPlayground/defer/twoDefer.go:9)     MOVQ    $1, ""..autotmp_24+152(SP)
        0x00ed 00237 (/Users/firaga/Gode/modPlayground/defer/twoDefer.go:9)     MOVB    $3, ""..autotmp_20+71(SP)
        0x00f2 00242 (/Users/firaga/Gode/modPlayground/defer/twoDefer.go:11)    MOVUPS  X0, ""..autotmp_19+72(SP)
        0x00f7 00247 (/Users/firaga/Gode/modPlayground/defer/twoDefer.go:11)    MOVQ    AX, ""..autotmp_19+72(SP)
        0x00fc 00252 (/Users/firaga/Gode/modPlayground/defer/twoDefer.go:11)    LEAQ    ""..stmp_2(SB), AX
        0x0103 00259 (/Users/firaga/Gode/modPlayground/defer/twoDefer.go:11)    MOVQ    AX, ""..autotmp_19+80(SP)
        0x0108 00264 (<unknown line number>)    NOP
        0x0108 00264 ($GOROOT/src/fmt/print.go:274)     MOVQ    os.Stdout(SB), AX
        0x010f 00271 ($GOROOT/src/fmt/print.go:274)     LEAQ    go.itab.*os.File,io.Writer(SB), CX
        0x0116 00278 ($GOROOT/src/fmt/print.go:274)     MOVQ    CX, (SP)
        0x011a 00282 ($GOROOT/src/fmt/print.go:274)     MOVQ    AX, 8(SP)
        0x011f 00287 ($GOROOT/src/fmt/print.go:274)     LEAQ    ""..autotmp_19+72(SP), AX
        0x0124 00292 ($GOROOT/src/fmt/print.go:274)     MOVQ    AX, 16(SP)
        0x0129 00297 ($GOROOT/src/fmt/print.go:274)     MOVQ    $1, 24(SP)
        0x0132 00306 ($GOROOT/src/fmt/print.go:274)     MOVQ    $1, 32(SP)
        0x013b 00315 ($GOROOT/src/fmt/print.go:274)     PCDATA  $1, $1
        0x013b 00315 ($GOROOT/src/fmt/print.go:274)     NOP
        0x0140 00320 ($GOROOT/src/fmt/print.go:274)     CALL    fmt.Fprintln(SB)
        0x0145 00325 (/Users/firaga/Gode/modPlayground/defer/twoDefer.go:12)    MOVB    $1, ""..autotmp_20+71(SP)
        0x014a 00330 (/Users/firaga/Gode/modPlayground/defer/twoDefer.go:12)    MOVQ    ""..autotmp_24+136(SP), AX
        0x0152 00338 (/Users/firaga/Gode/modPlayground/defer/twoDefer.go:12)    MOVQ    ""..autotmp_24+144(SP), CX
        0x015a 00346 (/Users/firaga/Gode/modPlayground/defer/twoDefer.go:12)    MOVQ    ""..autotmp_24+152(SP), DX
        0x0162 00354 (/Users/firaga/Gode/modPlayground/defer/twoDefer.go:12)    MOVQ    AX, (SP)
        0x0166 00358 (/Users/firaga/Gode/modPlayground/defer/twoDefer.go:12)    MOVQ    CX, 8(SP)
        0x016b 00363 (/Users/firaga/Gode/modPlayground/defer/twoDefer.go:12)    MOVQ    DX, 16(SP)
        0x0170 00368 (/Users/firaga/Gode/modPlayground/defer/twoDefer.go:12)    CALL    fmt.Println(SB)
        0x0175 00373 (/Users/firaga/Gode/modPlayground/defer/twoDefer.go:12)    MOVB    $0, ""..autotmp_20+71(SP)
        0x017a 00378 (/Users/firaga/Gode/modPlayground/defer/twoDefer.go:12)    MOVQ    ""..autotmp_22+160(SP), AX
        0x0182 00386 (/Users/firaga/Gode/modPlayground/defer/twoDefer.go:12)    MOVQ    ""..autotmp_22+168(SP), CX
        0x018a 00394 (/Users/firaga/Gode/modPlayground/defer/twoDefer.go:12)    MOVQ    ""..autotmp_22+176(SP), DX
        0x0192 00402 (/Users/firaga/Gode/modPlayground/defer/twoDefer.go:12)    MOVQ    AX, (SP)
        0x0196 00406 (/Users/firaga/Gode/modPlayground/defer/twoDefer.go:12)    MOVQ    CX, 8(SP)
        0x019b 00411 (/Users/firaga/Gode/modPlayground/defer/twoDefer.go:12)    MOVQ    DX, 16(SP)
        0x01a0 00416 (/Users/firaga/Gode/modPlayground/defer/twoDefer.go:12)    CALL    fmt.Println(SB)
        0x01a5 00421 (/Users/firaga/Gode/modPlayground/defer/twoDefer.go:12)    MOVQ    184(SP), BP
        0x01ad 00429 (/Users/firaga/Gode/modPlayground/defer/twoDefer.go:12)    ADDQ    $192, SP
        0x01b4 00436 (/Users/firaga/Gode/modPlayground/defer/twoDefer.go:12)    RET
        0x01b5 00437 (/Users/firaga/Gode/modPlayground/defer/twoDefer.go:12)    CALL    runtime.deferreturn(SB)
        0x01ba 00442 (/Users/firaga/Gode/modPlayground/defer/twoDefer.go:12)    MOVQ    184(SP), BP
        0x01c2 00450 (/Users/firaga/Gode/modPlayground/defer/twoDefer.go:12)    ADDQ    $192, SP
        0x01c9 00457 (/Users/firaga/Gode/modPlayground/defer/twoDefer.go:12)    RET
        0x01ca 00458 (/Users/firaga/Gode/modPlayground/defer/twoDefer.go:12)    NOP
        0x01ca 00458 (/Users/firaga/Gode/modPlayground/defer/twoDefer.go:7)     PCDATA  $1, $-1
        0x01ca 00458 (/Users/firaga/Gode/modPlayground/defer/twoDefer.go:7)     PCDATA  $0, $-2
        0x01ca 00458 (/Users/firaga/Gode/modPlayground/defer/twoDefer.go:7)     CALL    runtime.morestack_noctxt(SB)
        0x01cf 00463 (/Users/firaga/Gode/modPlayground/defer/twoDefer.go:7)     PCDATA  $0, $-1
        0x01cf 00463 (/Users/firaga/Gode/modPlayground/defer/twoDefer.go:7)     JMP     0