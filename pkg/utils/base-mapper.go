package utils

type BaseMapper[Entity any, DTO any] interface {
    ToEntity(dto DTO) Entity

    ToEntities(dtos []DTO) []Entity

    ToDTO(entity Entity) DTO

    ToDTOs(entities []Entity) []DTO

    UpdateEntityFromDTO(dto DTO, entity *Entity) Entity

    UpdateEntitiesFromDTOs(dtos []DTO, entities []Entity) []Entity
}
