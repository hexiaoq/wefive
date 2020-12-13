package router

import (
	"github.com/gin-gonic/gin"
	"wefive/controller"
	"wefive/middleware"
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
	r.POST("bus/get", controller.SendBusiness)
	r.POST("bus/update", controller.UpdateBusiness)

	r.POST("material/get", controller.SendMaterials)
	r.POST("material/add", controller.AddMaterial)
	r.POST("material/update", controller.UpdateMaterial)
	r.POST("material/delete", controller.DeleteMaterial)
	return r
}
