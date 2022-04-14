package main

import "time"

type TimeRange struct {
	startInterval time.Time
	endInterval   time.Time
	frecuency     float64
}

type Station struct {
	id       int
	name     string
	capacity int
	patients []Patient
	Resources []Resourse
	status   bool
}

type Resourse struct {
	id     int
	status bool
}

type Patient struct {
	id         int
	name 	   string
	arriveTime time.Time
	attentionTime  float64
	TimeRange
	Station
}

// SET
func (patient Patient) setPatientId(patients []Patient) int {
	patient.id = len(patients)
	return patient.id
}

func (patient Patient) setStationId(patients []Patient) int {
	patient.id = len(patients)
	return patient.id
}

func (resourse Resourse) setResourseId(resourses []Resourse) int {
	resourse.id = len(resourses)
	return resourse.id
}

func (patient Patient) setTimeRange(tr []TimeRange) {
	for _, patient := range tr {
		patient.TimeRange = 
		
}

func (station, Station) setArrivePatient(patients []Patient) {

}

func main()	{
	// Crear rangos de tiempo 
	// Pedir al usuario los parametros de entrada de la simulacion
	// Crear las estaciones 
	// Creamos los recursos
	
	// Hacer un for para moverse entre los dias a simular
		// Crear un for para moverse entre las horas a simular (Cada 6 horas hay cambio de turno)
			// Asignar los recursos a las estaciones
			// Cambiar el estado de recursos
			// Crear los numeros aleatorios entre 5 y 10 (float) que su suma de 6 horas o 360 mins
				// Crear un for para recorrer los el arreglo de numeros aleatorios entre 5 y 10
				// Crear al paciente (crear aleatoriamente el arriveTime)
					// Validar en que rango de tiempo está y en que iteracion de las 4 posibles del día está para determinar por ejemplo si está en la segunda iteración no y queda en el rango de tiempo de 1:31 PM a 6:30 PM no poder asignarle arriveTime mayor a las 2:30
					// Validar que el arriveTime del paciente no choque con los demas pacientes
				// Asignarle el tiempo de atencion al paciente
				// Asignar el paciente a la estacion aleatoriamente 
			// Cambiar el estado de estaciones asignadas
}