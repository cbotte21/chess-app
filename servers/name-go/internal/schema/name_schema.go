package schema

//user struct

const DATABASE string = "userdata"
const COLLECTION string = "name"

type Name struct { //Payload
	Id   string `bson:"_id,omitempty"`
	Name string `bson:"name,omitempty"`
}

func (name Name) Database() string {
	return DATABASE
}

func (name Name) Collection() string {
	return COLLECTION
}
