## 操作步骤

```
❯❯❯ cat -n 1.go
     1	package main
     2
     3	type person struct {
     4		age int
     5	}
     6
     7	func main() {
     8		var a = &person{111}
     9		println(a)
    10	}
```

我们要看第 8 行编译后变成啥了：

```
❯❯❯ go tool compile -S 1.go | grep "1.go:8"
	0x001d 00029 (1.go:8)	PCDATA	$0, $0
	0x001d 00029 (1.go:8)	PCDATA	$1, $0
	0x001d 00029 (1.go:8)	MOVQ	$0, ""..autotmp_2+8(SP)
	0x0026 00038 (1.go:8)	MOVQ	$111, ""..autotmp_2+8(SP)
```

两行的版本：

```
 ❯❯❯ cat -n 2.go
     1	package main
     2
     3	type person struct {
     4		age int
     5	}
     6
     7	func main() {
     8		var b = person{111}
     9		var a = &b
    10		println(a)
    11	}
    12
```

我们要看第 8 和第 9 行：

```
❯❯❯ go tool compile -S 2.go | grep -E '(2.go:8|2.go:9)'
	0x001d 00029 (2.go:8)	PCDATA	$0, $0
	0x001d 00029 (2.go:8)	PCDATA	$1, $0
	0x001d 00029 (2.go:8)	MOVQ	$0, "".b+8(SP)
	0x0026 00038 (2.go:8)	MOVQ	$111, "".b+8(SP)
```

可以看到，这里的一行版本的代码和两行版本的代码最终编译出的结果是完全一致的，没有任何区别。

