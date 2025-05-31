package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"TaipeiCityDashboardBE/app/models"

	"github.com/gin-gonic/gin"
)

// GeoJSON structures for response
type GeoJSONFeatureCollection struct {
	Type     string           `json:"type"`
	Features []GeoJSONFeature `json:"features"`
}

type GeoJSONFeature struct {
	ID         string                 `json:"id"`
	Type       string                 `json:"type"`
	Properties map[string]interface{} `json:"properties"`
	Geometry   Geometry               `json:"geometry"`
}

type Geometry struct {
	Type        string        `json:"type"`
	Coordinates [][][]float64 `json:"coordinates"`
}

// parseCoordinates parses the map string into coordinates array
func parseCoordinates(mapStr string) ([][][]float64, error) {
	if mapStr == "" {
		return nil, fmt.Errorf("empty map string")
	}

	// Parse the coordinate string like "[[121.46817932209012, 25.120902275924415], ...]"
	var coordinates [][]float64
	err := json.Unmarshal([]byte(mapStr), &coordinates)
	if err != nil {
		return nil, fmt.Errorf("failed to parse coordinates: %v", err)
	}

	// Convert to the format required by GeoJSON Polygon
	return [][][]float64{coordinates}, nil
}

// convertToGeoJSON converts courses to GeoJSON format
func convertToGeoJSON(courses []models.ElderlyFitnessCourse) GeoJSONFeatureCollection {
	var features []GeoJSONFeature

	for _, course := range courses {
		coordinates, err := parseCoordinates(course.Map)
		if err != nil {
			// Skip courses with invalid coordinates but continue processing
			continue
		}

		// Convert annual_district_ratio from string to int if possible
		annualRatio := 0
		if course.AnnualDistrictRatio != "" {
			if ratio, err := strconv.Atoi(course.AnnualDistrictRatio); err == nil {
				annualRatio = ratio
			}
		}

		properties := map[string]interface{}{
			"id":                           fmt.Sprintf("%d", course.ID),
			"org_name":                     course.OrgName,
			"provide_meal":                 course.ProvideMeal,
			"annual_district_ratio":        annualRatio,
			"annual_expected_participants": course.AnnualExpectedParticipants,
			"course_status":                course.CourseStatus,
			"popular_course":               course.PopularCourse,
		}

		feature := GeoJSONFeature{
			ID:         fmt.Sprintf("%d", course.ID), // Use database ID as feature ID
			Type:       "Feature",
			Properties: properties,
			Geometry: Geometry{
				Type:        "Polygon",
				Coordinates: coordinates,
			},
		}

		features = append(features, feature)
	}

	return GeoJSONFeatureCollection{
		Type:     "FeatureCollection",
		Features: features,
	}
}

/*
GetAllElderlyFitnessCourses returns all elderly fitness courses information in GeoJSON format
GET /api/v1/elderly-fitness-courses
*/
func GetAllElderlyFitnessCourses(c *gin.Context) {
	courses, err := models.GetAllElderlyFitnessCourses()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": err.Error()})
		return
	}

	// Convert to GeoJSON format
	geoJSON := convertToGeoJSON(courses)
	c.JSON(http.StatusOK, geoJSON)
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
