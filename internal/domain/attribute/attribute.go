package attribute

import (
	"errors"
	"fmt"

	"backend/internal/helper"
)

type Attribute struct {
	id     string
	code   string
	name   string
	values []AttributeValue

	valuesToAdd      []AttributeValue
	valueIDsToUpdate []string
	valueIDsToRemove []string
}

func NewAttribute(code, name string, values []AttributeValue) *Attribute {
	return &Attribute{
		code:   code,
		name:   name,
		values: values,
	}
}

func (a *Attribute) GetID() string {
	return a.id
}

func (a *Attribute) GetCode() string {
	return a.code
}

func (a *Attribute) SetCode(code string) {
	a.code = code
}

func (a *Attribute) GetName() string {
	return a.name
}

func (a *Attribute) SetName(name string) {
	a.name = name
}

func (a *Attribute) GetValues() []AttributeValue {
	return a.values
}

func (a *Attribute) GetValuesToAdd() []AttributeValue {
	return a.valuesToAdd
}

func (a *Attribute) GetValueIDsToRemove() []string {
	return a.valueIDsToRemove
}

func (a *Attribute) AddValues(values []AttributeValue) error {
	valueIDs := make([]string, 0, len(values))
	for _, value := range values {
		valueIDs = append(valueIDs, value.GetID())
	}

	if len(valueIDs) > 0 {
		existing := helper.FindExisting(
			a.values,
			valueIDs,
			func(v AttributeValue) string {
				return v.GetID()
			},
		)
		if len(existing) > 0 {
			return fmt.Errorf("values already exist with IDs: %v", existing)
		}
	}

	a.values = append(a.values, values...)
	a.valuesToAdd = append(a.valuesToAdd, values...)
	return nil
}

func (a *Attribute) UpdateValues(values []AttributeValue) error {
	valueIDs := make([]string, 0, len(values))
	for _, value := range values {
		if value.GetID() == "" {
			return errors.New("all values must have an ID to be updated")
		}
		valueIDs = append(valueIDs, value.GetID())
	}

	allExist := helper.CheckAllExist(a.values, valueIDs, func(v AttributeValue) string {
		return v.GetID()
	})
	if !allExist {
		nonExistent := helper.FindNonExistent(a.values, valueIDs, func(v AttributeValue) string {
			return v.GetID()
		})
		return fmt.Errorf("values do not exist with IDs: %v", nonExistent)
	}

	existValues := make(map[string]*AttributeValue, len(a.values))
	for i := range a.values {
		existValues[a.values[i].GetID()] = &a.values[i]
	}

	for _, value := range values {
		existValues[value.GetID()].SetValue(value.GetValue())
	}

	a.valueIDsToUpdate = append(a.valueIDsToUpdate, valueIDs...)
	return nil
}

func (a *Attribute) RemoveValuesByID(valueIDs []string) error {
	allExist := helper.CheckAllExist(a.values, valueIDs, func(v AttributeValue) string {
		return v.GetID()
	})
	if !allExist {
		nonExistent := helper.FindNonExistent(a.values, valueIDs, func(v AttributeValue) string {
			return v.GetID()
		})
		return fmt.Errorf("values do not exist with IDs: %v", nonExistent)
	}

	toRemoveSet := make(map[string]struct{}, len(valueIDs))
	for _, id := range valueIDs {
		toRemoveSet[id] = struct{}{}
	}

	newValues := a.values[:0]
	for _, value := range a.values {
		if _, shouldRemove := toRemoveSet[value.GetID()]; !shouldRemove {
			newValues = append(newValues, value)
		}
	}
	clear(a.values[len(newValues):])
	a.values = newValues

	a.valueIDsToRemove = append(a.valueIDsToRemove, valueIDs...)
	return nil
}
