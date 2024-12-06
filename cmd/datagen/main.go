package main

import (
	"fmt"
	"log"
	"math/rand/v2"
	"os"
	"time"
)

func getNewName(s []byte) []byte {
	n := len(s)

	i := n - 1
	for ; i >= 0 && s[i] == 'z'; i-- {
		s[i] = 'a'
	}

	if i < 0 {
		return []byte("aaa")
	}

	s[i]++
	return s
}

func getRandomDate() string {
	randomDate := time.Now().Add(time.Duration(-rand.IntN(35000) * int(time.Hour))).Add(time.Duration(-rand.IntN(3600) * int(time.Second)))

	return randomDate.Format("2006-01-02 15:04:05")
}

func main() {

	mp := make(map[int][]int)
	// Restaraunt and Food Data
	restarauntfile, err := os.Create("restaraunt.csv")
	if err != nil {
		log.Fatal("error creating restaraunt file")
	}

	foodfile, err := os.Create("food.csv")
	if err != nil {
		log.Fatal("error creating food file")
	}

	foodList := make([]string, 50)
	foodList = []string{
		"margaritha pizza", "pepperoni pizza", "bbq chicken pizza", "chicken pesto pizza", "paneer pizza", "corn pizza", "capsicum pizza", "chicken extravaganza pizza", "veg farmhouse pizza", "chicken dominator pizza", "tomato pizza", "mushroom pizza", "chicken corn pizza", "paneer corn pizza", "mushroom corn pizza",

		"chicken burger", "chicken filllet burger", "chicken bbq burger", "chicken strip burger", "veggie burger", "aloo tikki burger", "lamb burger", "paneer burger", "capsicum corn burger", "corn cheese burger",

		"chicken sandwich", "paneer sandwich", "capsicum corn sandwich", "corn cheese sandwich", "chilli sandwich", "mushroom sandwich", "chicken corn sandwich", "paneer corn sandwich", "mushroom corn sandwich",

		"chicken lollipop biriyani", "chicken biriyani", "mutton biriyani",

		"chocolate milkshake", "vanilla milkshake", "sweet lassi", "strawberry lassi", "chocolate mouse", "gulab jamun", "coke", "pepsi", "7up", "sprite", "mirinda", "chocolate ice cream", "vanilla ice cream",
	}

	foodindex := 1
	rname := []byte("aaa")
	for i := 1; i <= 10000; i++ {
		data := fmt.Sprintf("%v,%v\n", i, string(rname))
		restarauntfile.Write([]byte(data))

		st := rand.IntN(30)
		sz := rand.IntN(20)
		mp[i] = make([]int, max(5, sz))
		k := 0
		for j := st; j < st+max(5, sz); j++ {
			foodData := fmt.Sprintf("%v,%v,%v\n", foodindex, foodList[j], i)
			foodfile.Write([]byte(foodData))
			mp[i][k] = foodindex
			k++
			foodindex++
		}

		rname = getNewName(rname)
	}

	restarauntfile.Close()
	foodfile.Close()

	fmt.Println("Restaraunt and Food files created")

	fmt.Println(mp[589])

	// User
	userfile, err := os.Create("user.csv")
	if err != nil {
		log.Fatal("error creating user file")
	}

	for i := 1; i <= 10000; i++ {
		data := fmt.Sprintf("%v,%v\n", i, string(rname))
		userfile.Write([]byte(data))
		rname = getNewName(rname)
	}

	userfile.Close()
	fmt.Println("User file created")

	// Order and Ordered Food
	orderFile, err := os.Create("order.csv")
	if err != nil {
		log.Fatal("error creating order file")
	}

	orderedFoodfile, err := os.Create("orderedFood.csv")
	if err != nil {
		log.Fatal("error creating ordered food file")
	}

	metroCityList := []string{"Bangalore", "Delhi", "Chennai", "Hyderbad", "Gurugram", "Kolkata", "Mumbai"}
	otherCityList := []string{"dfdfd", "dfdfdf", "dfgrgt", "dyfj", "euhj", "ef", "diush", "ifji", "ifjds", "djfn", "udhsm", "dfdsfds", "ufhndsuf", "diohfnds", "fduisbniu", "udiyhj", "dufhjn", "dufhdj", "dufohjn", "dyfhvn", "duyfhn", "duhfnc", "gfhvb", "tofk", "tofkj", "eosk", "kjx", "ygh", "urdj", "trhf", "eiksl", "iejf", "kdjfnc", "yudjsk", "kmnbf",
		"ejrrjn", "osiducv", "xociuv", "mwne", "menrtb", "osdicu", "oikjedjc", "ruhfndjf", "ruhfjs", "uryhfnmki", "iruyt", "ruhfnmjc", "oiuyhnm", "aa", "aaa"}

	orderedFoodIndex := 1
	for i := 1; i <= 1000000; i++ {
		rid := rand.IntN(10000) + 1
		x := rand.IntN(7)
		var city string
		if x < 8 {
			city = metroCityList[x]
		} else {
			city = otherCityList[rand.IntN(50)]
		}
		data := fmt.Sprintf("%v,%v,%v,%v,%v,%v\n", i, rid, rand.IntN(10000)+1, city, getRandomDate(), 100+rand.IntN(400))
		orderFile.Write([]byte(data))

		for j := 0; j < 1+rand.IntN(5); j++ {
			foodData := fmt.Sprintf("%v,%v,%v,%v\n", orderedFoodIndex,
				mp[rid][rand.IntN(len(mp[rid]))], i, city)
			orderedFoodfile.Write([]byte(foodData))
			orderedFoodIndex++
		}

	}

	orderFile.Close()
	orderedFoodfile.Close()
	fmt.Println("Order and Ordered Food files created")

}
