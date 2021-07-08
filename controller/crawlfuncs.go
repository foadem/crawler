package controller

import (
	"fmt"
	"log"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/gocolly/colly"
)

func CrawlTechnolife() {

	es, err := elasticsearch.NewDefaultClient()
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}

	res, err := es.Info()
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}

	defer res.Body.Close()
	log.Println(res)

	c := colly.NewCollector()
	c.IgnoreRobotsTxt = false

	// product := models.Product{}

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
		return
	})

	// Crawl Products Category
	c.OnHTML(".ProductPrlist_product__3oA2g", func(e *colly.HTMLElement) {
		fmt.Println("product_link")
		link := "https://www.technolife.ir" + e.ChildAttr("a", "href")
		fmt.Println(link)
		e.Request.Visit(link)
	})

	// product_title_fa
	// product_title-en
	c.OnHTML(".product_productInfo__3Vkg3", func(e *colly.HTMLElement) {
		fmt.Println("product_title_fa")
		fmt.Println(e.ChildText("h1"))
		fmt.Println("product_title_en")
		fmt.Println(e.ChildText(".product_productHeader__HjqBl h2"))
	})

	// product_description
	// product_specifications
	c.OnHTML(".product_tInfo__3BP7K li", func(e *colly.HTMLElement) {
		fmt.Println("product_specifications")
		fmt.Println(e.Text)
	})

	// seller_name = technolife
	// seller_score = null
	// product_score = null
	// warranty_details
	c.OnHTML(".product_guaranteeInfo__2KwC-", func(e *colly.HTMLElement) {
		fmt.Println("warranty_details")
		fmt.Println(e.Text)
	})

	// warranty_details = null
	// availability_status = true
	// product_brand, sub_category
	c.OnHTML(".product_brandName__2O3bl .product_sameTransition__2Kr1J", func(e *colly.HTMLElement) {
		if e.Index == 0 {
			fmt.Println("product_brand")
			fmt.Println(e.Text)
		} else if e.Index == 1 {
			fmt.Println("sub_category")
			fmt.Println(e.Text)
		}
	})
	// main_category = consumer electronics
	// product_encoded_link = null
	// product_image_link
	c.OnHTML(".product_window__6C8HU", func(e *colly.HTMLElement) {
		fmt.Println("product_image_link")
		fmt.Println(e.ChildAttr("img", "src"))
	})
	// previous-price = null
	// current_price
	c.OnHTML(".product_productPrice__jgJIR .product_sameTransition__2Kr1J", func(e *colly.HTMLElement) {
		fmt.Println("current_price")
		fmt.Println(e.Text)
	})
	// color
	c.OnHTML(".product_selectedColor__eJzRT .product_sameTransition__2Kr1J", func(e *colly.HTMLElement) {
		if e.Index == 1 {
			fmt.Println("color")
			fmt.Println(e.Text)
		}
	})
	// in_stock_count = null

	// for i := 1; i < 5; i++ {
	// 	c.Visit("https://www.digikala.com/search/" + categoryLink + "/?has_selling_stock=1&pageno=" + strconv.Itoa(i) + "&sortby=1")
	// 	// fmt.Println("https://www.digikala.com/search/" + categoryLink + "/?has_selling_stock=1&pageno=" + strconv.Itoa(i) + "&sortby=1")
	// 	c.Wait()
	// }
	c.Visit("https://www.technolife.ir/product/list/69_800_801/%D8%AA%D9%85%D8%A7%D9%85%DB%8C-%DA%AF%D9%88%D8%B4%DB%8C%E2%80%8C%D9%87%D8%A7?keywords=only_available=true")
	c.Wait()
}
