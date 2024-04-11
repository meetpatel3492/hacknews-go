package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/enescakir/emoji"
	"github.com/gin-gonic/gin"
	"github.com/meetpatel3492/go-first-module/clients"
	"github.com/meetpatel3492/go-first-module/structs"
)

func HandleMain(c *gin.Context) {
	var message string = fmt.Sprintf("Hello,%v %v from first GO webapp",emoji. emoji.WavingHand)
	c.JSON(http.StatusOK, gin.H{
		"message": message,
	})
}

func TopStories(c *gin.Context) {
	// Get call to recieve top articleIds
	var url string = "https://hacker-news.firebaseio.com/v0/topstories.json?print=pretty"
	resp, err := clients.HackerNewsHttpClient.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	topStoriesResp := new(structs.TopStoriesResponse)
	json.NewDecoder(resp.Body).Decode(topStoriesResp)

	var array []int
	for _, v := range *topStoriesResp {
		array = append(array, v)
	}

	var articles []structs.ArticleResponseById
	for i := 0; i < len(*topStoriesResp); i++ {
		articleZ, _ := GetArticle(array[i])
		articles = append(articles, *articleZ)
	}

	articlesJson, err2 := json.Marshal(articles)
	if err2 != nil {
		log.Fatal(err2)
	}
	c.Header("Content-Type", "application/json")
	c.Writer.Write(articlesJson)
}

func TopStoriesConcurrent(c *gin.Context) {
	var url string = "https://hacker-news.firebaseio.com/v0/topstories.json?print=pretty"
	resp, err := clients.HackerNewsHttpClient.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	topStoriesResp := new(structs.TopStoriesResponse)
	json.NewDecoder(resp.Body).Decode(topStoriesResp)

	var array []int
	for _, v := range *topStoriesResp {
		array = append(array, v)
	}
	var articles []structs.ArticleResponseById
	wg := sync.WaitGroup{}

	for i := 0; i < len(array); i++ {
		wg.Add(1)
		go func(i int) {
			articleZ, _ := GetArticle(array[i])
			articles = append(articles, *articleZ)
			wg.Done()
		}(i)
	}
	wg.Wait()
	articlesJson, _ := json.Marshal(articles)
	c.Header("Content-Type", "application/json")
	c.Writer.Write(articlesJson)
}

func GetArticle(articleId int) (article *structs.ArticleResponseById, err error) {
	var baseUrl string = "https://hacker-news.firebaseio.com/v0/item/%d.json?print=pretty"
	url := fmt.Sprintf(baseUrl, articleId)
	resp, respErr := clients.HackerNewsHttpClient.Get(url)
	if respErr != nil {
		return nil, err
	}
	jsonErr := json.NewDecoder(resp.Body).Decode(&article)
	if jsonErr != nil {
		return nil, err
	}
	return article, nil
}
