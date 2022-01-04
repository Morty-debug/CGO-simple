gcc -fPIC ./include/saludo.c -shared -o ./lib/libsaludo.dll
gcc -fPIC ./include/saludo.c -shared -o ./bin/libsaludo.dll
go build main.go
copy main.exe bin\main.exe
del /f main.exe
