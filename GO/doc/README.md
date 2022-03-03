# 

    
<!-- Encabezado -->
## IS910-1700-Teoria de la Simulacion
### I PAC 2022  
### Seccción 1701
### Autor 
| Nombre | Numero De Cuenta | Correo Institucional |
|:-------------:| :-----:|:-----:|
| Edgar Josué Benedetto Godoy | `20171033802` | [UNAH](mailto:edgar.benedetto@unah.hn) |

</div>

_______
_______

### Configuración del entorno de trabajo
Para poder ejecutar código GO desde una carpeta en especifico, esta debe estár en el Path de las variables de entorno del sistema

1. Buscar variables de entorno
2. Agregar en Path de variables del sistema la carpeta "bin" de la carpeta donde se ejecutará el código


### Compilar & Ejecutar en terminal código GO 
1. Ejecutar código:
   
   ```go
   go run <NombrePrograma.go>
   ```
2. Compilar código (generando código binario o .exe en Win):
   
   ```go
   go build <NombrePrograma.go>
   ```
3. Ejecutar el código binario generado al compilar:
   ```go
   ./<NombrePrograma.go>
   ```
