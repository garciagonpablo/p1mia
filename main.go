package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("┌──────────────────────┐")
	fmt.Println("│  GESTOR DE ARCHIVOS  │")
	fmt.Println("├──────────────────────┤")
	fmt.Println("│     Pablo García     │")
	fmt.Println("│      201901107       │")
	fmt.Println("└──────────────────────┘")

	for {
		readInput()
	}
}

func readInput() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Ingrese comando:")
	comando, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("No se reconoció el comando: ", err)
		return
	}
	comando = strings.TrimSpace(comando)
	analize(comando)
}

func analize(comando string) {
	//mkdisk -size=3000 -unit=K login
	//Mkdisk
	//MKDISK
	//login
	//#inicio mkdisk del proyecto1

	//var comando ----> comando := "Hola este es un comando" ----> cambio en memoria
	//Apuntador al cambiar la variable, la cambia para todo el programa
	comandoSeparado := strings.Split(comando, " ")
	if strings.Contains(comandoSeparado[0], "#") {
		//Imprimir el comentario
		fmt.Println("Comentario: ")
		//Eliminar el # del comentario
		comandoSeparado[0] = strings.Replace(comandoSeparado[0], "#", "", -1)
		for _, comentario := range comandoSeparado {
			fmt.Println(comentario + " ")
		}
	} else {
		//Si no es un comentario, entonces es un comando
		//Iterar sobre el comando separado
		for _, valor := range comandoSeparado {
			//el primer valor del comando lo pasamos a minusculas
			valor = strings.ToLower(valor)
			//Si el valor es igual a mkdisk, entonces es un comando de creacion de disco
			if valor == "mkdisk" {
				fmt.Println("Ejecutando comando mkdisk")
				//analize Comando Mkdisk
				analizeMkdisk(&comandoSeparado)
				//Pasar a string el comando separado
				comandoSeparadoString := strings.Join(comandoSeparado, " ")
				analize(comandoSeparadoString)
			} else if valor == "rmdisk" {
				fmt.Println("Ejecutando comando rmdisk")
				//analize Comando Rmdisk
				analizeRmdisk(&comandoSeparado)
				//Pasar a string el comando separado
				comandoSeparadoString := strings.Join(comandoSeparado, " ")
				analize(comandoSeparadoString)
			} else if valor == "fdisk" {
				fmt.Println("Ejecutando comando fdisk")
				//analize Comando Fdisk
				analizeFdisk(&comandoSeparado)
				//Pasar a string el comando separado
				comandoSeparadoString := strings.Join(comandoSeparado, " ")
				analize(comandoSeparadoString)
			} else if valor == "rep" {
				fmt.Println("Ejecutando comando rep")
				// fs.ReporteDisk(&comandoSeparado)
			} else if valor == "\n" {
				continue
			} else if valor == "\r" {
				continue
			} else if valor == "" {
				continue
			} else if valor[0] == '-'{
				continue
			}else if valor == "mount" {
				fmt.Println("Ejecutando mount")
			}else if valor == "unmount" {
				fmt.Println("Ejecutando unmount")
			}else if valor == "mkfs" {
				fmt.Println("Ejecutando mkfs")
			}else if valor == "login" {
				fmt.Println("Ejecutando login")
			}else if valor == "logout" {
				fmt.Println("Ejecutando logout")
			}else if valor == "mkgrp" {
				fmt.Println("Ejectuando mkgrp")
			}else if valor == "rmgrp" {
				fmt.Println("Ejecutando rmgrp")
			}else if valor == "mkusr" {
				fmt.Println("Ejecutando mkusr")
			}else if valor == "rmusr" {
				fmt.Println("Ejecutando rmusr")
			}else if valor == "mkfile" {
				fmt.Println("Ejecutando mkfile")
			}else if valor == "cat" {
				fmt.Println("Ejecutando cat")
			}else if valor == "remove" {
				fmt.Println("Ejecutando remove")
			}else if valor == "edit" {
				fmt.Println("Ejecutando edit")
			}else if valor == "rename" {
				fmt.Println("Ejecutando rename")
			}else if valor == "mkdir" {
				fmt.Println("Ejecutando mkdir")
			}else if valor == "copy" {
				fmt.Println("Ejecutando copy")
			}else if valor == "move" {
				fmt.Println("Ejecutando move")
			}else if valor == "find" {
				fmt.Println("Ejecutando find")
			}else if valor == "chown" {
				fmt.Println("Ejecutando chown")
			}else if valor == "chgrp" {
				fmt.Println("Ejecutando chgrp")
			}else if valor == "chmod" {
				fmt.Println("Ejecutando chmod")
			}else if valor == "pause" {
				fmt.Println("Ejecutando pause")
			}else if valor == "loss" {
				fmt.Println("Ejecutando loss")
			}else if valor == "execute" {
				fmt.Println("Ejecutando execute")
			} else {
				fmt.Println("Comando No reconocido")
				fmt.Println(valor)
			}
		}
	}
}

