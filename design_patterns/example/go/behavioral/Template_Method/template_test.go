package Template_Method

import (
	"chenyx/design_patterns/example/go/behavioral/Template_Method/bad"
	"chenyx/design_patterns/example/go/behavioral/Template_Method/good"
	"log"
	"testing"
)

func TestTemplateMethod(t *testing.T) {
	log.Println("-----------")
	bad.MainRun()
	log.Println("-----------")
	good.MainRun()
	log.Println("-----------")
}
