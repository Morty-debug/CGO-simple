gcc -fPIC ./include/saludo.c -shared -o ./lib/libsaludo.so
export LD_LIBRARY_PATH=$PWD/lib
go build main.go
