package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/streadway/amqp"
)

// type TimeRange struct {
// 	Id            int
// 	StartInterval time.Time
// 	EndInterval   time.Time
// 	Frecuency     float64
// }

type Station struct {
	Id       int
	Name     string
	Patients Patient
	Resource Resource
	Status   bool
	Taken    bool
}

type Resource struct {
	Id     int
	Name   string
	Status bool
}

type Patient struct {
	Id            int
	Name          string
	ArriveTime    int
	AttentionTime int
	Attended      bool
}

type Parameters struct {
	Qtydays      int
	Qtystations  int
	Qtyresources int
	Stations     []Station
	Resources    []Resource
}

type Result struct {
	Day                           int
	TotalPatients                 int
	AttendedPatients              int
	UnattendedPatients            int
	AttendedPatientsTimeTotal     int
	UnattendedPatientsTimeTotal   int
	AttendedPatientsTimeAverage   float64
	UnattendedPatientsTimeAverage float64
}

func RemoveInPos(a []int, i int) []int {
	a[i] = a[len(a)-1]
	return a[:len(a)-1]
}

func main() {

	//Crear Parametros
	Parameters := CreateParameters()
	var Results []Result
	// Moverse por días simulados
	for i := 0; i < Parameters.Qtydays; i++ {

		var RandomFloat float64
		Pat := 1
		// Resultados de eficiencia en la atención
		var AttendedPatients []Patient
		var UnattendedPatients []Patient
		var AttendedPatientsTimeTotal int
		var UnattendedPatientsTimeTotal int
		var AttendedPatientsTimeAverage float64
		var UnattendedPatientsTimeAverage float64

		// Resultados de eficiencia en gestión de recursos y estaciones
		var NewPatient Patient

		// Moversee por las horas laborales
		// Los turnos son de 6 horas o 360 minutos por día laboral (930 minutos)
		for mins := 1; mins <= 930; mins++ {

			fmt.Println("\nMinuto: ", mins)
			//Crear Colas En RabiitMQ
			for _, param := range Parameters.Stations {
				CreateQueue(param.Name)
			}
			// Inicio de Turno o cambios de turno
			if mins == 1 || mins == 360 || mins == 720 {
				// Asignar recursos a estaciones
				Parameters = AssignResourceStation(Parameters)
			} else if mins == 930 {
				//Reiniciar Parametros
				parameters := ResetParameters(Parameters)

				// Eliminar las colas En RabiitMQ
				for _, param := range parameters.Stations {
					DeleteQueue(param.Name)
				}

				AttendedPatientsTimeAverage = float64(AttendedPatientsTimeTotal) / float64(len(AttendedPatients))
				UnattendedPatientsTimeAverage = float64(UnattendedPatientsTimeTotal) / float64(len(UnattendedPatients))

				Result := Result{
					Day:                           i + 1,
					TotalPatients:                 len(AttendedPatients) + len(UnattendedPatients),
					AttendedPatients:              len(AttendedPatients),
					UnattendedPatients:            len(UnattendedPatients),
					AttendedPatientsTimeTotal:     AttendedPatientsTimeTotal,
					UnattendedPatientsTimeTotal:   UnattendedPatientsTimeTotal,
					AttendedPatientsTimeAverage:   AttendedPatientsTimeAverage,
					UnattendedPatientsTimeAverage: UnattendedPatientsTimeAverage,
				}

				fmt.Println("\n")
				fmt.Println("--------------------------------------------------------------------------------------")
				fmt.Println("Resultados del día: ", i+1)
				fmt.Println("Total de pacientes: ", Result.TotalPatients)
				fmt.Println("Pacientes atendidos: ", Result.AttendedPatients)
				fmt.Println("Pacientes no atendidos: ", Result.UnattendedPatients)
				fmt.Println("Tiempo promedio de atención de pacientes atendidos: ", Result.AttendedPatientsTimeAverage)
				fmt.Println("Tiempo promedio de atención de pacientes no atendidos: ", Result.UnattendedPatientsTimeAverage)
				fmt.Println("Tiempo total de atención de pacientes atendidos: ", Result.AttendedPatientsTimeTotal)
				fmt.Println("Tiempo total de atención de pacientes no atendidos: ", Result.UnattendedPatientsTimeTotal)
				fmt.Println("--------------------------------------------------------------------------------------")
				fmt.Println("\n")

				Results = append(Results, Result)
				AttendedPatientsTimeTotal = 0
				UnattendedPatientsTimeTotal = 0
				AttendedPatientsTimeAverage = 0
				UnattendedPatientsTimeAverage = 0
				AttendedPatients = nil
				UnattendedPatients = nil
				Pat = 1

			} else if mins <= 180 {
				rand.Seed(int64(time.Now().UnixNano()))
				RandomFloat = rand.Float64()
				if RandomFloat > 0.31 {
					NewPatient = CreatePatient(Pat, mins)
					Pat++
				}
			} else if mins > 180 && mins <= 360 {
				rand.Seed(int64(time.Now().UnixNano()))
				RandomFloat = rand.Float64()
				if RandomFloat > 0.46 {
					NewPatient = CreatePatient(Pat, mins)
					Pat++
				}
			} else if mins > 360 && mins <= 450 {
				rand.Seed(int64(time.Now().UnixNano()))
				RandomFloat = rand.Float64()
				if RandomFloat > 0.55 {
					NewPatient = CreatePatient(Pat, mins)
					Pat++
				}
			} else if mins >= 540 && mins <= 840 {
				rand.Seed(int64(time.Now().UnixNano()))
				RandomFloat = rand.Float64()
				if RandomFloat > 0.73 {
					NewPatient = CreatePatient(Pat, mins)
					Pat++
				}
			} else if mins > 840 && mins < 930 {
				rand.Seed(int64(time.Now().UnixNano()))
				RandomFloat = rand.Float64()
				if RandomFloat > 0.88 {
					NewPatient = CreatePatient(Pat, mins)
					Pat++
				}
			}
			//-------------------------------------------------------------------------------
			IndexActiveStations := []int{}
			for index, param := range Parameters.Stations {
				if param.Taken == false && param.Status == true {
					IndexActiveStations = append(IndexActiveStations, index)
				}
			}

			if NewPatient.Name != "" && len(IndexActiveStations) > 0 {
				rand.Seed(int64(time.Now().UnixNano()))
				RandomStation := rand.Intn(len(IndexActiveStations))
				SelectedStation := Parameters.Stations[IndexActiveStations[RandomStation]]
				PatientAttended := false
				for index, param := range Parameters.Stations {
					if SelectedStation.Id != param.Id && Parameters.Stations[index].Patients.Id == NewPatient.Id {
						PatientAttended = true
					}
				}
				if SelectedStation.Taken == false && PatientAttended == false {
					TotalMins := NewPatient.ArriveTime + NewPatient.AttentionTime
					if TotalMins < 360 || (TotalMins > 370 && TotalMins < 450) || (TotalMins > 540 && TotalMins < 720) || (TotalMins > 730 && TotalMins < 930) {
						Parameters = AssignPatient(NewPatient, Parameters, SelectedStation.Id-1)
					} else {
						UnattendedPatients = append(UnattendedPatients, NewPatient)
						UnattendedPatientsTimeTotal += NewPatient.AttentionTime
					}
				}

			} else if NewPatient.Name != "" && len(IndexActiveStations) == 0 {
				UnattendedPatients = append(UnattendedPatients, NewPatient)
				UnattendedPatientsTimeTotal += NewPatient.AttentionTime
			}

			//-------------------------------------------------------------------------------
			// Verificar si hay pacientes en la estación
			for index, param := range Parameters.Stations {
				// Verificar si el paciente ya está atendido
				if param.Patients.ArriveTime+param.Patients.AttentionTime == mins {
					ExitTime := Parameters.Stations[index].Patients.ArriveTime + Parameters.Stations[index].Patients.AttentionTime
					fmt.Printf("\nSaliendo de la %s el %s a los %d minutos", Parameters.Stations[index].Name, Parameters.Stations[index].Patients.Name, ExitTime)
					AttendedPatients = append(AttendedPatients, Parameters.Stations[index].Patients)
					AttendedPatientsTimeTotal += Parameters.Stations[index].Patients.AttentionTime
					Parameters.Stations[index].Taken = false
					Consumer(param.Name)
					fmt.Printf("\n[x] %s saliendo de la cola\n", Parameters.Stations[index].Patients.Name)
					Parameters.Stations[index].Patients = Patient{}
				}
			}
			//time.Sleep(time.Duration(mins) * time.Millisecond)
		}
	}
}

