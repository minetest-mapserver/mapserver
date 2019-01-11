package mapblockrenderer

import (
  "github.com/sirupsen/logrus"
)

var log *logrus.Entry
func init(){
	log = logrus.WithFields(logrus.Fields{"prefix": "mapblockrenderer"})
}
