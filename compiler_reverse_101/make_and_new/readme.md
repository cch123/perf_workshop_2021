## 观察 make

```go
❯❯❯ cat -n make.go
     1	package main
     2
     3	func main() {
     4		// make slice
     5		// 空间开的比较大，是为了让这个 slice 分配在堆上，栈上的 slice 结果不太一样
     6		var sl = make([]int, 100000)
     7		println(sl)
     8
     9		// make channel
    10		var ch = make(chan int, 5)
    11		println(ch)
    12
    13		// make map
    14		var m = make(map[int]int, 22)
    15		println(m)
    16	}
```

```shell
go build make.go && go tool objdump ./make | grep -E "make.go:6|make.go:10|make.go:14"
```

这里使用 go tool compile -S 也是可以的。

## 观察 new(**输出内容难以读懂，不推荐**)

```
❯❯❯ cat -n new.go
     1	package main
     2
     3	type person struct{ age int }
     4
     5	func main() {
     6		var a = new(int)
     7		var b = new(person)
     8		var c = new(chan int)
     9		var d = new(map[int]int)
    10
    11		println(a, b, c, d)
    12	}
```

```
❯❯❯ go build -gcflags="-N -l" new.go && go tool objdump new | grep -E "new.go:6|new.go:7|new.go:8|new.go:9"
  new.go:6		0x1051591		48c744241000000000	MOVQ $0x0, 0x10(SP)
  new.go:6		0x105159a		488d442410		LEAQ 0x10(SP), AX
  new.go:6		0x105159f		4889442430		MOVQ AX, 0x30(SP)
  new.go:7		0x10515a4		48c744240800000000	MOVQ $0x0, 0x8(SP)
  new.go:7		0x10515ad		488d442408		LEAQ 0x8(SP), AX
  new.go:7		0x10515b2		4889442428		MOVQ AX, 0x28(SP)
  new.go:8		0x10515b7		48c744244000000000	MOVQ $0x0, 0x40(SP)
  new.go:8		0x10515c0		488d442440		LEAQ 0x40(SP), AX
  new.go:8		0x10515c5		4889442420		MOVQ AX, 0x20(SP)
  new.go:9		0x10515ca		48c744243800000000	MOVQ $0x0, 0x38(SP)
  new.go:9		0x10515d3		488d442438		LEAQ 0x38(SP), AX
  new.go:9		0x10515d8		4889442418		MOVQ AX, 0x18(SP)
```

被优化搞得面目全非了。

```
❯❯❯ go tool compile -N -S  new.go | grep -E "new.go:6|new.go:7|new.go:8|new.go:9"
	0x0021 00033 (new.go:6)	MOVQ	$0, ""..autotmp_4+16(SP)
	0x002a 00042 (new.go:6)	LEAQ	""..autotmp_4+16(SP), AX
	0x002f 00047 (new.go:6)	MOVQ	AX, "".a+48(SP)
	0x0034 00052 (new.go:7)	MOVQ	$0, ""..autotmp_5+8(SP)
	0x003d 00061 (new.go:7)	LEAQ	""..autotmp_5+8(SP), AX
	0x0042 00066 (new.go:7)	MOVQ	AX, "".b+40(SP)
	0x0047 00071 (new.go:8)	MOVQ	$0, ""..autotmp_6+64(SP)
	0x0050 00080 (new.go:8)	LEAQ	""..autotmp_6+64(SP), AX
	0x0055 00085 (new.go:8)	MOVQ	AX, "".c+32(SP)
	0x005a 00090 (new.go:9)	MOVQ	$0, ""..autotmp_7+56(SP)
	0x0063 00099 (new.go:9)	LEAQ	""..autotmp_7+56(SP), AX
	0x0068 00104 (new.go:9)	MOVQ	AX, "".d+24(SP)
```

可见就是生成了一些临时变量。
