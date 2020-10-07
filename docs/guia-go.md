# Mini-guía Go

## Nociones básicas

- -> Editor y compilador on-line: [play.golang.org](https://play.golang.org)

- -> Todos los programas de go tienen un `package main` y una función `main()`.  
Ahi es donde estaría el "inicio" del programa.

- -> No necesitas poner `;` por que el compilador los incluye, pero se pueden poner.

- **COMPILAR**: `go build [*.go]`

- **LIBRERÍA "ESTÁNDAR"**: `import fmt`

## Declaración de variables

->*Lenguaje fuertemente tipado*

Estructura declaración variable:

```go
    var [nombre] [tipo]

    //ejemplos:
    var x,y,z int
    var booleano bool
    var nombre string
```

-> todas las variables tienen un valor inicial, No esta a null.

- En caso de números `int`=0
- En caso `string`= cadena vaciá
- En caso de `bool`=`false`

-> NO USAR VARIABLES ES MOTIVO DE ERROR DE COMPILACIÓN.

Forma abreviada de declarar una variable (sin especificar tipo). Tipado dinámico:

```go
func main(){
    x := 23
    nombre := "Manuel"
}
```

### Conversiones de tipos

-> Cuando intentamos concatenar un numero con una cadena, nos da un error de compilación debemos hacer **una conversion de tipos**.

- Usamos paquete `strconv`
- `Itoa`: int -> string
- `Atoi`: string -> int, **DEVUELVE DOS VALORES**.

```go
import(
    "fmt"
    "strconv"
)

edad := 42
edad_str := strconv.Itoa(edad)
fmt.println("mi edad es "+edad_str)
//o fmt.println("mi edad es "+ strconv.Itoa(edad))
```

Atoi:

```go
edad_str:="42"
edad_int,_:=strconv.Atoi(edad)
//Atoi retorna múltiples valores
//Podemos declarar una variable en ese mismo momento PERO DEBEMOS USARLA
//O DESECHAMOS EL VALOR CON _
//En este caso el 2º valor es un código de err
```

## Input y output de datos

- Usa el paquete `fmt`
IMPRIMIR
- `Print`,`Println`, 
- `Printf`(usa formato como en C %)
    - %d int.
    - %v, %s string
    - %t bool
    - %b %e %f float, double etc..


```go
    var name string
    name="Manuel"
    fmt.println(name)
    fmt.println("mi nombre es "+name)//concatenación con +
```

LEER DATOS

- `Scanf`
```go
var edad int
fmt.Scanf(%d\n, &edad)
```
### Usando bufio y os

```go
import(
    "fmt"
    "bufio"
    "os"
)
func main(){
    reader := bufio.NewReader(os.Stdin)
    fmt.Println("Ingresa tu nombre: ")
    nombre,err := reader.ReadString('\n')
    if(err != nil){
        fmtPrintln(err)
    }else{
        fmt.println("Hola "+nombre)
    }
}
```

## Condicionales

- ==, !=, <, >, <=, >=

```go
if {

}else if{

}else
```

- no es obligatorio poner ()
- Siempre se debe poner `{}`
- el `{}` debe de estar en la misma linea que el `if`.

## Bucles

- Solo existe For 
    - `for [initialize];[condición];[ejecución]`
    - Se puede quitar esos elementos

```go
for i:=0;i<10;i++{}

for i<10{}

for{break}
```

- sentencia `continue` comienza de nuevo el bucle.
- sentencia `break`termina el bucle. 

## Arrais, vectores o arreglos

- `var vector [10]int`
- funciona `fmt.Println(vector)`

Formas de declaración:

```go
    vector := [3]int{1,2,3}
    vector2 := [3]int{1,2}

    //tamaño de un vector
    tam := len(vector)  //3

    var matriz [2][2]int
    fmt.Println(matriz)
```

## Slice

Es como un vector dinámico. Estructura basada en array.

- Declaración: `var sli []int`, `var sli []int{1,2,3}`
- Cuando tenemos un slice sin inicializar su valor es `nil`

La estructura del slice tiene 3 datos importantes

- **Puntero al array referenciado**
- **Longitud del array al que apunta**
- **Capacidad**

### Crear un slice con un vector

Se puede crear el slice con un fragmento del vector.

HACER SLICER
```go
arr:=[3]int{1,2,3}

sl1:=arr[:2]    // 1 2
sl2:=arr[1:2]   // 2
sl3:=arr[0:]     //todo el vector 1 2 3

```
### Crear slices con make

- `make([]Tipo, tamaño, capacidad(opcional))`

```go
//make()
slice := make([]int, 3, 5)

len(slice) //3
cap(slice) //5

```

- `append()` añade un elemento

```go
slice = append(slice, 2)
```

### Funcion `copy()`

Copia un array en otro. Copia el minimo de elementos de ambos vectores.

- `copy(destino, fuente)`

```go
    arr:=[]int{1,2,3,4}
    copia:=make([]int, len(arr))

    copy(copia, arr)

    //El tamaño de la copia es el tamaña menor de los dos vectores
    //len = min(copia,fuente)
```

## Punteros

1. una dirección de memoria
2. En lugar del valor, tenemos la dirección en la que esta el valor

DECLARACIÓN:

- `*T` => `*int`, `*string`...
- el valor cero => `nil`
- `&` devuelve la dirección de memoria de una variable
- `*` devuelve el valor de una dirección de memoria

Ejemplo

```go
var x,y *int
num:=5

x = &num
```

## Structs

```go
type User struct{
    edad int
    nombre, apellido string
}

func main(){
    var manu User

    fmt.Println(manu)

    //Asignación indicando nombre de atributo
    manu := User{nombre: "manu", apellido:"HC"}

    fmt.Println(manu)

    //Asignación según orden de declaración de atributos 
    //Se incluyen todos los atributos
    manu := User{20,"",""}

    //new devuelve un puntero
    Manuel:=new(User)
}
```

## Métodos para Structs

- `func (id(this) Struct) nombreFunc() (datoDeveulto){}`

id: identificador del struct en la función
Struct: struct de la función

```go
func (usuario User) nombre_completo() string{
    return usuario.nombre + " " + usuario.apellido
}

func (this *User) set_name(n string) string{
    this.nombre = n
}
```

## Campos anónimos

- Se asemeja mucho a la herencia en lenguajes OO

```go
type Persona struct{
    name string
}

type tutor struct{
    Persona
}

func main(){
    tutor :=Tutor{Humano{"manu"}}
    //HEREDA el nombre de Persona
    fmt.Println(tutor.name)
    //también se puede con:
    fmt.Println(tutor.Persona.name)
}
```

- También podemos usar los métodos y sobrescribir los métodos con `this.campo.func()`

## Interfaces

- Es un tipo de dato
- se define con `type nombre interface{//definiciones de métodos}`