include test/Makefile

all:demo dev2 dev3 dev4 dev5 dev6 dev7 dev8 dev9 dev10 dev10.1 dev11 for switch if const
build:
	go build
n:switch
jock:
	find .|grep -v git|grep -v .so|grep -v code-cache|grep -v ast|grep -v parse|grep -v token|grep -v scanner|xargs cat|wc -l
