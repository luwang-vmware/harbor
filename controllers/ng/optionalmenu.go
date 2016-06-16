package ng

import (
	"net/http"

	"github.com/vmware/harbor/dao"
	"github.com/vmware/harbor/models"
	"github.com/vmware/harbor/utils/log"
)

// OptionalMenuController handles request to /ng/optional_menu
type OptionalMenuController struct {
	BaseController
}

// Get renders optional menu, Admin user has "Add User" menu
func (omc *OptionalMenuController) Get() {
	sessionUserID := omc.GetSession("userId")

	var hasLoggedIn bool
	if sessionUserID != nil {
		hasLoggedIn = true
		userID := sessionUserID.(int)
		u, err := dao.GetUser(models.User{UserID: userID})
		if err != nil {
			log.Errorf("Error occurred in GetUser, error: %v", err)
			omc.CustomAbort(http.StatusInternalServerError, "Internal error.")
		}
		if u == nil {
			log.Warningf("User was deleted already, user id: %d, canceling request.", userID)
			omc.CustomAbort(http.StatusUnauthorized, "")
		}
		omc.Data["Username"] = u.Username
	}
	omc.Data["HasLoggedIn"] = hasLoggedIn
	omc.TplName = "ng/optional-menu.htm"
	omc.Render()

}