func AssignPatient(patient Patient, parameters Parameters, IndexActiveStation int) Parameters {
	// Asignar paciente a la estación que le corresponde

	parameters.Stations[IndexActiveStation].Patients.Id = patient.Id
	parameters.Stations[IndexActiveStation].Patients.Name = patient.Name
	parameters.Stations[IndexActiveStation].Patients.ArriveTime = patient.ArriveTime
	parameters.Stations[IndexActiveStation].Patients.AttentionTime = patient.AttentionTime
	parameters.Stations[IndexActiveStation].Patients.Attended = true
	parameters.Stations[IndexActiveStation].Taken = true
	fmt.Printf("\nEntrando a la %s el %s a los %d minutos\n", parameters.Stations[IndexActiveStation].Name, parameters.Stations[IndexActiveStation].Patients.Name, parameters.Stations[IndexActiveStation].Patients.ArriveTime)
	Producer(parameters.Stations[IndexActiveStation].Name, parameters.Stations[IndexActiveStation].Patients.Name)

	return parameters
}

func CreatePatient(id int, mins int) Patient {
	// Crear paciente con un tiempo de atención entre 5 y 10 mins
	rand.Seed(int64(time.Now().UnixNano()))
	max := 10
	min := 5
	radomAttention := rand.Intn(max-min) + min
	patient := Patient{
		Id:            id,
		Name:          "Paciente " + strconv.Itoa(id),
		ArriveTime:    mins,
		AttentionTime: radomAttention,
		Attended:      false,
	}
	return patient
}

