// router.go

package routes

import (
	compose_handler "github.com/OswaldoRodriguesM14/BankSincGo/internal/infra/webserver/compose"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/jwtauth"
	"go.mongodb.org/mongo-driver/mongo"
)

var TokenAuth = jwtauth.New("HS256", []byte("Osw4ld0ro"), nil)

func NewRouter(SessionDB *mongo.Client) *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.WithValue("jwt", TokenAuth))

	RouteAdmin(r, SessionDB)

	return r
}

func RouteAdmin(r *chi.Mux, db *mongo.Client) {
	// Crie um novo roteador para as rotas relacionadas à ReuniSinc
	ReuniSincRouter := chi.NewRouter()
	CoraHandler := compose_handler.CoraHanlder(db)

	// Defina o handler para a rota /nucleos
	ReuniSincRouter.Post("/cadastra_certificado", CoraHandler.SalvaCertificadoCora)
	ReuniSincRouter.Get("/gera_certificado/{id}", CoraHandler.GeraAccesToken)

	ReuniSincRouter.Get("/consulta_extrato", CoraHandler.ConsultaExtrato)
	ReuniSincRouter.Get("/consulta_extrato/tipo/{id}", CoraHandler.ConsultaExtratoType)

	ReuniSincRouter.Post("/emite_boleto/cora/", CoraHandler.EmiteBoleto)

	// Monte o roteador ReuniSincRouter em /admin/
	r.Mount("/admin/", ReuniSincRouter)
}
