package data

import (
	"fmt"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"test/week4_service_framework/internal/pkg/database"
)

const (
	tableMachineLifecycleTemplate = "machine_lifecycle_template"
	tableMachineLifecycleNode     = "machine_lifecycle_node"
	tableMachineLifecycleLink     = "machine_lifecycle_link"
)

type MachineLifecycleTemplate struct {
	ID   int
	Name string
}

type MachineLifecycleNode struct {
	ID         int
	TemplateID int
	Name       string
	FuncName   string
}

type MachineLifecycleLink struct {
	ID         int
	TemplateID int
	Name       string
	PrevNodeID string
	NextNodeID string
}

func GetNextNode(templateID int, currNodeID int) (*MachineLifecycleNode, error) {
	var err error
	nextNode := &MachineLifecycleNode{}
	link := &MachineLifecycleLink{}
	err = database.DB.
		Table(tableMachineLifecycleLink).
		Where("template_id = ? and prev_node_id = ?", templateID, currNodeID).
		Find(link).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nextNode, nil
	} else if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("select table %s, template id %d, prev_nodeID %d",
			tableMachineLifecycleLink,
			templateID,
			currNodeID,
		))
	}
	err = database.DB.
		Table(tableMachineLifecycleNode).
		Where("id = ?", link.NextNodeID).
		Find(nextNode).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nextNode, nil
	} else if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("select table %s, node id %d",
			tableMachineLifecycleNode,
			link.NextNodeID,
		))
	}
	return nextNode, nil
}
