package router

import (
	"claps-admin/controller"
	"claps-admin/middleware"
	"github.com/gin-gonic/gin"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	r.Use(middleware.CORSMiddleware())
	r.POST("tokenAuth", middleware.TokenAuthHandler)

	r.POST("login", controller.LoginCheck)
	r.POST("loginAdmin", controller.LoginAdminCheck)

	r.POST("department/add", controller.AddDepartment)
	r.GET("department/all", controller.SendAllDepartments)
	r.POST("department/delete", controller.DepartmentDelete)

	r.GET("gover/all", controller.SendAllGovernors)
	r.POST("gover/add", controller.AddGovernor)
	r.POST("gover/delete", controller.DeleteGovernor)

	r.GET("business/:name/all", controller.SendAllBusinessOfDept)
	r.POST("business/:name/add", controller.AddBusiness)
	r.POST("business/:name/delete", controller.DeleteBusiness)
	return r
}
