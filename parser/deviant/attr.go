package deviant

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/wizard7414/epos_v2/domain/miner"
	"strings"
)

func FindObjectAttributes(document *goquery.Document) []miner.AttrV {
	var resultAttributes []miner.AttrV

	entryTitleSelection, titleResult := FindSelectionByMeta(document, "og:title")
	if titleResult {
		title, _ := entryTitleSelection.Attr("content")
		author, authorResult := GetAuthorFromTitle(title)
		if authorResult {
			resultAttributes = append(resultAttributes, author)
		}
	}

	tagsSelection, tagsResult := FindSelectionArrayByClass(document, "span", "_3uQxz")
	if tagsResult {
		tags, tagsAttributeResult := FindTags(tagsSelection)
		if tagsAttributeResult {
			resultAttributes = append(resultAttributes, tags)
		}
	}

	infoSelection, infoResult := FindSelectionByClass(document, "div", "_19ms0")
	if infoResult {
		info, infoAttributeResult := FindInfo(infoSelection)
		if infoAttributeResult {
			resultAttributes = append(resultAttributes, info)
		}
	}

	graphicsSelection, graphicsResult := FindSelectionByClass(document, "img", "_1izoQ")
	if graphicsResult {
		graphics, graphicsAttributeResult := FindGraphics(graphicsSelection)
		if graphicsAttributeResult {
			resultAttributes = append(resultAttributes, graphics)
		}
	}

	return resultAttributes
}

func GetAuthorFromTitle(title string) (miner.AttrV, bool) {
	var author miner.AttrV

	author.ID = 0
	author.AttrType = miner.AttrTypeV{
		ID:    0,
		Title: "string",
	}
	author.Code = miner.AttrCodeV{
		ID:    0,
		Title: "author",
	}

	titleArray := strings.Split(title, " ")
	for i := range titleArray {
		if titleArray[i] == "by" && (i+1 < len(titleArray)) {
			author.Value = titleArray[i+1]
		}
	}

	return author, author.Value != ""
}

func FindTags(tagsSelection []*goquery.Selection) (miner.AttrV, bool) {
	var tags miner.AttrV
	var tagsArray []string

	tags.ID = 0
	tags.AttrType = miner.AttrTypeV{
		ID:    0,
		Title: "list",
	}
	tags.Code = miner.AttrCodeV{
		ID:    0,
		Title: "tags",
	}

	for tag := range tagsSelection {
		if len(tagsSelection[tag].Nodes) == 1 {
			tagsArray = append(tagsArray, tagsSelection[tag].Nodes[0].FirstChild.Data)
		}
	}

	if len(tagsArray) > 0 {
		tags.Value = strings.Join(tagsArray, "; ")
	}

	return tags, tags.Value != ""
}

func FindInfo(infoSelection *goquery.Selection) (miner.AttrV, bool) {
	var info miner.AttrV

	info.ID = 0
	info.AttrType = miner.AttrTypeV{
		ID:    0,
		Title: "text",
	}
	info.Code = miner.AttrCodeV{
		ID:    0,
		Title: "info",
	}

	if infoSelection != nil && len(infoSelection.Nodes) == 1 {
		if infoSelection.Nodes[0].FirstChild.Data == "span" {
			if infoSelection.Nodes[0].FirstChild.FirstChild != nil {
				info.Value = infoSelection.Nodes[0].FirstChild.FirstChild.Data
			} else {
				info.Value = infoSelection.Text()
			}
		} else {
			info.Value = infoSelection.Nodes[0].FirstChild.Data
		}
	}

	return info, info.Value != ""
}

func FindGraphics(graphicsSelection *goquery.Selection) (miner.AttrV, bool) {
	var graphics miner.AttrV

	graphics.ID = 0
	graphics.AttrType = miner.AttrTypeV{
		ID:    0,
		Title: "url",
	}
	graphics.Code = miner.AttrCodeV{
		ID:    0,
		Title: "graphics",
	}

	graphicsLink, exist := graphicsSelection.Attr("src")
	if exist && graphicsLink != "" {
		graphics.Value = graphicsLink
	}

	return graphics, graphics.Value != ""
}
