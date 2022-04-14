# 
<div align="center">
    <img src="https://miro.medium.com/max/1400/1*_WUT6QJF_vnniSERwRje1w.png" width="300px"> </img> 
    
<!-- Encabezado -->
### IS910-1700 Teoria de la Simulacion
### Proyecto 
#### I PAC 2022  
#### Seccción 1701
#### Catedratico 

| Nombre | Correo |
|:-------------:| :-----:|
|  Claudio José Paz Fonseca | [Gmail](mailto:cjosepfonseca@gmail.com) |


### Integrantes 
| Nombre | Numero De Cuenta |
|:-------------:| :-----:|
| Edgar Josué Benedetto Godoy | `20171033802` |
| Tiffany Monique Matamoros Gonzalez | `20181002925` |
| William Alberto Gómez López | 20161900396 |
| Yohelis Lindeth Ordoñez Alvarez | 20161002245 |
| Jose Luis Maradiaga Zambrano | 20091011842 |

</div>

______

### Algoritmo
#### Datos de entrada

1. El número de estaciones con las que se desarrollará la simulación, teniendo como máximo 15 estaciones, debido al espacio físico del lugar.
2. El rango de tiempo de duración de la simulación en términos de días.
3. El numero de recursos que serán asignados en las estaciones.

#### Generación de datos aleatorios
1. Las llegadas de los pacientes serán generadas de manera aleatoria dentro del rango de tiempo especificado
   * La frecuencia de llegada por minuto esta especificada por la siguiente tabla, dependiendo al intervalo de tiempo:

    | Intervalo de tiempo | Frecuencia por minuto |  
    |:-----------------:|:-----------------------:|
    | 4:30 AM a 7:30 AM | 0.31 | 
    | 7:31 AM a 10:30 AM | 0.46 | 
    | 10.31 AM a 12:00 M | 0.55 | 
    | 12:00 AM a 1:30 PM | 0.00 | 
    | 1:31 PM a 6:30 PM | 0.73 | 
    | 6:30 PM a 8:00 PM | 0.88 | 

2. La llegada de los pacientes serán enviados de manera aleatoria a cada una de las estaciones.
3. La política de asignación del recurso que aplica la vacuna en cada estación es aleatoria en función al numero de recursos suministrado al algoritmo
   * Cualquier recurso pudiese estar asignado a la estación dependiendo a la disponibilidad y este a su vez rotar el turno cada 6 horas.

#### Consideraciones
* El tiempo de duración de la atención se distribuye de manera uniforme entre 5 y 10 minutos, tiempo en el cual la estación de atención queda inhabilitada para atender a otro paciente.

#### Ejemplo
* Se recibe como parámetro 5 estaciones y 8 recursos, 5 de estos estarán asignados a las estaciones y 3 de ellos en espera a ser asignados al efectuarse el cambio de turno.
    * Nota: en este ejemplo, al efectuarse el cambio de turno 3 estaciones estarían disponibles y las otras 2 estarían inhabilitadas, hasta el próximo cambio de turno. 

