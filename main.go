package main

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"

	"github.com/go-chi/chi/v5"
)

func main() {
	// Crear un enrutador y configurar el server
	router := chi.NewRouter()
	server := &http.Server{
		Addr:    ":3000",
		Handler: router,
	}
	router.Post("/execute-go-function", MyFunction)

	FileServer(router)
	fmt.Println("Running http://localhost:3000")
	panic(server.ListenAndServe())

}

func FileServer(router *chi.Mux) {
	//servir archivos estáticos
	root := "./TestM/dist"
	fs := http.FileServer(http.Dir(root))

	router.Get("/*", func(w http.ResponseWriter, r *http.Request) {
		if _, err := os.Stat(root + r.RequestURI); os.IsNotExist(err) {
			http.StripPrefix(r.RequestURI, fs).ServeHTTP(w, r)
		} else {
			fs.ServeHTTP(w, r)
		}
	})
}

type Person struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
	Email string `json:"email"`
}

// Hacer la solicitud GET
func MyFunction(w http.ResponseWriter, r *http.Request) {
	req, err := http.NewRequest("GET", "https://8j5baasof2.execute-api.us-west-2.amazonaws.com/production/tests/trucode/samples?size=5", nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Content-Type", "text/csv")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	log.Println(resp.StatusCode)

	// Parse CSV
	reader := csv.NewReader(resp.Body)
	var people []Person

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		person := Person{
			Name:  formatname(record[0], record[1]),
			Phone: formatPhone(record[1]),
			Email: record[2],
		}
		people = append(people, person)
	}

	postURL := "https://8j5baasof2.execute-api.us-west-2.amazonaws.com/production/tests/trucode/items"

	// Iterar a través de cada objeto JSON y enviar una solicitud POST
	for _, person := range people {
		// Convertir el objeto JSON en JSON
		jsonData, err := json.Marshal(person)
		if err != nil {
			log.Fatal(err)
		}

		// Crear una nueva solicitud POST con el JSON
		req, err := http.NewRequest("POST", postURL, bytes.NewBuffer(jsonData))
		if err != nil {
			log.Fatal(err)
		}
		req.Header.Set("Content-Type", "application/json")

		// Realizar la solicitud POST
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()

		//Verificar que todos los datos sean transmitidos
		failedResponses := []string{}
		if resp.StatusCode != http.StatusOK {
			failedResponses = append(failedResponses, fmt.Sprintf("Status Code %d: %s", resp.StatusCode, string(jsonData)))
		} else {
			fmt.Println("POST Status Code:", resp.StatusCode)
		}

		resp.Body.Close()

		// Imprimir las respuestas con status code diferente a correcto
		if len(failedResponses) > 0 {
			fmt.Println("Failed Responses:")

			for _, response := range failedResponses {
				fmt.Println(response)
			}
		}
	}
	fmt.Fprintln(w, "Solicitud realizada")

	DownloadData()
}

func formatPhone(phone string) string {
	// Extraer solo los dígitos del número de teléfono
	regex := regexp.MustCompile("[0-9]+")
	digits := regex.FindAllString(phone, -1)

	// Asegurarse de que haya 10 dígitos completando con ceros
	if len(digits[0]) == 7 {
		digits[0] = digits[0] + "000"
	}
	if len(digits[0]) == 8 {
		digits[0] = digits[0] + "00"
	}
	if len(digits[0]) == 9 {
		digits[0] = digits[0] + "0"
	}
	// Formatear según "xxx-xxx-xxxx"
	formatted := fmt.Sprintf("%s-%s-%s", digits[0][:3], digits[0][3:6], digits[0][6:])
	return formatted

}

func formatname(name string, phone string) string {
	// Extraer solo los dígitos del número de teléfono
	regex := regexp.MustCompile("[0-9]+")
	//Agregar asteriscos al nombre según le falten digito al número
	digits := regex.FindAllString(phone, -1)
	if len(digits[0]) == 7 {
		name = "***" + name
	}
	if len(digits[0]) == 8 {
		name = "**" + name
	}
	if len(digits[0]) == 9 {
		name = "*" + name
	}
	return name
}

// Descargar los datos
func DownloadData() {
	req1, err1 := http.NewRequest("GET", "https://8j5baasof2.execute-api.us-west-2.amazonaws.com/production/tests/trucode/items", nil)

	if err1 != nil {
		log.Fatal(err1)
	}

	req1.Header.Set("Content-Type", "application/json")

	resp1, err1 := http.DefaultClient.Do(req1)
	if err1 != nil {
		log.Fatal(err1)
	}
	defer resp1.Body.Close()

	log.Println(resp1.StatusCode)
	Body1, err1 := io.ReadAll(resp1.Body)
	if err1 != nil {
		log.Fatal(err1)
	}

	fmt.Println(string(Body1))
}
