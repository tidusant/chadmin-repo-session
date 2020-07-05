package session

import (
	"time"

	c3mcommon "github.com/tidusant/chadmin-common"
	"github.com/tidusant/chadmin-log"
	"github.com/tidusant/chadmin-repo-models"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func SaveFBWhook(whook models.FBWhook) string {

	col := db.C("addons_fbwhook")

	err := col.Insert(whook)
	log.Debugf("SaveFBWhook done")
	c3mcommon.CheckError("SaveFBWhook", err)
	return "1"
}
func SaveWhook(whook models.Whook) string {

	col := db.C("addons_whook")
	whook.Created = int(time.Now().Unix())
	err := col.Insert(whook)
	c3mcommon.CheckError("whook insert", err)
	return "1"
}
func GetWhookByLabel(label string) []models.Whook {
	var rs []models.Whook
	col := db.C("addons_whook")

	err := col.Find(bson.M{"data": bson.M{"$regex": bson.RegEx{label, "si"}}}).All(&rs)
	c3mcommon.CheckError("GetWhook by label", err)
	return rs
}
func GetWhook() models.Whook {

	col := db.C("addons_whook")
	var bs models.Whook
	change := mgo.Change{
		Update:    bson.M{"$set": bson.M{"status": 1, "modified": time.Now().Unix()}},
		ReturnNew: true,
	}
	_, err := col.Find(bson.M{"status": 0}).Apply(change, &bs)
	c3mcommon.CheckError("GetWhook script", err)
	return bs
}

func GetFBWhook() models.FBWhook {

	col := db.C("addons_fbwhook")
	var bs models.FBWhook
	change := mgo.Change{
		Update:    bson.M{"$set": bson.M{"status": 1, "modified": time.Now().Unix()}},
		ReturnNew: true,
	}
	_, err := col.Find(bson.M{"status": bson.M{"$ne": 1}}).Apply(change, &bs)
	c3mcommon.CheckError("GetFBWhook script", err)
	return bs
}

func SaveFBComment(fbcomment models.FBComment) string {

	col := db.C("addons_fbcomment")
	err := col.Insert(fbcomment)
	c3mcommon.CheckError("SaveFBComment insert", err)
	return "1"
}
func SaveFBConversation(fbcon models.FBConversation) string {

	col := db.C("addons_fbconversation")
	err := col.Insert(fbcon)
	c3mcommon.CheckError("SaveFBConversation insert", err)
	return "1"
}
