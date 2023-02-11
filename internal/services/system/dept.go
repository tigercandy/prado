package system

import "github.com/tigercandy/prado/internal/models/system"

type sysDeptService struct {
}

var (
	SysDeptService = new(sysDeptService)
	deptModel      = new(system.DeptModel)
)

func (sds *sysDeptService) SetDept() ([]system.DeptModel, error) {
	list, err := deptModel.GetPage()

	m := make([]system.DeptModel, 0)
	for i := 0; i < len(list); i++ {
		if list[i].ParentId != 0 {
			continue
		}
		info := sds.Digui(&list, list[i])
		m = append(m, info)
	}
	return m, err
}

func (sds *sysDeptService) Digui(deptList *[]system.DeptModel, menu system.DeptModel) system.DeptModel {
	list := *deptList

	min := make([]system.DeptModel, 0)
	for j := 0; j < len(list); j++ {
		if menu.DeptId != list[j].ParentId {
			continue
		}
		m := system.DeptModel{}
		m.DeptId = list[j].DeptId
		m.ParentId = list[j].ParentId
		m.DeptPath = list[j].DeptPath
		m.DeptName = list[j].DeptName
		m.Sort = list[j].Sort
		m.Leader = list[j].Leader
		m.Phone = list[j].Phone
		m.Email = list[j].Email
		m.Status = list[j].Status
		m.CreatedAt = list[j].CreatedAt
		m.UpdatedAt = list[j].UpdatedAt
		m.Children = []system.DeptModel{}
		ms := sds.Digui(deptList, m)
		min = append(min, ms)
	}
	menu.Children = min
	return menu
}

func (sds *sysDeptService) SetDeptLabel() {

}
