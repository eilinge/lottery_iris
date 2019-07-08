package main

/*
curl http://localhost:8888/
curl --data "users=eilinge, duzi, lin" http://localhost:8888/import
curl http://localhost:8888/lucky
*/
import (
	"fmt"
	"math/rand"
	"strings"
	"sync"
	"time"

	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

type lotteryController struct {
	Ctx iris.Context
}

var mu sync.Mutex // 保证线程安全: 多个线程修改共享变量, 导致结果不一致

var userList []string

func NewApp() (app *iris.Application) {
	app = iris.New()
	mvc.New(app.Party("/")).Handle(&lotteryController{})
	return
}
func main() {
	app := NewApp()
	userList = []string{}
	mu = sync.Mutex{}
	app.Run(iris.Addr(":8888"))
}

func (c *lotteryController) Get() string {
	count := len(userList)
	return fmt.Sprintf("\nonline person num: %d\n", count)
}

func (c *lotteryController) PostImport() string {
	strUsers := c.Ctx.FormValue("users")
	users := strings.Split(strUsers, ",")
	mu.Lock()
	defer mu.Unlock()

	count1 := len(userList)
	for _, u := range users {
		u = strings.TrimSpace(u)
		if len(u) > 0 {
			userList = append(userList, u)
		}
	}

	count2 := len(userList)
	return fmt.Sprintf("\nonline person sum:%d, import user successfully:%d", count1, (count2 - count1))
}

func (c *lotteryController) GetLucky() string {
	mu.Lock()
	mu.Unlock()
	count := len(userList)
	if count > 1 {
		seed := time.Now().UnixNano()
		index := rand.New(rand.NewSource(seed)).Int31n(int32(count))
		user := userList[index]
		userList = append(userList[0:index], userList[index+1:]...)
		return fmt.Sprintf("\nthe lucky man: %s, other person num: %d", user, len(userList)-1)
	} else if count == 1 {
		userList = userList[0:0]
		return fmt.Sprintf("\nthe lucky man: %s, other person num: %d", userList[0], count)
	} else {
		return fmt.Sprintf("\nthe userlist in null, please import users by get /import")
	}

}
