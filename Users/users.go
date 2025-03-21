package Users

import (
	"fmt"
	"time"

	"github.com/GoLangEcomerce/models"
)

func AltaUsuario() {
	u := new(models.User)
	u.AddUser(10, "pablo ", time.Now(), true)
	fmt.Println(u)
}
