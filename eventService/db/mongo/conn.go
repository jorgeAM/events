package mongo

import (
	"encoding/json"
	"sync"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/jorgeAM/events/eventService/models"
)

// constants
const (
	DATABASE        = "eventsDB"
	EVENTCOLLECTION = "event"
)

var (
	once sync.Once
	db   *dbLayer
	e    error
)

// DBLayer struct handle mongoDB connection
type dbLayer struct {
	conn *mgo.Session
}

// NewMongoDBLayer initialze dbLayer struct
func NewMongoDBLayer(url string) (*dbLayer, error) {
	once.Do(func() {
		session, err := mgo.Dial(url)
		db.conn = session
		e = err
	})

	return db, e
}

func (db *dbLayer) CreateEvent(event *models.Event) ([]byte, error) {
	ss := db.conn.Clone()
	defer ss.Close()

	if !event.ID.Valid() {
		event.ID = bson.NewObjectId()
	}

	err := ss.DB(DATABASE).C(EVENTCOLLECTION).Insert(event)

	if err != nil {
		return nil, err
	}

	return json.Marshal(event)
}

func (db *dbLayer) ListAvailableEvents() ([]models.Event, error) {
	ss := db.conn.Clone()
	defer ss.Close()

	var events []models.Event
	query := bson.M{"Available": true}
	err := ss.DB(DATABASE).C(EVENTCOLLECTION).Find(query).All(&events)

	if err != nil {
		return nil, err
	}

	return events, nil
}

func (db *dbLayer) GetEvent(id string) (*models.Event, error) {
	ss := db.conn.Clone()
	defer ss.Close()

	var event models.Event
	err := ss.DB(DATABASE).C(EVENTCOLLECTION).FindId(bson.ObjectIdHex(id)).One(&event)

	if err != nil {
		return nil, err
	}

	return &event, nil
}
