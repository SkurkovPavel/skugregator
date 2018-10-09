package skuParser

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"sync"
)

const req = 4

var Sites = map[string]string{
	"ozone.ru":      "http://static.ozone.ru/multimedia/yml/facet/div_soft.xml",
	"trenazhery.ru": "http://www.trenazhery.ru/market2.xml",
	"radio-liga.ru": "http://www.radio-liga.ru/yml.php",
	"armprodukt.ru": "http://armprodukt.ru/bitrix/catalog_export/yandex.php",
}

type bodyArr struct {
	name string
	val  string
}

func SkuPars() string {

	var wg sync.WaitGroup

	OnenSkuDB()

	defer CloseSkuDB(SkuDB)

	for site, url := range Sites {
		wg.Add(1)

		go skuReguest(&wg, site, url)
	}

	wg.Wait()
	return "finito"

}

//Выполняем запросы
func skuReguest(wg *sync.WaitGroup, site, url string) {
	var res = make(map[string]string)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprint(os.Stderr, "Ошибка при запросе данных\n")
		os.Exit(1)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprint(os.Stderr, "Ошибка при чтении данных\n")
		os.Exit(1)
	}

	resp.Body.Close()

	if resp.StatusCode != 200 {
		fmt.Fprintf(os.Stderr, "Ошибка при получении данных с сервера: код %s", resp.Status)
		os.Exit(1)
	}
	res["body"] = string(body)
	res["site"] = string(site)

	SaveData(res)

	wg.Done()
}
