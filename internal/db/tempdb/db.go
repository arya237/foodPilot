package tempdb

import (
	"sync"

	"github.com/arya237/foodPilot/internal/models"
)

type FakeDb struct {
	UserMu      sync.RWMutex
	FoodMu      sync.RWMutex
	RateMu      sync.RWMutex
	Users       map[int]*models.User
	Foods       map[int]*models.Food
	Rates       map[int]map[int]*models.Rate
	FoodCounter int
	UserCounter int
}

func NewDb(cnf *Config) *FakeDb {
	users := map[int]*models.User{
		0: {
			Id: 0,
			Username: cnf.AdminUsername,
			Password: cnf.AdminUsername,
			AutoSave: true,
			Role: models.RoleAdmin,
		},
	}

	foods := map[int]*models.Food{
		1:  {Name: "چلو کباب کوبیده زعفرانی", Id: 1},
		3:  {Name: "خوراک گوشت چرخ‌کرده با سیب زمینی", Id: 3},
		4:  {Name: "چلو خورشت آلو با اسفناج", Id: 4},
		5:  {Name: "چلو جوجه کباب", Id: 5},
		6:  {Name: "فرنی", Id: 6},
		7:  {Name: "لوبیا پلو", Id: 7},
		8:  {Name: "خوراک فلافل", Id: 8},
		9:  {Name: "خوراک فیله سوخاری", Id: 9},
		10: {Name: "چلو خورشت قیمه", Id: 10},
		11: {Name: "کلم پلو", Id: 11},
		12: {Name: "خوراک عدسی", Id: 12},
		13: {Name: "زرشک پلو با مرغ", Id: 13},
		14: {Name: "شکلات صبحانه ،شیر،پنیر ،چای", Id: 14},
		15: {Name: "تخم مرغ(24عدد)", Id: 15},
		16: {Name: "پنیر ،مربا،خامه،چای", Id: 16},
		17: {Name: "تخم مرغ(14 عدد)", Id: 17},
		18: {Name: "شیرموز(2عدد)، شیر کاکائو(2عدد)،کیک(4عدد)،چایی(5عدد)", Id: 18},
		19: {Name: "غلات صبحانه، شیر", Id: 19},
		20: {Name: "خوراک پاستا", Id: 20},
		21: {Name: "چلو ماهی قزل آلا", Id: 21},
		22: {Name: "خوراک ناگت مرغ", Id: 22},
		23: {Name: "استانبولی پلو", Id: 23},
		24: {Name: "چلو تن ماهی", Id: 24},
		25: {Name: "عدس پلو", Id: 25},
		26: {Name: "چلو خورشت قورمه سبزی", Id: 26},
		27: {Name: "خوراک دلمه", Id: 27},
		28: {Name: "چلو کباب کوبیده مرغ", Id: 28},
		29: {Name: "خوراک شنیسل مرغ", Id: 29},
		30: {Name: "خوراک لوبیا + پوره", Id: 30},
		31: {Name: "سوپ جو", Id: 31},
		32: {Name: "چلو خورشت بادمجان", Id: 32},
	}
	return &FakeDb{
		Users: users,
		Foods: foods,

		Rates:       map[int]map[int]*models.Rate{},
		FoodCounter: len(foods),
		UserCounter: len(users),
	}
}
