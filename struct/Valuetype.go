package valuetype

type (
	Employee struct {
		Name       string
		Age        int
		Skill      []string
		Department Department
	}
	Department struct {
		DepartmentName string
	}
)
