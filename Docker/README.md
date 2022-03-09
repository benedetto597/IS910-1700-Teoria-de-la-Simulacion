<div>  
<!-- Encabezado -->
## IS910-1700-Teoria de la Simulacion
### I PAC 2022  
### Seccción 1701
### Autor 
| Nombre | Numero De Cuenta | Correo Institucional |
|:-------------:| :-----:|:-----:|
| Edgar Josué Benedetto Godoy | `20171033802` | [UNAH](mailto:edgar.benedetto@unah.hn) |

</div>

#### Qué es Docker
1. Permite ejecutar la app en el mismo entorno 
    * Todas las dependencias del aplicativo estarán dentro del contenedor
2. Sandbox
3. Facil de mover

#### Diferencias entre Docker y maquinas virtuales
La maquina virtual permite ejecutar una aplicación junto con el sistema operativo, la cual corre sobre un hipervisor y una estructura fisica como un servidor.
En Docker se comparte el kernel por lo que un contenedor linux solamente puede correr en un servidor que tenga kernel de linux. 
* ¿Cómo ejecutar contenedores linux en una máquina con windows?
  * Docker Desktop lo que hace es crear un kernel linux en el servidor y ejecutar el contenedor en el servidor con windows.

#### Docker corre contenedores a partir de una imagen
Para correr un contenedor se necesita una imagen que tenga:
    * Sistema Operativo (no el Kernel)
    * Software para ejecutar la app (Apache, PHP, Librerias PHP, etc)
    * Aplicación (Go, Python, Ruby, etc)

¿Cómo crear una imagen?
R/ Se generan usando un DockerFile, el cual es un archivo que indica como crear una imagen, luego se ejecuta el comando **docker-build** para generar la imagen y para correr el contenedor se ejecuta el comando **docker-run**
Para convertir la imagen en contenedor se corre la imagen, la ventaja de esto es que se puede **descargar una imagen de internet** y luego correrla, ahorrando el trabajo de tener que instalar manualmente todas las dependencias de una aplicación.

#### Comandos importantes

1. Correr una imagen:
    ```
    docker run <nombre_imagen>
    ```

    

