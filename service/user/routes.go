package user

type Handler struct {

}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/login", h.handlelogin).Methods("POST")


}

func (h *Handler) handlelogin(w http.ResponseWriter, r *http.Request) {
	fmt.Println("login")
}

func (h *Handler) handlelogin(w http.ResponseWriter, r *http.Request) {
	fmt.Println("login")
}