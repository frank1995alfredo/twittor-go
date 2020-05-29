package models

//Tweet ... se graba el mensaje del tweet
type Tweet struct {
	Mensaje string `bson:"mensaje" json:"mensaje"`
}
