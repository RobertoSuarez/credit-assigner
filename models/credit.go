package models

// Credit Modelo que se base la tabla en la db.
type Credit struct {
	ID              int
	Investment      int32
	Credit_type_300 int32
	Credit_type_500 int32
	Credit_type_700 int32
	Successful      bool
}
