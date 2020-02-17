.PHONY: clean

OBJS = y.tab.o lex.yy.o
COMPILER = plhtml

lex.yy.c: scanner.l
	flex -d -o $@ $^

lex.yy.o: lex.yy.c
	g++ -c -o $@ $<

y.tab.h y.tab.c: parser.y
	bison -dvty parser.y

y.tab.o: y.tab.c
	g++ -c -o y.tab.o y.tab.c

$(COMPILER): $(OBJS)
	g++ -o $@ $(OBJS)

clean:
	rm -f lex.yy.c y.tab.h y.tab.c y.output *.o $(COMPILER)
