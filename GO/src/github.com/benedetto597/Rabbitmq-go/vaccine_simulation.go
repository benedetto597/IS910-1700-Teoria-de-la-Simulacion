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
}

type Parameters struct {
	Qtydays      int
	Qtystations  int
	Qtyresources int
	Stations     []Station
	Resources    []Resource
}

func RemoveInPos(a []int, i int) []int {
	a[i] = a[len(a)-1]
	return a[:len(a)-1]
}

func main() {

	fmt.Println("*****************************************************************************************************")
	fmt.Println("Proyecto de Implementación y evaluación de centro de vacunación en la ciudad de Tegucigalpa, Honduras")
	fmt.Println("*****************************************************************************************************")

	//Crear Parametros
	parameters := CreateParameters()

	// Moverse por días simulados
	for i := 0; i < parameters.Qtydays; i++ {
		fmt.Println(".....................................................................................................")
		fmt.Printf("                                           Dia No. %d \n", i+1)
		fmt.Println(".....................................................................................................")

		var RandomFloat float64
		var Pat int
		// Resultados de eficiencia en la atención
		var NonAttendedPatients []Patient
		var AttendedPatients []Patient
		var OneInterval int
		var TwoInterval int
		var ThreeInterval int
		var FiveInterval int
		var SixInterval int
		var AttendedPatientsTimeTotal int
		var NonAttendedPatientsTimeTotal int

		// Resultados de eficiencia en gestión de recursos y estaciones
		var NewPatient Patient

		// Moversee por las horas laborales
		// Los turnos son de 6 horas o 360 minutos por día laboral (930 minutos)
		for mins := 1; mins <= 930; mins++ {

			//Crear Colas En RabiitMQ
			for _, param := range parameters.Stations {
				CreateQueue(param.Name)
			}
			// Inicio de Turno o cambios de turno
			if mins == 1 || mins == 360 || mins == 720 {
				// Asignar recursos a estaciones

				fmt.Println("-----------------------------------------------------------------------------------------------------")
				fmt.Println("                                     Asignación de Recursos                                          ")
				fmt.Println("-----------------------------------------------------------------------------------------------------")
				parameters = AssignResourceStation(parameters)
			} else if mins == 930 {
				//Reiniciar Parametros
				parameters := ResetParameters(parameters)

				// Eliminar las colas En RabiitMQ
				for _, param := range parameters.Stations {
					DeleteQueue(param.Name)
				}
			} else if mins <= 180 {
				rand.Seed(int64(time.Now().UnixNano()))
				RandomFloat = rand.Float64()
				if RandomFloat > 0.31 {
					Pat++
					OneInterval++
					NewPatient = CreatePatient(Pat, mins)
					NonAttendedPatients = append(NonAttendedPatients, NewPatient)

				}

			} else if mins > 180 && mins <= 360 {
				rand.Seed(int64(time.Now().UnixNano()))
				RandomFloat = rand.Float64()
				if RandomFloat > 0.46 {
					Pat++
					TwoInterval++
					NewPatient = CreatePatient(Pat, mins)
					NonAttendedPatients = append(NonAttendedPatients, NewPatient)

				}
			} else if mins > 360 && mins <= 450 {
				rand.Seed(int64(time.Now().UnixNano()))
				RandomFloat = rand.Float64()
				if RandomFloat > 0.55 {
					Pat++
					ThreeInterval++
					NewPatient = CreatePatient(Pat, mins)
					NonAttendedPatients = append(NonAttendedPatients, NewPatient)
				}
			} else if mins > 540 && mins <= 840 {
				rand.Seed(int64(time.Now().UnixNano()))
				RandomFloat = rand.Float64()
				if RandomFloat > 0.73 {
					Pat++
					FiveInterval++
					NewPatient = CreatePatient(Pat, mins)
					NonAttendedPatients = append(NonAttendedPatients, NewPatient)

				}
			} else if mins > 840 && mins < 930 {
				rand.Seed(int64(time.Now().UnixNano()))
				RandomFloat = rand.Float64()
				if RandomFloat > 0.88 {
					Pat++
					SixInterval++
					NewPatient = CreatePatient(Pat, mins)
					NonAttendedPatients = append(NonAttendedPatients, NewPatient)

				}
			}

			//time.Sleep(time.Duration(mins) * time.Millisecond)

			//-------------------------------------------------------------------------------

			IndexActiveStations := []int{}
			for index, param := range parameters.Stations {
				if param.Status == true && param.Taken == false {
					IndexActiveStations = append(IndexActiveStations, index)
				}
			}

			if len(NonAttendedPatients) == 0 {

			}

			if mins <= 450 || mins > 540 {

				if len(IndexActiveStations) > 0 && len(NonAttendedPatients) > 0 {
					parameters = AssignPatient(mins, NonAttendedPatients[0], parameters, IndexActiveStations)
					NonAttendedPatients = RemovePatient(NonAttendedPatients, 0)
					AttendedPatientsTimeTotal++

				} else {

					NonAttendedPatientsTimeTotal += NewPatient.AttentionTime
				}

			}
			//-------------------------------------------------------------------------------

			// Verificar si hay pacientes en la estación
			for index, param := range parameters.Stations {
				// Verificar si el paciente ya está atendido
				if param.Patients.ArriveTime+param.Patients.AttentionTime == mins {

					AttendedPatients = append(AttendedPatients, param.Patients)
					parameters.Stations[index].Taken = false
					parameters.Stations[index].Patients = Patient{}
					Consumer(param.Name)

				}
			}

			if mins == 930 {
				fmt.Println("-----------------------------------------------------------------------------------------------------")
				fmt.Println("                                         Informe Final del Día ")
				fmt.Println("-----------------------------------------------------------------------------------------------------")

				fmt.Println("\n                                      Pacientes Generados  Por Intervalos")
				fmt.Println("_____________________________________________________________________________________________________")

				fmt.Printf("Pacientes Generados de 4:30 AM a 7:30 AM: %d \n", OneInterval)
				fmt.Printf("Pacientes Generados de 7:31 AM a 10:30 AM: %d \n", TwoInterval)
				fmt.Printf("Pacientes Generados de 10.31 AM a 12:00 M: %d \n", ThreeInterval)
				fmt.Printf("Pacientes Generados de 12:00 AM a 1:30 PM: 0 \n")
				fmt.Printf("Pacientes Generados de 1:31 PM a 6:30 PM: %d \n", FiveInterval)
				fmt.Printf("Pacientes Generados de 6:30 PM a 8:00 PM : %d \n", SixInterval)

				fmt.Println("\n                                      Datos Generales")
				fmt.Println("_____________________________________________________________________________________________________")
				fmt.Printf("Número de Estaciones: %d \n", parameters.Qtystations)
				fmt.Printf("Número de Recurso: %d \n", parameters.Qtyresources)
				fmt.Printf("Total de Pacientes que llegaron: %d \n", Pat)
				fmt.Printf("Cantidad de Pacientes Atendidos: %d \n", AttendedPatientsTimeTotal)
				fmt.Printf("Cantidad de Pacientes no Atendidos: %d \n\n", len(NonAttendedPatients))

			}

		}
	}
}

