package deviant

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/wizard7414/epos_v2/domain/miner"
	"time"
)

func setResource() miner.ResourceV {
	return miner.ResourceV{
		ID:    0,
		Title: "deviantart",
		Url:   "https://www.deviantart.com",
	}
}

func setSource(document *goquery.Document) miner.SourceV {
	var objectSource miner.SourceV
	objectSource.Resource = setResource()
	objectSource.ID = 0
	objectSource.ChangeDate = time.Now()

	selection, result := FindSelectionByClass(document, "a", "user-link _35MgG qPf26")

	if result {
		sourceTitle, _ := selection.Attr("title")
		sourceUrl, _ := selection.Attr("href")

		objectSource.Title = sourceTitle
		objectSource.Url = sourceUrl
	}

	return objectSource
}

func FindObjectTitle(document *goquery.Document) string {
	var objectTitle string

	titleSelection, result := FindSelectionByHookAndClass(document, "h1", "deviation_title", "_2p6cd")

	if result {
		if len(titleSelection.Nodes) == 1 {
			objectTitle = titleSelection.Nodes[0].FirstChild.Data
		}
	}

	return objectTitle
}

func FindObjectEntry(document *goquery.Document) miner.EntryV {
	var entry miner.EntryV
	entry.Source = setSource(document)

	entryTitleSelection, titleResult := FindSelectionByMeta(document, "og:title")
	if titleResult {
		entryTitle, _ := entryTitleSelection.Attr("content")
		entry.Title = entryTitle
	}

	entryUrlSelection, urlResult := FindSelectionByMeta(document, "og:url")
	if urlResult {
		entryUrl, _ := entryUrlSelection.Attr("content")
		entry.Url = entryUrl
	}

	return entry
}
