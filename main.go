// Las lineas que contienen # no son comentarios, pero esta linea SI
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
