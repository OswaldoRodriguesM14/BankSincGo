package compose_handler

import (
	usecases_cora_api "github.com/OswaldoRodriguesM14/BankSincGo/internal/entity/usecases"
	usecases_cora_db "github.com/OswaldoRodriguesM14/BankSincGo/internal/infra/database/usecases"
	"github.com/OswaldoRodriguesM14/BankSincGo/internal/infra/webserver/handlers"
	pkg "github.com/OswaldoRodriguesM14/BankSincGo/pkg/utilits"
	"go.mongodb.org/mongo-driver/mongo"
)

func CoraHanlder(Session *mongo.Client) *handlers.CoraHandler {
	Utilits := pkg.UtilitsCompose()
	CoraDB := usecases_cora_db.CoraDBCompose(Session)
	CoraAPI := usecases_cora_api.CoraAPICompose(Utilits)
	return handlers.CoraHandlerComposer(CoraDB, CoraAPI, Utilits)

}
