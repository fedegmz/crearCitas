# Programador de citas

este proyecto fue creado para realizar citas para ser implemenatado en lugar donde se requiera dicha funciopnalidad. Se implemento ciertas validaciones.
- el usuario no puede agregar una cita cuando el dia y la hora sea menor al actual
- no se permite empalmar las citas ya que por ende esta ya esta ocupada
- podemos eliminar dicha cita
- podemos actualizar (drag and drop)

## Estructura del Proyecto

- **frontend**: Contiene el código del lado del cliente desarrollado en JavaScript, html
- **backend**: Contiene el código del lado del servidor desarrollado en Go.
- **docs**: Documentación adicional.

## Configuración

1. **Clonar repositorio**

```bash
git clone https://github.com/fedegmz/crearCitas
```
2. **Base de datos:** Deberá crear una base de datos llamada `cita` y copia la información que se encuentra en el archivo `schema.db` para crear la tabla.

3. **navegar al proyecto y ejecutar el servidor:** navegando en el archivo en el carpeta `api` ejecutamos el siguiente comando para poner en marcha nuestro servidor.

``` bash
go run .
```

4. Con nuestro servidor funcionado ya podríamos abrir nuestro archivo index.html 
