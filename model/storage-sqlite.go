package model

import (
	"strings"

	"github.com/crlspe/snp/logging"
	"github.com/crlspe/snp/settings"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

type SnippetModel struct {
	Snippet
	Slug string `gorm:"index"`
}

var db *gorm.DB

func init() {
	if settings.Config.DataUseSqliteDb {
		var err error
		db, err = gorm.Open(
			sqlite.Open(settings.Config.DataSqliteDbName),
			&gorm.Config{},
		)
		if err != nil {
			logging.Debug("Can't connect to Db.")
		}
		db.AutoMigrate(&SnippetModel{})
	}
}

func NewSqliteStorage() SnippetModel {
	return SnippetModel{}
}

func (sm *SnippetModel) Add(snippet Snippet) {
	*sm = SnippetModel{
		Snippet: snippet,
		Slug:    snippet.Description + " " + snippet.Scopes + " " + snippet.Tags,
	}
	db.Create(sm)
}

func (sm *SnippetModel) Search(searchTerms ...string) ([]SnippetModel, error) {
	var snippetsFound []SnippetModel
	for _, term := range searchTerms {
		var snippets []SnippetModel
		var query = db.Model(&snippets)
		var subterms = strings.Split(strings.TrimSpace(term), " ")
		for _, subterm := range subterms {
			query = query.Where("slug LIKE ?", "%"+subterm+"%")
		}
		if err := query.Find(&snippets).Error; err != nil {
			logging.Debug("Error executing query:", err)
		} else {
			snippetsFound = append(snippetsFound, snippets...)
		}
	}
	return snippetsFound, nil
}

func TestGorm() {

	// var snippet = Snippet{
	// 	Description: "Description",
	// 	Scopes:      "Scopes",
	// 	Tags:        "Tags",
	// 	Code:        "Code",
	// }

	// var sm = NewSnippetModel()

	// var searchTerms = []string{"ta"}
	// var result, _ = sm.Search(searchTerms...)

	// fmt.Println(result)
}