func AssignPatient(mins int, patient Patient, parameters Parameters, IndexActiveStations []int) Parameters {
	// Asignar paciente a la estación que le corresponde
	rand.Seed(int64(time.Now().UnixNano()))
	RandomStation := rand.Intn(len(IndexActiveStations))

	if parameters.Stations[IndexActiveStations[RandomStation]].Taken == false {

		parameters.Stations[IndexActiveStations[RandomStation]].Patients = patient
		parameters.Stations[IndexActiveStations[RandomStation]].Patients.ArriveTime = mins
		parameters.Stations[IndexActiveStations[RandomStation]].Taken = true
		Producer(parameters.Stations[IndexActiveStations[RandomStation]].Name, parameters.Stations[IndexActiveStations[RandomStation]].Patients.Name)

	}
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
		ArriveTime:    0,
		AttentionTime: radomAttention,
	}

	return patient
}

func AssignResourceStation(parameters Parameters) Parameters {
	// Validar que estaciones están libres para usarse en el siguiente turnos
	AvaibleStations := []int{}
	for i := 0; i < len(parameters.Stations); i++ {
		//	if parameters.Stations[i].Status == false {
		AvaibleStations = append(AvaibleStations, i)
		//	} else {
		// Cambiar el estado de las estaciones a libre
		//	parameters.Stations[i].Status = false
		//}

	}

	// Validar que recursos están libres para usarse en el siguiente turnos
	AvaibleResources := []int{}

	for i := 0; i < len(parameters.Resources); i++ {
		if len(parameters.Resources) < len(parameters.Stations) {
			parameters.Resources[i].Status = false

		}

		if parameters.Resources[i].Status == false {

			AvaibleResources = append(AvaibleResources, i)
		} else {
			// Cambiar el estado de las estaciones a libre
			parameters.Resources[i].Status = false
		}
	}

	// Seleccionar los recursos y estaciones a asignar en el siguiente turno

	AssignedS := ToAssign(AvaibleStations)
	AssignedR := ToAssign(AvaibleResources)
	//AcceptedLen := false

	//Asignar Recursos A las estaciones
	rand.Seed(int64(time.Now().UnixNano()))
	//centinel:=len(AssignedR)
	i := false

	for _, param := range parameters.Resources {
		for i == false {

			if param.Status == false {
				var randomIndexS int
				var randomIndexR int

				if len(AssignedS) == 1 {
					randomIndexS = 0
					i = true
				} else {
					randomIndexS = rand.Intn(len(AssignedS))
				}

				if len(AssignedR) == 1 {
					randomIndexR = 0
					i = true

				} else {
					randomIndexR = rand.Intn(len(AssignedR))
				}

				randomS := AssignedS[randomIndexS]
				randomR := AssignedR[randomIndexR]

				parameters.Stations[randomS].Resource.Id = parameters.Resources[randomR].Id
				parameters.Stations[randomS].Resource.Name = parameters.Resources[randomR].Name
				parameters.Stations[randomS].Resource.Status = true
				//Cambio de Estado Recursos & Estaciones
				parameters.Resources[randomR].Status = true
				parameters.Stations[randomS].Status = true

				AssignedR = RemoveInPos(AssignedR, randomIndexR)
				AssignedS = RemoveInPos(AssignedS, randomIndexS)

				fmt.Printf("Estación  %d : %s \n", parameters.Stations[randomS].Id, parameters.Stations[randomS].Resource.Name)

			}
		}
	}

	return parameters
}

