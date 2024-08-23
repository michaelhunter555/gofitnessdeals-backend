package models

import "go.mongodb.org/mongo-driver/bson/primitive"

//set user, reviews and credibility rating
type User struct {
	ID       primitive.ObjectID   `bson:"_id,omitempty" json:"id,omitempty"`
	Username string               `bson:"username,omitempty" json:"username,omitempty"`
	Email    string               `bson:"email,omitempty" json:"email,omitempty"`
	Password string               `bson:"password,omitempty" json:"password,omitempty"`
	Reviews  []primitive.ObjectID `bson:"reviews,omitempty" json:"reviews,omitempty"`
	Cred     int                  `bson:"cred, omitempty" json:"cred,omitempty"`
}

//review ref to User && product
type Review struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	ProductId   primitive.ObjectID `bson:"product_id,omitempty" json:"product_id,omitempty"`
	UserId      primitive.ObjectID `bson:"user_id,omitempty" json:"user_id,omitempty"`
	Description string             `bson:"description,omitempty" json:"description,omitempty"`
	Rating      int                `bson:"rating,omitempty" json:"rating,omitempty"`
	ImgUrl      string             `bson:"img_url,omitempty" json:"img_url,omitempty"`
}

//product details
type Product struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	ProductName string             `bson:"name,omitempty" json:"name,omitempty"`
	ProductType string             `bson:"type,omitempty" json:"type,omitempty"`
	ImgUrl      string             `bson:"img_Url,omitempty" json:"img_Url,omitempty"`
	Category    string             `bson:"category,omitempty" json:"category,omitempty"`
	Description string             `bson:"description,omitempty" json:"description,omitempty"`
	Price       string             `bson:"" json:"price,omitempty"`
	Rating      int                `bson:"rating,omitempty" json:"rating,omitempty"`
}

//Blog
type Blog struct {
	ID                primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	ReviewTitle       string             `bson:"review_title,omitempty" json:"review_title,omitempty"`
	ReviewDescription string             `bson:"review_description,omitempty" json:"review_description,omitempty"`
	ReviewRating      int                `bson:"review_rating,omitempty" json:"review_rating,omitempty"`
	ImgUrl            string             `bson:"img_url,omitempty" json:"img_url,omitempty"`
}
