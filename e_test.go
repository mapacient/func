package e 
import (
	"testing"
	"sort"
	"math/rand"
	"time"
)

func Test_e(t *testing.T) {

	n:=77
	osl:=make([]int,n)
	for i:=0;i<n;i++{
		osl[i]=i+1
	}

	rand.Seed(time.Now().UnixNano())
	for i:=0;i<n;i++{
		sel:=rand.Intn(n-i)
		osl[i],osl[sel]=osl[sel],osl[i]
	}

	t.Log(osl)


	sorted := make([][]int, 7)
	for i := 0; i < 7; i++ {
		sl := make([]int, 11)
		for l := 0; l < 11; l++ {
			sl[l] = osl[i*11+l]
		}
		sort.Ints(sl)
		sorted[i] = sl
		
	}
	sorted[2]=append(sorted[2],70)

	r := e(sorted, 17)
	t.Log(r)
	r = e(nil, 0)
	t.Log(r)
	for i:=1;i<78;i++{
	r = e(sorted, i)
		t.Log(i,r)
}
	t.Log(sorted[0])
	t.Log(sorted[1])
	t.Log(sorted[2])

}
