package model

type Arrival struct {
	UserID int  `bson:"user_id"`
	Year   int  `bson:"year"`
	Week   int  `bson:"week"`
	Spot   Spot `bson:"spot"`
}

type Quest struct {
	Desc  string `bson:"desc"`
	Year  int    `bson:"year"`
	Week  int    `bson:"week"`
	Spots []Spot
}

type Spot struct {
	Name      string  `bson:"name"`
	Type      int     `bson:"type"`
	Latitude  float64 `bson:"latitude"`
	Longitude float64 `bson:"longitude"`
	Finished  bool    `bson:"finished"`
}