func analizeMkdisk(comandoSeparado *[]string) {
	//mkdisk -size=3000 -unit=K -fit
	*comandoSeparado = (*comandoSeparado)[1:]
	//Iterar sobre el comando separado
	var size, fit, unit bool
	//Variables para almacenar los valores de los parametros
	var sizeValor, fitValor, unitValor string
	fitValor = "f"
	unitValor = "m"
	//Iterar sobre el comando separado
	for _, valor := range *comandoSeparado {
		bandera := ObtenerBandera(valor)
		banderaValor := ObtenerBanderaValor(valor)		
		if bandera == "-size" {
			size = true
			sizeValor = banderaValor
			*comandoSeparado = (*comandoSeparado)[1:]
		} else if bandera == "-fit" {
			fit = true
			fitValor = banderaValor
			fitValor = strings.ToLower(fitValor)
			*comandoSeparado = (*comandoSeparado)[1:]
		} else if bandera == "-unit" {
			unit = true
			unitValor = banderaValor
			unitValor = strings.ToLower(unitValor)
			*comandoSeparado = (*comandoSeparado)[1:]
		} else {
			fmt.Println("Parametro no reconocido: ", bandera)
		}
	}

	//Verificar si se ingresaron los parametros obligatorios
	if !size {
		fmt.Println("El parametro -size es obligatorio")
		return
	} else {
		if fit {
			if fitValor != "bf" && fitValor != "ff" && fitValor != "wf" {
				fmt.Println("El valor del parametro -fit no es valido")
				return
			} else {
				if fitValor == "bf" {
					fitValor = "b"
				} else if fitValor == "ff" {
					fitValor = "f"
				} else if fitValor == "wf" {
					fitValor = "w"
				}
			}
		}
		if unit {
			if unitValor != "k" && unitValor != "m" {
				fmt.Println("El valor del parametro -unit no es valido")
				return
			}
		}
		//Pasar a entero el valor del size
		fmt.Println(sizeValor)
		sizeInt, err := strconv.Atoi(sizeValor)
		if err != nil {
			fmt.Println("El valor del parametro -size no es valido")
			return
		}
		if sizeInt <= 0 {
			fmt.Println("El valor del parametro -size no es valido")
			return
		}

		//Imprimir los valores de los parametros
		fmt.Println("Size: ", sizeValor)
		fmt.Println("Fit: ", fitValor)
		fmt.Println("Unit: ", unitValor)
		//Llamar a la funcion para crear el disco
		//CrearDisco(sizeInt, fitValor, unitValor)
		// fs.CrearDisco(sizeInt, fitValor, unitValor)
	}

}

func analizeRmdisk(comandoSeparado *[]string) {
	//rmdisk -driveletter=A
	*comandoSeparado = (*comandoSeparado)[1:]
	//Iterar sobre el comando separado
	var driveletter string
	var drive bool
	//Iterar sobre el comando separado
	for _, valor := range *comandoSeparado {
		bandera := ObtenerBandera(valor)
		banderaValor := ObtenerBanderaValor(valor)
		if bandera == "-driveletter" {
			driveletter = banderaValor
			driveletter = strings.ToUpper(driveletter)
			drive = true
			*comandoSeparado = (*comandoSeparado)[1:]
		} else {
			fmt.Println("Parametro no reconocido: ", bandera)
		}
	}
	//Verificar si se ingresaron los parametros obligatorios
	if !drive {
		fmt.Println("El parametro -driveletter es obligatorio")
		return
	} else {
		//Imprimir los valores de los parametros
		fmt.Println("Driveletter: ", driveletter)
		//Llamar a la funcion para eliminar el disco
		//Buscar el disco con la letra en el directorio Discos
		//EliminarDisco(driveletter)
		//os.Remove("Discos/" + driveletter + ".dsk")
	}
}

