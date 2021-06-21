翻墙困难的话，

git clone https://github.com/golang/perf.git

进 benchstat 的目录，go install

go test -run=NONE -bench=Garbage -benchmem -count=5 > old.txt

change code

go test -run=NONE -bench=Garbage -benchmem -count=5 > new.txt

then benchstat old.txt new.txt

