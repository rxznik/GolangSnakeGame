package world

import "github.com/rxznik/GolangSnakeGame/internal/entity"

type World struct {
	entities []entity.Entity
}

func New() *World {
	return &World{
		entities: []entity.Entity{},
	}
}

func (w *World) AddEntity(entity entity.Entity) {
	w.entities = append(w.entities, entity)
}

func (w World) Entities() []entity.Entity {
	return w.entities
}

func (w World) GetEntities(tag string) []entity.Entity {
	var result []entity.Entity
	for _, e := range w.entities {
		if e.Tag() == tag {
			result = append(result, e)
		}
	}
	return result
}

func (w World) GetFirstEntity(tag string) entity.Entity {
	for _, e := range w.entities {
		if e.Tag() == tag {
			return e
		}
	}
	return nil
}