func AssignResourceStation(parameters Parameters) Parameters {
	// Validar que estaciones están libres para usarse en el siguiente turnos
	AvaibleStations := []int{}
	for i := 0; i < len(parameters.Stations); i++ {
		if parameters.Stations[i].Status == false {
			AvaibleStations = append(AvaibleStations, i)
		} else {
			// Cambiar el estado de las estaciones a libre
			parameters.Stations[i].Status = false
		}

	}

	// Validar que recursos están libres para usarse en el siguiente turnos
	AvaibleResources := []int{}

	for i := 0; i < len(parameters.Resources); i++ {
		if parameters.Resources[i].Status == false {

			AvaibleResources = append(AvaibleResources, i)
		} else {
			// Cambiar el estado de las estaciones a libre
			parameters.Resources[i].Status = false
		}
	}

	// Seleccionar los recursos y estaciones a asignar en el siguiente turno
	fmt.Println("\nEstaciones Disponibles: ", AvaibleStations)
	fmt.Println("Recursos Disponibles: ", AvaibleResources)
	AssignedS := ToAssign(AvaibleStations)
	AssignedR := ToAssign(AvaibleResources)
	//AcceptedLen := false

	/*for AcceptedLen == false {
		if len(AssignedR) <= len(AssignedS) && len(AssignedR) > 0 {
			AcceptedLen = true
		} else {
			AssignedR = ToAssign(AvaibleResources)
		}
	}*/

	fmt.Printf("\nPosibles estaciones a asignar: %d", AssignedS)
	fmt.Printf("\nPosibles recursos a asignar: %d\n", AssignedR)

	//Asignar Recursos a las estaciones
	rand.Seed(int64(time.Now().UnixNano()))
	for i := 0; i < len(AssignedR); i++ {

		var randomIndexS int
		var randomIndexR int

		// Asignar el ultimo recurso o la ultima estación
		if len(AssignedS) == 1 {
			randomIndexS = 0
			i = len(AssignedR)
		} else {
			randomIndexS = rand.Intn(len(AssignedS))
		}

		if len(AssignedR) == 1 {
			randomIndexR = 0
			i = len(AssignedR)
		} else {
			randomIndexR = rand.Intn(len(AssignedR))
		}

		randomS := AssignedS[randomIndexS]
		randomR := AssignedR[randomIndexR]

		// fmt.Printf("\nRandom Index Resource: %d\n", randomR)
		// fmt.Printf("\nObject Resource: %d\n", parameters.Resources[randomR].Id)
		if parameters.Stations[randomS].Status == false {
			parameters.Stations[randomS].Resource.Id = parameters.Resources[randomR].Id
			parameters.Stations[randomS].Resource.Name = parameters.Resources[randomR].Name
			parameters.Stations[randomS].Resource.Status = true
			//Cambio de Estado Recursos & Estaciones
			parameters.Resources[randomR].Status = true
			parameters.Stations[randomS].Status = true

			fmt.Printf("\n%s asignado a la %s", parameters.Stations[randomS].Resource.Name, parameters.Stations[randomS].Name)
		}
		AssignedR = RemoveInPos(AssignedR, randomIndexR)
		AssignedS = RemoveInPos(AssignedS, randomIndexS)

	}

	return parameters
}

