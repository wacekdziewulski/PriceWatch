package strategy

import (
	"PriceWatch/model"
	"encoding/json"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
	"golang.org/x/net/html"
)

type banggoodJSONSeller struct {
	Name string `json:"name"`
}

type banggoodJSONOffer struct {
	PriceCurrency string             `json:"priceCurrency"`
	Price         string             `json:"price"`
	Seller        banggoodJSONSeller `json:"seller"`
}

type banggoodJSONLd struct {
	ImageURL string              `json:"image"`
	Offers   []banggoodJSONOffer `json:"offers"`
	Name     string              `json:"name"`
}

// JSONLdPriceStrategy defines the strategy for processing Json+LD data
type JSONLdPriceStrategy struct {
}

// NewJSONLdPriceStrategy returns the JSONLdPriceStrategy
func NewJSONLdPriceStrategy() *JSONLdPriceStrategy {
	return &JSONLdPriceStrategy{}
}

func (p *JSONLdPriceStrategy) findChildOfNodeWithName(root *html.Node, name string, attributeMatch bool) *html.Node {
	for node := root.FirstChild; node != nil; node = node.NextSibling {
		if node.Type == html.ElementNode {
			logrus.Debugf("Looking at node: %s, %s", node.Type, node.Data)
			if node.Data == name {
				if attributeMatch == true && !p.findAttributeTypeWithValue(node.Attr, "type", "application/ld+json") {
					logrus.Debugf("Attribute match requested, but not found: %s", node.Attr)
					continue
				}
				logrus.Debugf("Found child of node: %s", name)
				return node.FirstChild
			}
		}
	}
	logrus.Infof("Failed to find node with name: %s", name)
	return nil
}

func (p *JSONLdPriceStrategy) findJSONLdNode(root *html.Node) string {
	// find html
	node := p.findChildOfNodeWithName(root, "html", false)
	node = p.findChildOfNodeWithName(node, "script", true)

	return node.Data
}

func (p *JSONLdPriceStrategy) findAttributeTypeWithValue(attributes []html.Attribute, attrName string, attrValue string) bool {
	for _, attr := range attributes {
		if attr.Key == attrName && attr.Val == attrValue {
			return true
		}
	}
	return false
}

// FillPriceData extracts the prices from Json+LD object - implements PriceStrategy
func (p *JSONLdPriceStrategy) FillPriceData(pageBody string, priceData *model.PriceData) {
	doc, err := html.Parse(strings.NewReader(pageBody))
	if err != nil {
		logrus.Warnf("Failed to parse HTML looking for json+ld because of: %+v", err)
		return
	}

	jsonLdContent := p.findJSONLdNode(doc)

	logrus.Infof("Found Json+LD content: %s", jsonLdContent)
	jsonLdContent = strings.Trim(jsonLdContent, "; \n\t\r")
	logrus.Debugf("Trimmed Json+LD content: %s", jsonLdContent)

	var banggoodJSON banggoodJSONLd

	if err := json.Unmarshal([]byte(jsonLdContent), &banggoodJSON); err != nil {
		logrus.Warnf("Failed to process JSON+LD, because: %+v", err)
	}

	priceData.Title = banggoodJSON.Name
	priceData.Site = banggoodJSON.Offers[0].Seller.Name
	priceData.ImageURL = banggoodJSON.ImageURL
	priceData.PriceCurrency = banggoodJSON.Offers[0].PriceCurrency
	priceData.PriceAmount, _ = strconv.ParseFloat(banggoodJSON.Offers[0].Price, 32)
}
