package chronos

import (
	"github.com/dollarkillerx/chronos/adapter/redis_adapter"
	"log"
	"testing"
	"time"
)

var enf *Chronos

func init() {
	log.SetFlags(log.Llongfile | log.LstdFlags)

	adapter := redis_adapter.New("127.0.0.1:6379", 10)
	enforcer, err := NewEnforcer("./exp/base.conf", adapter)
	if err != nil {
		log.Fatalln(err)
	}
	enf = enforcer
}

func TestAllTotal(t *testing.T) {
	policy, err := enf.AddPolicy("Person", "0_advanced_search_times", "r.sub.Count<10")
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(policy)

	filteredPolicy, err := enf.GetFilteredPolicy("Person", "0_advanced_search_times")
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(filteredPolicy)

	abac := ABACEnforce{
		Name:  "Person",
		Count: 2,
	}

	enforce, err := enf.Enforce(abac, "0_advanced_search_times")
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(enforce)

	abac = ABACEnforce{
		Name:  "Person",
		Count: 20,
	}

	enforce, err = enf.Enforce(abac, "0_advanced_search_times")
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(enforce)

	policy, err = enf.AddPolicy("test_team_0", "RimePro", "r.sub.Time<1635936065")
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(policy)

	filteredPolicy, err = enf.GetFilteredPolicy("test_team_0", "RimePro")
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(filteredPolicy)

	abac = ABACEnforce{
		Name: "test_team_0",
		Time: time.Now().Unix(),
	}

	enforce, err = enf.Enforce(abac, "RimePro")
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(enforce)

	abac = ABACEnforce{
		Name: "test_team_0",
		Time: time.Now().UnixNano(),
	}

	enforce, err = enf.Enforce(abac, "RimePro")
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(enforce)
}

func TestAdd(t *testing.T) {
	policy, err := enf.AddPolicy("Person", "0_advanced_search_times", "r.sub.Count<10")
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(policy)

	filteredPolicy, err := enf.GetFilteredPolicy("Person", "0_advanced_search_times")
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(filteredPolicy)
}

type ABACEnforce struct {
	Name  string
	Time  int64
	Count int64
}

func TestAb(t *testing.T) {
	abac := ABACEnforce{
		Name:  "Person",
		Time:  time.Now().Unix(),
		Count: 2,
	}

	enforce, err := enf.Enforce(abac, "0_advanced_search_times")
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(enforce)
}
