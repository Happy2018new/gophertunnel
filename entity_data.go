package main

// PersistenceData ..
type PersistenceData struct {
	WorldEntity []*EntityData
}

// EntityData holds the basic data for a single entity.
type EntityData struct {
	EntityType      string // e.g. `minecraft:falling_block`
	EntityRuntimeID uint64 // e.g. `2018`
	EntityUniqueID  int64  // e.g. `-2018`
}

// AddWorldEntity ..
func (e *PersistenceData) AddWorldEntity(entityData EntityData) {
	e.WorldEntity = append(e.WorldEntity, &entityData)
}

// GetWorldEntityByRuntimeID ..
func (e *PersistenceData) GetWorldEntityByRuntimeID(entityRuntimeID uint64) *EntityData {
	for _, value := range e.WorldEntity {
		if value.EntityRuntimeID == entityRuntimeID {
			return value
		}
	}
	return nil
}

// GetWorldEntityByUniqueID ..
func (e *PersistenceData) GetWorldEntityByUniqueID(entityUniqueID int64) *EntityData {
	for _, value := range e.WorldEntity {
		if value.EntityUniqueID == entityUniqueID {
			return value
		}
	}
	return nil
}

// DeleteWorldEntityByRuntimeID ..
func (e *PersistenceData) DeleteWorldEntityByRuntimeID(entityRuntimeID uint64) {
	newer := make([]*EntityData, 0)
	for _, value := range e.WorldEntity {
		if value.EntityRuntimeID == entityRuntimeID {
			continue
		}
		newer = append(newer, value)
	}
	e.WorldEntity = newer
}

// DeleteWorldEntityByUniqueID ..
func (e *PersistenceData) DeleteWorldEntityByUniqueID(entityUniqueID int64) {
	newer := make([]*EntityData, 0)
	for _, value := range e.WorldEntity {
		if value.EntityUniqueID == entityUniqueID {
			continue
		}
		newer = append(newer, value)
	}
	e.WorldEntity = newer
}
