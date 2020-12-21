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
	r.GET("department/get/:deptId", controller.SendDepartment)
	r.POST("department/delete", controller.DepartmentDelete)
	r.POST("department/update", controller.UpdateDepartment)

	r.GET("gover/all", controller.SendAllGovernors)
	r.POST("gover/add", controller.AddGovernor)
	r.POST("gover/delete", controller.DeleteGovernor)

	r.GET("business/all/:deptId", controller.SendAllBusinessOfDept)
	r.POST("business/:name/add", controller.AddBusiness)
	r.POST("business/:name/delete", controller.DeleteBusiness)
	r.POST("businesses/get", controller.SendAllBusinessOfDeptByDeptId)
	r.GET("businesses/getHot", controller.SendHotBusiness)
	r.POST("bus/get", controller.SendBusiness)
	r.POST("bus/update", controller.UpdateBusiness)

	r.POST("material/get", controller.SendMaterials)
	r.POST("material/add", controller.AddMaterial)
	r.POST("material/update", controller.UpdateMaterial)
	r.POST("material/delete", controller.DeleteMaterial)

	r.POST("process/add", controller.AddProcessForBus)
	r.POST("process/addMaterial/:processId", controller.AddProcessMaterial)
	r.GET("process/all/:busId", controller.SendAllProcessOfBus)
	/*r.GET("process/busGet/:busId", controller.SendBusProcess)*/
	r.POST("process/delete", controller.DeleteProcess)
	r.POST("process/deleteMaterial", controller.DeleteProcessMaterial)

	return r
}
