package UserList

import (
	"net/http"
	"order/app/Common"
	"order/app/Http/Models/User"
	"order/app/Http/Request"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type PageInfo struct {
	CurrentPage int    `json:"current_page" binding:"required,min=1" label:"当前页数"`       // UUID 类型
	PageSize    int    `json:"page_size" binding:"required,min=5,max=9999" label:"每页条数"` // 自定义校验
	Date        string `json:"date" binding:"omitempty,datetime=2006-01-02,check_date"`
}

func Info(ctx *gin.Context) {
	currentPage := Common.Input(ctx, "current_page")
	pageSize := Common.Input(ctx, "page_size")
	Date := Common.Input(ctx, "date")
	page, _ := strconv.Atoi(currentPage)
	size, _ := strconv.Atoi(pageSize)
	s := PageInfo{
		CurrentPage: page,
		PageSize:    size,
		Date:        Date,
	}
	// fmt.Println(s)
	if err := ctx.ShouldBind(&s); err != nil {
		// 获取validator.ValidationErrors类型的errors
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			// 非validator.ValidationErrors类型错误直接返回
			ctx.JSON(http.StatusOK, gin.H{
				"code":    500,
				"message": err.Error(),
			})
		}
		// validator.ValidationErrors类型错误则进行翻译
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": Request.RemoveTopStruct(errs.Translate(Request.Trans)),
		})
		return
	}

	user, _ := User.GetUserList(page, size)
	// 接口返回
	back := gin.H{
		"userInfo": user,
	}
	ctx.JSON(200, back)
}

func Active(ctx *gin.Context) {
}
