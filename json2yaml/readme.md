json yaml 相互转换

[Go语言encoding/json库源码分析](https://zhuanlan.zhihu.com/p/37165706)

[json google document](https://pkg.go.dev/encoding/json#example-Decoder)

marshal 递归方式encode生成json
unmarshal 两层分析,底层词法分析器, 上层并不是语法分析,利用下层生成的opcode最终生成结果


todo:
* customMarshalJson
* textMarshalJson