// Las lineas comentadas antes del import "C" es codigo C para el compilador CGO
package main
/*
#include <stdio.h>
void saludo(const char *nombre){
	printf("Hola %s desde C\n", nombre);
}
*/
import "C"  
func main(){
	nombre := C.CString("Ricky")
	C.saludo(nombre)  
}  
