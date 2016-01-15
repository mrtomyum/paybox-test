package mqtt
import (
	"testing"
	//"os"
	"fmt"
	"strings"
)

func Test_print_Json_Single_Record(t *testing.T){
	mqMsg := sub("#")
	//os.Stdout.Write(mqMsg)
	fmt.Println(mqMsg)
	if !strings.Contains(mqMsg,"SiteID") {
		t.Fatal("Error can not find SiteID Field in Return val ",mqMsg)
	}


//	if mqMsg.SiteID != 1 {
//		t.Fatal("error ",mqMsg)
//	}
}