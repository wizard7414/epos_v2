package deviant

import (
	"github.com/PuerkitoBio/goquery"
	"strings"
)

func FindSelectionByHookAndClass(document *goquery.Document, tag string, hook string, class string) (*goquery.Selection, bool) {
	var target *goquery.Selection

	document.Find(tag).Each(func(i int, selection *goquery.Selection) {
		hookOfNode, _ := selection.Attr("data-hook")
		classOfNode, _ := selection.Attr("class")
		if hookOfNode == hook && strings.Contains(classOfNode, class) {
			target = selection
		}
	})

	return target, true
}

func FindSelectionByClass(document *goquery.Document, tag string, class string) (*goquery.Selection, bool) {
	var target *goquery.Selection

	document.Find(tag).Each(func(i int, selection *goquery.Selection) {
		classOfNode, _ := selection.Attr("class")
		if strings.Contains(classOfNode, class) {
			target = selection
		}
	})

	return target, true
}

func FindSelectionArrayByClass(document *goquery.Document, tag string, class string) ([]*goquery.Selection, bool) {
	var target []*goquery.Selection

	document.Find(tag).Each(func(i int, selection *goquery.Selection) {
		classOfNode, _ := selection.Attr("class")
		if strings.Contains(classOfNode, class) {
			target = append(target, selection)
		}
	})

	return target, true
}

func FindSelectionByMeta(document *goquery.Document, property string) (*goquery.Selection, bool) {
	var target *goquery.Selection

	document.Find("meta").Each(func(i int, selection *goquery.Selection) {
		propertyOfNode, _ := selection.Attr("property")
		if propertyOfNode == property {
			target = selection
		}
	})

	return target, true
}
