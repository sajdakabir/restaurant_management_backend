// package models

// import (
// 	"time"

// 	"go.mongodb.org/mongo-driver/bson/primitive"
// )

// type Invoice struct {
// 	ID               primitive.ObjectID `bson:"_id"`
// 	Invoice_id       string             `json:"invoice_id"`
// 	Order_id         string             `json:"order_id" validate:"eq=CARD|eq=CASH|eq="`
// 	Payment_method   *string
// 	Payment_status   *string
// 	Payment_due_date time.Time
// 	Created_at       time.Time
// 	Updated_at       time.Time
// }
