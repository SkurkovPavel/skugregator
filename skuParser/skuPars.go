package skuParser


import (
	db  "../alias"

	"flag"
	"fmt"
	"net/http"
	"os"
	"regexp"
	"strings"
)

//Все что связано с запросами и парсером
func skuParser(){

}
//подтягиваем настройки из скуджисона
func skuGetSetings(){

}
//формируем очередь запросов
func skuSetQueue()  {

}
//Выполняем запросы
func SkuReguest()  {

}
//Обрабатываем полученные ответы
func skuProcessingResults(){
	
}
//Отдаем в скугрегатор
func SkuSkugregatorInit()  {
	
}
//Ставим счетчик на час. После которого снова отправим запросы. такой себе демон
func SkuSetTimer( time int) error{
	return nil
}
//Отправляем запросы вручную. Для мануального рестарта парсера
func SkugregatorManualInit()  {

	fmt.Println("Слушаю")
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		request.ParseForm()

		fmt.Fprintf(writer,"Ответ из GO %v",request.Form)
	})
	http.ListenAndServe(":3000",nil)

}

var url = "http://www.zakupki.gov.ru/epz/order/notice/ea44/view/common-info.html?regNumber="

var fields = map[string]string{
	"number": "", "method": "", "platform": "", "object": "", "stage": "", "date": "", "nmc": "",
}

var regEx = map[string]string{
	"date":     `(?is)Размещено:.+?(\d{2}.\d{2}.\d{4})`,
	"method":   `(?is)Способ.+?определен.+?поставщик.+?<td>(.+?)<\/td>`,
	"platform": `(?is)Наименован.+?электрон.+?площадк.+?<td>(.+?)<\/td>`,
	"object":   `(?is)Наименован.+?объект.+?закупк.+?<td>(.+?)<\/td>`,
	"stage":    `(?is)Этап.+?закупк.+?<td>(.+?)<\/td>`,
	"nmc":      `(?is)Начальн.+?цен.+?контракт.+?<td>(.+?)</td>`,
}

var law = flag.Int("f", 44, "Федеральный закон")
var number = flag.String("n", "", "Номер извещения")

func main() {
	flag.Parse()

	if len(*number) == 0 {
		fmt.Fprint(os.Stderr, "Не указан номер извещения\n")
		os.Exit(1)
	}

	if match, _ := regexp.MatchString("^[0-9]{19}$", *number); match != true {
		fmt.Fprint(os.Stderr, "Некорректный номер извещения\n")
		os.Exit(1)
	}

	data, _ := GetData(*number)
	if len(data) == 0 {
		body, err := request.Send(url + *number)
		if err != nil {
			fmt.Fprint(os.Stderr, err, "\n")
			//os.Exit(1)
		}

		fillStructure(body)

		err = SaveData(fields)
		if err != nil {
			fmt.Fprint(os.Stderr, "Произошла ошибка при сохранении данных в базе\n")
			fmt.Println(err)
			os.Exit(1)
		}

		viewInfo(fields)
	} else {
		viewInfo(data)
	}
}

// Обработка данных и заполнение структуры
func fillStructure(body string) {
	fields["number"] = *number

	for k, v := range regEx {
		re := regexp.MustCompile(v)
		match := re.FindStringSubmatch(body)
		if len(match) != 0 {
			fields[k] = strings.TrimSpace(match[1])
		}
	}
}

// Вывод данных
func viewInfo(fields map[string]string) {
	fmt.Print(strings.Repeat("▬", 100) + "\n")
	fmt.Printf("Дата публикации: %s\n", fields["date"])
	fmt.Printf("Номер извещения: %s\n", fields["number"])
	fmt.Printf("Способ определения поставщика: %s\n", fields["method"])
	fmt.Printf("Наименование электронной площадки: %s\n", fields["platform"])
	fmt.Printf("Наименование объекта закупки: %s\n", fields["object"])
	fmt.Printf("Цена контракта: %s\n", fields["nmc"])
	fmt.Printf("Этап закупки: %s\n", fields["stage"])
	fmt.Print(strings.Repeat("▬", 100))
}




