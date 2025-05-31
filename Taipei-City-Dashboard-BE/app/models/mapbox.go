package models

/* ----- Models ----- */

// ElderlyFitnessCourse represents the elderly fitness courses distribution data
type ElderlyFitnessCourse struct {
	ID                         int64  `json:"id" gorm:"column:id;autoincrement;primaryKey"`
	City                       string `json:"city" gorm:"column:城市;type:varchar(10)"`
	District                   string `json:"district" gorm:"column:地區;type:varchar(255)"`
	Address                    string `json:"address" gorm:"column:地址;type:text"`
	Course                     string `json:"course" gorm:"column:課程;type:varchar(255)"`
	Category                   string `json:"category" gorm:"column:課程分類;type:varchar(255)"`
	OrgName                    string `json:"org_name" gorm:"column:組織名稱;type:varchar(255)"`
	ProvideMeal                string `json:"provide_meal" gorm:"column:是否提供共餐;type:varchar(5)"`
	AnnualDistrictRatio        string `json:"annual_district_ratio" gorm:"column:年度開課地區比例;type:varchar(255)"`
	AnnualExpectedParticipants int    `json:"annual_expected_participants" gorm:"column:年度預計參與人數;type:int"`
	CourseStatus               string `json:"course_status" gorm:"column:開課狀態;type:varchar(255)"`
	PopularCourse              string `json:"popular_course" gorm:"column:熱門課程;type:varchar(255)"`
	Map                        string `json:"map" gorm:"column:地圖;type:text"`
}

/* ----- Handlers ----- */

// GetAllElderlyFitnessCourses retrieves all elderly fitness courses from the database
func GetAllElderlyFitnessCourses() (courses []ElderlyFitnessCourse, err error) {
	err = DBDashboard.Table("elderly_fitness_distribution_all").Find(&courses).Error
	return courses, err
}

// GetElderlyFitnessCoursesByCity retrieves elderly fitness courses filtered by city
func GetElderlyFitnessCoursesByCity(city string) (courses []ElderlyFitnessCourse, err error) {
	err = DBDashboard.Table("elderly_fitness_distribution_all").
		Where("城市 = ?", city).
		Find(&courses).Error
	return courses, err
}

// GetElderlyFitnessCoursesByDistrict retrieves elderly fitness courses filtered by district
func GetElderlyFitnessCoursesByDistrict(district string) (courses []ElderlyFitnessCourse, err error) {
	err = DBDashboard.Table("elderly_fitness_distribution_all").
		Where("地區 = ?", district).
		Find(&courses).Error
	return courses, err
}

// GetElderlyFitnessCoursesByCategory retrieves elderly fitness courses filtered by category
func GetElderlyFitnessCoursesByCategory(category string) (courses []ElderlyFitnessCourse, err error) {
	err = DBDashboard.Table("elderly_fitness_distribution_all").
		Where("課程分類 = ?", category).
		Find(&courses).Error
	return courses, err
}

// GetElderlyFitnessCoursesByOrganization retrieves elderly fitness courses filtered by organization name
func GetElderlyFitnessCoursesByOrganization(orgName string) (courses []ElderlyFitnessCourse, err error) {
	err = DBDashboard.Table("elderly_fitness_distribution_all").
		Where("組織名稱 = ?", orgName).
		Find(&courses).Error
	return courses, err
}

// GetElderlyFitnessCoursesByMealProvision retrieves elderly fitness courses filtered by meal provision
func GetElderlyFitnessCoursesByMealProvision(provideMeal string) (courses []ElderlyFitnessCourse, err error) {
	err = DBDashboard.Table("elderly_fitness_distribution_all").
		Where("是否提供共餐 = ?", provideMeal).
		Find(&courses).Error
	return courses, err
}

// GetElderlyFitnessCoursesByStatus retrieves elderly fitness courses filtered by course status
func GetElderlyFitnessCoursesByStatus(status string) (courses []ElderlyFitnessCourse, err error) {
	err = DBDashboard.Table("elderly_fitness_distribution_all").
		Where("開課狀態 = ?", status).
		Find(&courses).Error
	return courses, err
}

// GetPopularElderlyFitnessCourses retrieves popular elderly fitness courses
func GetPopularElderlyFitnessCourses() (courses []ElderlyFitnessCourse, err error) {
	err = DBDashboard.Table("elderly_fitness_distribution_all").
		Where("熱門課程 IS NOT NULL AND 熱門課程 != ''").
		Find(&courses).Error
	return courses, err
}

// GetElderlyFitnessCoursesWithParticipants retrieves elderly fitness courses filtered by participant count range
func GetElderlyFitnessCoursesWithParticipants(minParticipants, maxParticipants int) (courses []ElderlyFitnessCourse, err error) {
	query := DBDashboard.Table("elderly_fitness_distribution_all")

	if minParticipants > 0 {
		query = query.Where("年度預計參與人數 >= ?", minParticipants)
	}

	if maxParticipants > 0 {
		query = query.Where("年度預計參與人數 <= ?", maxParticipants)
	}

	err = query.Find(&courses).Error
	return courses, err
}

// GetElderlyFitnessCoursesWithPagination retrieves elderly fitness courses with pagination support
func GetElderlyFitnessCoursesWithPagination(pageSize int, pageNum int, sort string, order string) (courses []ElderlyFitnessCourse, totalCourses int64, err error) {
	tempDB := DBDashboard.Table("elderly_fitness_distribution_all")

	// Count the total amount of courses
	tempDB.Count(&totalCourses)

	// Sort the courses
	if sort != "" {
		tempDB = tempDB.Order(sort + " " + order)
	} else {
		tempDB = tempDB.Order("id asc")
	}

	// Paginate the courses
	if pageSize > 0 {
		tempDB = tempDB.Limit(pageSize)
		if pageNum > 0 {
			tempDB = tempDB.Offset((pageNum - 1) * pageSize)
		}
	}

	// Get the courses
	err = tempDB.Find(&courses).Error

	return courses, totalCourses, err
}
