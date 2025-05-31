package controllers

import (
	"net/http"

	"TaipeiCityDashboardBE/app/models"

	"github.com/gin-gonic/gin"
)

/*
GetAllElderlyFitnessCourses returns all elderly fitness courses information
GET /api/v1/elderly-fitness-courses
*/
func GetAllElderlyFitnessCourses(c *gin.Context) {
	courses, err := models.GetAllElderlyFitnessCourses()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "total": len(courses), "data": courses})
}

/*
GetElderlyFitnessCoursesWithPagination returns elderly fitness courses with pagination
GET /api/v1/elderly-fitness-courses/paginated
*/
func GetElderlyFitnessCoursesWithPagination(c *gin.Context) {
	type courseQuery struct {
		PageSize int    `form:"pagesize"`
		PageNum  int    `form:"pagenum"`
		Sort     string `form:"sort"`
		Order    string `form:"order"`
	}

	// Get query parameters
	var query courseQuery
	c.ShouldBindQuery(&query)

	courses, totalCourses, err := models.GetElderlyFitnessCoursesWithPagination(query.PageSize, query.PageNum, query.Sort, query.Order)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "total": totalCourses, "data": courses})
}

/*
GetElderlyFitnessCoursesByCity returns elderly fitness courses filtered by city
GET /api/v1/elderly-fitness-courses/city/:city
*/
func GetElderlyFitnessCoursesByCity(c *gin.Context) {
	city := c.Param("city")
	if city == "" {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "City parameter is required"})
		return
	}

	courses, err := models.GetElderlyFitnessCoursesByCity(city)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "total": len(courses), "data": courses})
}

/*
GetElderlyFitnessCoursesByDistrict returns elderly fitness courses filtered by district
GET /api/v1/elderly-fitness-courses/district/:district
*/
func GetElderlyFitnessCoursesByDistrict(c *gin.Context) {
	district := c.Param("district")
	if district == "" {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "District parameter is required"})
		return
	}

	courses, err := models.GetElderlyFitnessCoursesByDistrict(district)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "total": len(courses), "data": courses})
}

/*
GetElderlyFitnessCoursesByCategory returns elderly fitness courses filtered by category
GET /api/v1/elderly-fitness-courses/category/:category
*/
func GetElderlyFitnessCoursesByCategory(c *gin.Context) {
	category := c.Param("category")
	if category == "" {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Category parameter is required"})
		return
	}

	courses, err := models.GetElderlyFitnessCoursesByCategory(category)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "total": len(courses), "data": courses})
}

/*
GetElderlyFitnessCoursesByOrganization returns elderly fitness courses filtered by organization
GET /api/v1/elderly-fitness-courses/organization/:orgName
*/
func GetElderlyFitnessCoursesByOrganization(c *gin.Context) {
	orgName := c.Param("orgName")
	if orgName == "" {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Organization name parameter is required"})
		return
	}

	courses, err := models.GetElderlyFitnessCoursesByOrganization(orgName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "total": len(courses), "data": courses})
}

/*
SearchElderlyFitnessCourses returns elderly fitness courses with multiple filter options
GET /api/v1/elderly-fitness-courses/search
*/
func SearchElderlyFitnessCourses(c *gin.Context) {
	type searchQuery struct {
		City     string `form:"city"`
		District string `form:"district"`
		Category string `form:"category"`
		OrgName  string `form:"org_name"`
		PageSize int    `form:"pagesize"`
		PageNum  int    `form:"pagenum"`
		Sort     string `form:"sort"`
		Order    string `form:"order"`
	}

	var query searchQuery
	c.ShouldBindQuery(&query)

	// Start with all courses and apply filters
	var courses []models.ElderlyFitnessCourse
	var err error

	if query.City != "" {
		courses, err = models.GetElderlyFitnessCoursesByCity(query.City)
	} else if query.District != "" {
		courses, err = models.GetElderlyFitnessCoursesByDistrict(query.District)
	} else if query.Category != "" {
		courses, err = models.GetElderlyFitnessCoursesByCategory(query.Category)
	} else if query.OrgName != "" {
		courses, err = models.GetElderlyFitnessCoursesByOrganization(query.OrgName)
	} else {
		// If no specific filter, use pagination
		var totalCourses int64
		courses, totalCourses, err = models.GetElderlyFitnessCoursesWithPagination(query.PageSize, query.PageNum, query.Sort, query.Order)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "success", "total": totalCourses, "data": courses})
		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "total": len(courses), "data": courses})
}
