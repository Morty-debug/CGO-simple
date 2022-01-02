// Las lineas que contienen #cgo no son comentarios, pero esta linea SI
package main  
/* 
#cgo  CFLAGS:  -I./include 
#cgo  LDFLAGS:  -L./lib  -lsaludo 
#include "saludo.h" 
*/  
import "C"  
func main(){
	nombre := C.CString("Ricky")
	C.saludo(nombre)  
}  
