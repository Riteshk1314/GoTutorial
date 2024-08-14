package main
import(
	"log"
	"github.com/Riteshk1314/GoTutorial/cmd/api"

)
func main() {
	server := api.NewAPIServer("localhost:8080", nill)
	if err := server.Run(); err != nil {
		log.Fatalf("error running server: %v", err)
	}
}