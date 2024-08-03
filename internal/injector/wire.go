package injector

//
//import (
//	"github.com/google/wire"
//	"go-gin-web/internal/controller"
//	"go-gin-web/internal/database"
//	"go-gin-web/internal/repository"
//	"go-gin-web/internal/server"
//	"go-gin-web/internal/service"
//	"gorm.io/gorm"
//)
//
//type App struct {
//	DB             *gorm.DB
//	RoleRepository repository.RoleRepository
//	UserRepository repository.UserRepository
//	RoleService    service.RoleService
//	UserService    service.UserService
//	AuthService    service.AuthService
//	UserController controller.UserController
//	RoleController controller.RoleController
//	AppController  controller.AppController
//	AppServer      *server.AppServer
//}
//
//func InitializeApp() (*App, error) {
//	wire.Build(
//		database.New,
//		repository.NewRoleRepository,
//		repository.NewUserRepository,
//		service.NewRoleService,
//		service.NewUserService,
//		service.NewAuthService,
//		controller.NewUserController,
//		controller.NewRoleController,
//		controller.NewAppController,
//		server.NewAppServer,
//		wire.Struct(new(App), "*"),
//	)
//
//	return nil, nil
//}