func CreateParameters() Parameters {
	reader := bufio.NewReader(os.Stdin)
	var parameters Parameters
	// Pedir al usuario los parametros de entrada de la simulacion
	// Cantidad de días a simular
	fmt.Println("\nIngrese la cantidad de días a simular: ")
	entry1, _ := reader.ReadString('\n')
	strDays := strings.TrimRight(entry1, "\r\n") // Remove \n and \r
	if strDays == "" {
		fmt.Println("Debe ingresar la cantidad de días a simular")
		os.Exit(1)
	}
	Qtydays, _ := strconv.Atoi(strDays)

	// Cantidad de estaciones --> MAXIMO 15 estaciones
	fmt.Println("\nIngrese la cantidad de estaciones (máximo 15): ")
	entry2, _ := reader.ReadString('\n')
	strStations := strings.TrimRight(entry2, "\r\n") // Remove \n and \r
	if strStations == "" {
		fmt.Println("Debe ingresar la cantidad de estaciones para la simulacion")
		os.Exit(1)
	}
	Qtystations, _ := strconv.Atoi(strStations)

	if Qtystations > 15 {
		fmt.Println("La cantidad de estaciones no puede ser mayor a 15")
		os.Exit(1)
	}

	// Cantidad de recursos
	fmt.Println("\nIngrese la cantidad de recursos: ")
	entry3, _ := reader.ReadString('\n')
	strResources := strings.TrimRight(entry3, "\r\n") // Remove \n and \r
	if strResources == "" {
		fmt.Println("Debe ingresar la cantidad de recursos para la simulacion")
		os.Exit(1)
	}
	Qtyresources, _ := strconv.Atoi(strResources)

	// Crear las estaciones
	var stations []Station
	for i := 1; i <= Qtystations; i++ {
		station := Station{
			Id:       i,
			Name:     "Estacion " + strconv.Itoa(i),
			Patients: Patient{},
			Resource: Resource{},
			Taken:    false,
			Status:   false,
		}
		stations = append(stations, station)
	}

	// Crear los recursos
	var resources []Resource
	for i := 1; i <= Qtyresources; i++ {
		resource := Resource{
			Id:     i,
			Name:   "Recurso " + strconv.Itoa(i),
			Status: false,
		}
		resources = append(resources, resource)
	}

	parameters = Parameters{
		Qtydays:      Qtydays,
		Qtystations:  Qtystations,
		Qtyresources: Qtyresources,
		Stations:     stations,
		Resources:    resources,
	}
	return parameters
}

func ResetParameters(parameters Parameters) Parameters {
	// Crear las estaciones
	var stations []Station
	for i := 1; i <= parameters.Qtystations; i++ {
		station := Station{
			Id:       i,
			Name:     "Estacion " + strconv.Itoa(i),
			Patients: Patient{},
			Resource: Resource{},
			Status:   false,
		}
		stations = append(stations, station)
	}

	// Crear los recursos
	var resources []Resource
	for i := 1; i <= parameters.Qtyresources; i++ {
		resource := Resource{
			Id:     i,
			Name:   "Recurso " + strconv.Itoa(i),
			Status: false,
		}
		resources = append(resources, resource)
	}

	parameters = Parameters{
		Qtydays:      parameters.Qtydays,
		Qtystations:  parameters.Qtystations,
		Qtyresources: parameters.Qtyresources,
		Stations:     stations,
		Resources:    resources,
	}
	return parameters
}

