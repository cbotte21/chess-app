package schema

//user struct

const DATABASE string = "userdata"
const COLLECTION string = "username"

type Username struct { //Payload
	Id       string `bson:"_id,omitempty"`
	Username string `bson:"username,omitempty"`
}

func (username Username) Database() string {
	return DATABASE
}

func (username Username) Collection() string {
	return COLLECTION
}