func CreateParameters() Parameters {
	reader := bufio.NewReader(os.Stdin)
	var parameters Parameters
	// Pedir al usuario los parametros de entrada de la simulacion
	// Cantidad de días a simular
	fmt.Println("Ingrese la cantidad de días a simular: ")
	entry1, _ := reader.ReadString('\n')
	strDays := strings.TrimRight(entry1, "\r\n") // Remove \n and \r
	Qtydays, _ := strconv.Atoi(strDays)

	// Cantidad de estaciones --> MAXIMO 15 estaciones
	fmt.Println("Ingrese la cantidad de estaciones (máximo 15): ")
	entry2, _ := reader.ReadString('\n')
	strStations := strings.TrimRight(entry2, "\r\n") // Remove \n and \r
	Qtystations, _ := strconv.Atoi(strStations)

	if Qtystations > 15 {
		fmt.Println("La cantidad de estaciones no puede ser mayor a 15")
		os.Exit(1)
	}

	// Cantidad de recursos
	fmt.Println("Ingrese la cantidad de recursos: ")
	entry3, _ := reader.ReadString('\n')
	strResources := strings.TrimRight(entry3, "\r\n") // Remove \n and \r
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
			Name:   "Vacuna " + strconv.Itoa(i),
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
			Name:   "Vacuna " + strconv.Itoa(i),
			Status: true,
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
	j := 0
	i := 0
	RandomN := 0
	var temp []int

	for j < 1 {

		Found := false

		RandomL := rand.Intn(len(Array))
		RandomN = Array[RandomL]
		sort.Ints(temp)
		i = sort.SearchInts(temp, RandomN)

		if len(Assigned) > 0 && i < len(temp) {

			if temp[i] == RandomN || temp[0] == RandomN {
				Found = true

			}
			if Found == false {
				Assigned = append(Assigned, RandomN)
				temp = append(temp, RandomN)

			}

		} else {
			Assigned = append(Assigned, RandomN)
			temp = append(temp, RandomN)

		}
		if len(Assigned) == len(Array) {
			j++
		}
	}

	return Assigned
}

func RemovePatient(s []Patient, index int) []Patient {
	return append(s[:index], s[index+1:]...)
}

// func remove(s []int, index int) []int {
// 	return append(s[:index], s[index+1:]...)
// }

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

/*func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}

}*/

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
	/*	var num int
		num, err = canal.QueueDelete(
			queue.Name,
			false,
			false,
			false,
		)
		log.Println("[RABBITMQ_CLIENT]", num, "message purged")*/

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
	//log.Printf(" [x] Sent %s", msg)

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

	_, mensaje := channel.Consume(
		queue.Name, // name of the queue
		"",         // consumer
		true,       // auto-ack
		false,      // exclusive
		false,      // no-local
		false,      // no-wait
		nil,        // args
	)

	if err != nil {
		log.Fatalf("%s: %s", "Error creating the consume channel", mensaje)
	}

	// Anonymous function to recive messages
	/*go func() {
		for queue := range message {
			fmt.Printf("\n%s: %s", "Message received", string(queue.Body))
		}
	}()*/
}
