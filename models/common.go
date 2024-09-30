package models

const (
	EVENT int = iota
	BOOKING
)

type (
	ModelType[T any] struct {
		model   T
		typeRef int
	}

	IModelType[T any] interface {
		GetModel() T
		GetTypeRef() int
	}
)

var ModelTypes = map[int]string{
	EVENT:   "Event",
	BOOKING: "Booking",
}

func GetModelType(modelType int) string {
	return ModelTypes[modelType]
}

func (model ModelType[T]) GetModel() T {
	return model.model
}

func GetBookingModel(model IModelType[Booking]) Booking {
	return model.GetModel()
}

func GetEventModel(model IModelType[Event]) Event {
	return model.GetModel()
}

func (modelType ModelType[T]) GetTypeRef() int {
	return modelType.typeRef
}

func NewModelType[T any](model T, typeRef int) ModelType[T] {
	return ModelType[T]{model, typeRef}
}
