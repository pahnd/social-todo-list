package ginuser

import (
	"net/http"
	"social-todo-list/common"
	"social-todo-list/component/tokenprovider"
	"social-todo-list/module/user/biz"
	"social-todo-list/module/user/model"
	"social-todo-list/module/user/storage"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Login(db *gorm.DB, tokenProvider tokenprovider.Provider) gin.HandlerFunc {
	return func(c *gin.Context) {
		var loginUserData model.UserLogin

		if err := c.ShouldBind(&loginUserData); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := storage.NewSQLStore(db)
		HahsedPassword := common.NewHashPassword()

		business := biz.NewLoginBusiness(store, tokenProvider, HahsedPassword, 60*60*24*7)
		account, err := business.Login(c.Request.Context(), &loginUserData)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(account))
	}
}