func ToAssign(Array []int) []int {
	var Assigned []int

	rand.Seed(int64(time.Now().UnixNano()))
	for i := 0; i < len(Array)-1; i++ {
		Found := false
		RandomL := rand.Intn(len(Array))
		RandomN := Array[RandomL]
		i := sort.SearchInts(Assigned, RandomN)
		if i < len(Assigned) && Assigned[i] == RandomN {
			Found = true
		}
		if Found == false {
			Assigned = append(Assigned, RandomN)
		}
	}
	return Assigned
}

func CreateQueue(Name string) amqp.Queue {
	connection, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatalf("%s: %s", "Error de conexion", err)
	}
	defer connection.Close()

	ch, err := connection.Channel()
	if err != nil {
		log.Fatalf("%s: %s", "Error de conexion dentro del canal", err)
	}
	defer ch.Close()

	queue, err := ch.QueueDeclare(
		Name,  // name
		false, // durable
		false, // elimine automaticamente al usarla
		false, // exclusiva
		false, // timeout
		nil,
	)

	if err != nil {
		log.Fatalf("%s: %s", "Error cuando creamos la queue", err)
	}

	err = ch.Publish(
		"",
		queue.Name,
		false,
		false,

		amqp.Publishing{
			ContentType: "text/plain",
		})

	if err != nil {
		log.Fatalf("%s: %s", "Error cuando enviamos el mensaje", err)
	}

	_, err = ch.Consume(
		queue.Name, // name of the queue
		"",         // consumer
		true,       // auto-ack
		false,      // exclusive
		false,      // no-local
		false,      // no-wait
		nil,        // args
	)

	return queue
}

func DeleteQueue(Name string) {
	connection, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatalf("%s: %s", "Error de conexion", err)
	}
	defer connection.Close()

	ch, err := connection.Channel()
	if err != nil {
		log.Fatalf("%s: %s", "Error de conexion dentro del canal", err)
	}
	defer ch.Close()

	queue, err := ch.QueueDeclare(
		Name,  // name
		false, // durable
		false, // elimine automaticamente al usarla
		false, // exclusiva
		false, // timeout
		nil,
	)

	if err != nil {
		log.Fatalf("%s: %s", "Error cuando creamos la queue", err)
	}

	_, err = ch.QueueDelete(
		queue.Name,
		false,
		false,
		false,
	)

	if err != nil {
		log.Fatalf("%s: %s", "Error cuando eliminamos la queue", err)
	}

}

func Producer(Name string, msg string) {
	connection, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatalf("%s: %s", "Error de conexion", err)
	}
	defer connection.Close()

	canal, err := connection.Channel()
	if err != nil {
		log.Fatalf("%s: %s", "Error de conexion dentro del canal", err)
	}
	defer canal.Close()

	queue, err := canal.QueueDeclare(
		Name,  // name
		false, // durable
		false, // elimine automaticamente al usarla
		false, // exclusiva
		false, // timeout
		nil,
	)

	if err != nil {
		log.Fatalf("%s: %s", "Error cuando creamos la queue", err)
	}

	err = canal.Publish(
		"",
		queue.Name,
		false,
		false,

		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(msg),
		})

	fmt.Printf("[x] %s agregado a la cola\n", msg)

	if err != nil {
		log.Fatalf("%s: %s", "Error cuando enviamos el mensaje", err)
	}

}

func Consumer(Name string) {
	connection, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatalf("%s: %s", "Error connecting to RabbitMQ", err)
	}
	defer connection.Close()

	// failOnError(err, "Failed to connect to RabbitMQ")

	channel, err := connection.Channel()
	if err != nil {
		log.Fatalf("%s: %s", "Error connecting to the channel", err)
	}
	defer channel.Close()

	queue, err := channel.QueueDeclare(
		Name,  // name of the queue
		false, // durable or persistent
		false, // delete when is used
		false, // exclusive
		false, // timeout
		nil,   // arguments
	)

	if err != nil {
		log.Fatalf("%s: %s", "Error creating the queue", err)
	}

	message, err := channel.Consume(
		queue.Name, // name of the queue
		"",         // consumer
		true,       // auto-ack
		false,      // exclusive
		false,      // no-local
		false,      // no-wait
		nil,        // args
	)

	if err != nil {
		log.Fatalf("%s: %s", "Error creating the consume channel", err)
	}

	// Anonymous function to recive messages
	go func() {
		for queue := range message {
			fmt.Printf("\n%s: %s", "Message received", string(queue.Body))
		}
	}()
}
