package entity

import (
	"github.com/kzmake/example-uber-go-dig/domain/vo"
)

// Task ...
type Task struct {
	ID   vo.TaskID
	Data vo.TaskData
}
