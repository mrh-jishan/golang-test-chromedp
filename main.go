// Command logic is a chromedp example demonstrating more complex logic beyond
// simple actions.
package main

import (
	"bufio"
	"context"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/chromedp/chromedp"
)

// ud contains a url, description for a project.
type ud struct {
	URL, Description string
}

type marcent struct {
	name, location, contact string
}

func main() {

	x_path_product := `.//*[contains(concat(" ",normalize-space(@class)," ")," ui ")][contains(concat(" ",normalize-space(@class)," ")," vertical ")][contains(concat(" ",normalize-space(@class)," ")," segment ")][contains(concat(" ",normalize-space(@class)," ")," page ")][contains(concat(" ",normalize-space(@class)," ")," no ")][contains(concat(" ",normalize-space(@class)," ")," padding ")][contains(concat(" ",normalize-space(@class)," ")," borderless ")]`
	open_browser(x_path_product)
}

func visit_page(url string) marcent {
	var res marcent
	return res
}

func open_browser(x_path string) {
	var err error

	// create context
	ctxt, cancel := context.WithCancel(context.Background())
	defer cancel()

	// create chrome instance
	c, err := chromedp.New(ctxt, chromedp.WithLog(log.Printf))
	if err != nil {
		log.Fatal(err)
	}

	file_name := `result.csv`
	res := read_csv(file_name)

	var save_data_marcent []marcent

	for i := 0; i < len(res); i++ {
		fmt.Println(res[i].URL)
		res, err := receive_scrap_data(ctxt, c, x_path, res[i].URL)
		if err != nil {
			log.Fatalf("could not list awesome go projects: %v", err)
		}

		fmt.Println(res)
		save_data_marcent = append(save_data_marcent, res)

		//var save_data [][]string
		//
		//for _, value := range res {
		//	save_data = append(save_data, []string{value.URL, value.Description})
		//	write_into_file(save_data)
		//}
		//
		//write_into_file(save_data)

	}
	var save_data [][]string
	for _, value := range save_data_marcent {
		save_data = append(save_data, []string{value.name, value.location})
		write_into_file(save_data)
	}

	write_into_file(save_data)

	// list awesome go projects for the "Selenium and browser control tools."
	//res, err := receive_scrap_data(ctxt, c, x_path, url)
	//if err != nil {
	//	log.Fatalf("could not list awesome go projects: %v", err)
	//}
	//
	//fmt.Println(res)

	//var save_data [][]string
	//
	//for _, value := range res {
	//	save_data = append(save_data, []string{value.URL, value.Description})
	//	write_into_file(save_data)
	//}
	//
	//write_into_file(save_data)

	// shutdown chrome
	err = c.Shutdown(ctxt)
	if err != nil {
		log.Fatal(err)
	}

	// wait for chrome to finish
	err = c.Wait()
	if err != nil {
		log.Fatal(err)
	}

	//// output the values
	//for k, v := range res {
	//	log.Printf("project %s (%s): '%s'", k, v.URL, v.Description)
	//}
}

//scrap the page
func receive_scrap_data(ctxt context.Context, c *chromedp.CDP, sect string, url string) (m marcent, error error) {
	// force max timeout of 15 seconds for retrieving and processing the data
	var cancel func()
	ctxt, cancel = context.WithTimeout(ctxt, 25*time.Second)
	defer cancel()

	sel := fmt.Sprintf(sect)

	// navigate
	if err := c.Run(ctxt, chromedp.Navigate(url)); err != nil {
		return marcent{}, fmt.Errorf("could not navigate to github: %v", err)
	}

	// wait visible
	if err := c.Run(ctxt, chromedp.WaitVisible(sel)); err != nil {
		return marcent{}, fmt.Errorf("could not get section: %v", err)
	}

	// get merchants name
	var marcants_name string
	if err := c.Run(ctxt, chromedp.Text(sel+`//a[contains(concat(" ",normalize-space(@class)," ")," header ")]`, &marcants_name)); err != nil {
		return marcent{}, fmt.Errorf("could not get links: %v", err)
	}

	fmt.Println(`------------------marcent name----------------------------------`, marcants_name)

	// get merchants name
	var marcants_location string
	if err := c.Run(ctxt, chromedp.Text(sel+`//h3[contains(concat(" ",normalize-space(@class)," ")," is-meta ")]`, &marcants_location)); err != nil {
		return marcent{}, fmt.Errorf("could not get links: %v", err)
	}

	fmt.Println(`-----------------------marcent location-----------------------------`, marcants_location)

	//// get merchants name
	//var marcants_contact string
	//if err := c.Run(ctxt, chromedp.Text(`//*[contains(concat(" ",normalize-space(@class)," ")," ui ")][contains(concat(" ",normalize-space(@class)," ")," text ")][contains(concat(" ",normalize-space(@class)," ")," markdown ")]//ul/li/ul/li[contains(text(),"+")]`, &marcants_contact)); err != nil {
	//	return marcent{}, fmt.Errorf("could not get links: %v", err)
	//}
	//
	//fmt.Println(`-----------------------marcants_contact-----------------------------`, marcants_contact)

	//// get project link Nodes
	//var projects []*cdp.Node
	//if err := c.Run(ctxt, chromedp.Nodes(sel+`//a/h3/text()`, &projects)); err != nil {
	//	return marcent{}, fmt.Errorf("could not get projects: %v", err)
	//}
	//
	//// get links text
	//var links []*cdp.Node
	//if err := c.Run(ctxt, chromedp.Nodes(sel+`//a`, &links)); err != nil {
	//	return marcent{}, fmt.Errorf("could not get links: %v", err)
	//}
	//
	//// get description text
	//var descriptions []*cdp.Node
	//if err := c.Run(ctxt, chromedp.Nodes(sel+`//a/h3/text()`, &descriptions)); err != nil {
	//	return marcent{}, fmt.Errorf("could not get descriptions: %v", err)
	//}

	// process data
	//res := make(map[string]ud)
	//for i := 0; i < len(projects); i++ {
	//	res[projects[i].NodeValue] = ud{
	//		URL:         url + links[i].AttributeValue("href"),
	//		Description: descriptions[i].NodeValue,
	//	}
	//
	//}

	return marcent{name: marcants_name, location: marcants_location, contact: `+5457496554458`}, nil
}

func write_into_file(data [][]string) {
	//Create output file
	file, err := os.Create("result2.csv")
	checkError("Cannot create file", err)
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	err = writer.WriteAll(data)
	checkError("Cannot write to file", err)
}

func checkError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}

func read_csv(file_name string) []ud {
	csvFile, _ := os.Open(file_name)
	reader := csv.NewReader(bufio.NewReader(csvFile))
	var url []ud

	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}
		url = append(url, ud{URL: line[0], Description: line[1]})
	}
	return url
}