func analizeFdisk(comandoSeparado *[]string) {
	//fdisk -size=300 -driveletter=A -name=Particion1
	*comandoSeparado = (*comandoSeparado)[1:]
	//Booleanos para verificar si se ingresaron los parametros
	var banderaLetter, banderaName, banderaFit, banderaUnit, banderaType, banderaDelete, banderaAdd bool
	//Variables para almacenar los valores de los parametros
	var sizeValor, letterValor, nameValor, fitValor, unitValor, typeValor, deleteValor, addValor string
	//Setear valores por defecto
	fitValor = "w"
	unitValor = "k"
	typeValor = "p"
	deleteValor = "0"
	addValor = "0"
	sizeValor = "0"
	//Iterar sobre el comando separado
	for _, valor := range *comandoSeparado {
		//-size
		bandera := ObtenerBandera(valor)
		//300
		banderaValor := ObtenerBanderaValor(valor)
		if bandera == "-size" {
			sizeValor = banderaValor
			*comandoSeparado = (*comandoSeparado)[1:]
		} else if bandera == "-driveletter" {
			banderaLetter = true
			letterValor = banderaValor
			letterValor = strings.ToUpper(letterValor)
			*comandoSeparado = (*comandoSeparado)[1:]
		} else if bandera == "-name" {
			banderaName = true
			nameValor = banderaValor
			*comandoSeparado = (*comandoSeparado)[1:]
		} else if bandera == "-fit" {
			banderaFit = true
			fitValor = banderaValor
			fitValor = strings.ToLower(fitValor)
			*comandoSeparado = (*comandoSeparado)[1:]
		} else if bandera == "-unit" {
			banderaUnit = true
			unitValor = banderaValor
			unitValor = strings.ToLower(unitValor)
			*comandoSeparado = (*comandoSeparado)[1:]
		} else if bandera == "-type" {
			banderaType = true
			typeValor = banderaValor
			typeValor = strings.ToLower(typeValor)
			*comandoSeparado = (*comandoSeparado)[1:]
		} else if bandera == "-delete" {
			banderaDelete = true
			deleteValor = banderaValor
			*comandoSeparado = (*comandoSeparado)[1:]
		} else if bandera == "-add" {
			banderaAdd = true
			addValor = banderaValor
			*comandoSeparado = (*comandoSeparado)[1:]
		} else {
			fmt.Println("Parametro no reconocido: ", bandera)
		}
	}
	//Obligatorios: -size(al crear), driveletter, name
	//Verificar si se ingresaron los parametros obligatorios
	if !banderaLetter {
		fmt.Println("El parametro -driveletter es obligatorio")
		return
	} else if !banderaName {
		fmt.Println("El parametro -name es obligatorio")
		return
	} else {
		//Pasar a entero el valor del size
		sizeInt, err := strconv.Atoi(sizeValor)
		if err != nil {
			fmt.Println("El valor del parametro -size no es valido")
			return
		}
		if sizeInt <= 0 {
			fmt.Println("El valor del parametro -size no es valido")
			return
		}

		if banderaFit {
			if fitValor != "bf" && fitValor != "ff" && fitValor != "wf" {
				fmt.Println("El valor del parametro -fit no es valido")
				return
			} else {
				if fitValor == "bf" {
					fitValor = "b"
				} else if fitValor == "ff" {
					fitValor = "f"
				} else if fitValor == "wf" {
					fitValor = "w"
				}
			}
		}
		if !banderaUnit {
			unitValor = "k"
		} else {
			if unitValor != "k" && unitValor != "m" && unitValor != "b" {
				fmt.Println("El valor del parametro -unit no es valido")
				return
			}
		}
		if !banderaType {
			typeValor = "p"
		} else {
			if typeValor != "p" && typeValor != "e" && typeValor != "l" {
				fmt.Println("El valor del parametro -type no es valido")
				return
			}
		}
		if banderaDelete {
			if deleteValor != "full" {
				fmt.Println("El valor del parametro -delete no es valido")
				return

			}
		}
		var addInt int
		if banderaAdd {
			//Intentar pasar a entero el valor del size a entero
			addInt, err := strconv.Atoi(addValor)
			if err != nil {
				fmt.Println("El valor del parametro -add no es valido")
				return
			}
			if addInt != 0 {
				fmt.Println("El valor del parametro -add no es valido")
				return
			}
		}
		//Imprimir los valores de los parametros
		fmt.Println("Size: ", sizeInt)
		fmt.Println("Driveletter: ", letterValor)
		fmt.Println("Name: ", nameValor)
		fmt.Println("Fit: ", fitValor)
		fmt.Println("Unit: ", unitValor)
		fmt.Println("Type: ", typeValor)
		fmt.Println("Delete: ", deleteValor)
		fmt.Println("Add: ", addInt)
		//Llamar a la funcion para crear la particion
		// Filesystem.Fdisk(sizeInt, letterValor, nameValor, fitValor, unitValor, typeValor, deleteValor, addInt)
	}
}

func ObtenerBandera(bandera string) string {
	//mkdisk -size=3000 -unit=K
	var banderaValor string
	for _, valor := range bandera {
		if valor == '=' {
			break
		}
		banderaValor += string(valor)
	}
	banderaValor = strings.ToLower(banderaValor)
	return banderaValor
}

func ObtenerBanderaValor(bandera string) string {
	//mkdisk -size=3000 -unit=K
	var banderaValor string
	var banderaEncontrada bool
	for _, valor := range bandera {
		if banderaEncontrada {
			banderaValor += string(valor)
		}
		if valor == '=' {
			banderaEncontrada = true
		}
	}
	return banderaValor
}