## Funciones como Tipos de Primera Clase en Go

En Go, las funciones son **tipos de primera clase**, lo que significa que pueden:

- Asignarse a variables
- Pasarse como argumentos a otras funciones
- Retornarse desde funciones
- Almacenarse en estructuras o slices

Como

```go
type Handler func(string)
```
Aquí Handler es un tipo función, lo que permite tratar funciones como cualquier otro valor.

**Funciones de Orden Superior**

Un middleware tiene la forma:
```go
func(next Handler) Handler
```
Esto significa que:

- Recibe una función <code>(next)</code>
- Retorna otra función
- Encapsula comportamiento adicional antes o después de ejecutar <code>next</code>

Ejemplo simplificado:
```go
func Logging(next Handler) Handler {
    return func(request string) {
        // lógica antes
        next(request)
        // lógica después
    }
}
```
Gracias a que Go permite que las funciones sean valores, podemos construir cadenas dinámicas como:
```go
handler := Chan(finalHandler, Logging, Auth)
```
equivalente a:
```go
Logging(Auth(finalHandler))
```
