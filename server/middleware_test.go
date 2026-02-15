package server

import (
	"math/rand"
	"strconv"
	"testing"
	"time"
)

func generarIDPeticion() string {
	rand.Seed(time.Now().UnixNano())
	id := rand.Intn(10000)
	return "REQ-" + strconv.Itoa(id)
}

// TEST: verifica que los middlewares se ejecuten en el orden correcto
func TestOrdenEjecucion(t *testing.T) {
	ResetTracker()
	finalHandler := func(request string) {
		Tracker = append(Tracker, "Handler")
	}

	handler := Chan(
		finalHandler,
		Logging,
		Auth,
	)

	requestID := generarIDPeticion()
	t.Logf("ID de petición: %s", requestID)

	handler("admin")

	ordenEsperado := []string{"Logging", "Auth", "Handler"}
	for i, nombre := range ordenEsperado {
		if Tracker[i] != nombre {
			t.Errorf("Orden incorrecto. Esperado: %s, Obtenido: %s", nombre, Tracker[i])
		}
	}
	t.Logf("Orden de ejecución correcto: %v", Tracker)
}

// TEST: verifica que si Auth falla, el Handler no se ejecute
func TestInterrupcionAuth(t *testing.T) {
	ResetTracker()
	finalHandler := func(request string) {
		Tracker = append(Tracker, "Handler")
	}
	handler := Chan(
		finalHandler,
		Logging,
		Auth,
	)

	requestID := generarIDPeticion()
	t.Logf("ID de petición: %s", requestID)

	handler("invitado")

	for _, nombre := range Tracker {
		if nombre == "Handler" {
			t.Error("El Handler se ejecutó cuando debería haberse interrumpido por Auth")
		}
	}
	t.Logf("Cadena interrumpida correctamente en Auth: %v", Tracker)
}
