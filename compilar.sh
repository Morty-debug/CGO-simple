gcc -fPIC ./include/saludo.c -shared -o ./lib/libsaludo.so
gcc -fPIC ./include/saludo.c -shared -o ./bin/libsaludo.so
export LD_LIBRARY_PATH=$PWD/lib
go build main.go
cp main ./bin/main.bin
rm main
