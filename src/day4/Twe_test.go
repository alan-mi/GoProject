package day4
import (
	"os"
	"testing"
	"time"
	"fmt"
)


func TestB(t *testing.T) {
	fmt.Println(os.Args)
	if os.Args[len(os.Args)-1] == "b" {
		t.Parallel()
	}
	time.Sleep(time.Second * 2)
}
func TestA(t *testing.T) {
	t.Parallel()
	time.Sleep(time.Second * 2)

}
