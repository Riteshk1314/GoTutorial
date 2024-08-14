package main
import(
	"log"


)
func main() {
	server := api.NewAPIServer("localhost:8080", nill)
	if err := server.Run(); err != nil {
		log.Fatalf("error running server: %v", err)
	}
}