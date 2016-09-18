LIBS = -lkernel32 -luser32 -lgdi32 -lwinspool -lcomdlg32 -ladvapi32 \
-lshell32 -lole32 -loleaut32 -luuid -lodbc32 -lodbccp32

mingwt : main.o
	gcc -o mingwt main.o -L../libyara/Release -L../lib/ -L$(MINGW)/lib -llibyara32 -ljansson32 -llibeay32 $(LIBS)
main.o : main.c
	gcc -c main.c -I../../libyara/include

clean : 
	rm main.o mingwt