package uid

type UniqueIdentifier[T uid] struct {
	value T
}

func (ui *UniqueIdentifier[T]) Value() any {
	return ui.value
}

func NewUniqueIdentifier[T uid](s T) (Identifier[T], error) {
	ui := &UniqueIdentifier[T]{value: s}
	return ui, nil
}
