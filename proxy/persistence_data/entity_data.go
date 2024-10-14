package persistence_data

// 描述单个实体的数据
type EntityData struct {
	EntityType      string // 该实体的英文 ID
	EntityRuntimeID uint64 // 该实体的运行时 ID
	EntityUniqueID  int64  // 该实体的唯一 ID
}

// ...
func (e *PersistenceData) AddWorldEntity(entityData EntityData) {
	e.WorldEntity = append(e.WorldEntity, &entityData)
}

// ...
func (e *PersistenceData) GetWorldEntityByRuntimeID(entityRuntimeID uint64) *EntityData {
	for _, value := range e.WorldEntity {
		if value.EntityRuntimeID == entityRuntimeID {
			return value
		}
	}
	return nil
}

// ...
func (e *PersistenceData) GetWorldEntityByUniqueID(entityUniqueID int64) *EntityData {
	for _, value := range e.WorldEntity {
		if value.EntityUniqueID == entityUniqueID {
			return value
		}
	}
	return nil
}

// ...
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

// ...
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
