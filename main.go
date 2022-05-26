package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Field struct {
	Name        string `json:"field_name"`
	Type        string `json:"type"`
	Color       string `json:"color,omitempty"`
	DisplayName string `json:"displayName,omitempty"`
}

type Edge struct {
	ID           string `json:"id"`
	Source       string `json:"source"`
	Target       string `json:"target"`
	MainStat     string `json:"mainStat,omitempty"`
	ResponseTime string `json:"detail__responsetime,omitempty"`
}

type Node struct {
	ID            string  `json:"id"`
	Title         string  `json:"title"`
	SubTitle      string  `json:"subTitle,omitempty"`
	MainStat      string  `json:"mainStat,omitempty"`
	SecondaryStat float32 `json:"secondaryStat,omitempty"`
	ArcFailed     float32 `json:"arc__failed,omitempty"`
	ArcPassed     float32 `json:"arc__passed,omitempty"`
	DetailRole    string  `json:"detail__role,omitempty"`
}

func main() {
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodOptions},
		AllowHeaders: []string{echo.HeaderAccept, echo.HeaderContentType},
	}))

	e.Use(middleware.Logger())

	edgeFields := []Field{
		{
			Name: "id",
			Type: "string",
		},
		{
			Name: "source",
			Type: "string",
		},
		{
			Name: "target",
			Type: "string",
		},
		{
			Name: "mainStat",
			Type: "string",
		},
		{
			Name: "detail__responsetime",
			Type: "string",
		},
	}

	nodeFields := []Field{
		{
			Name: "id",
			Type: "string",
		},
		{
			Name: "title",
			Type: "string",
		},
		{
			Name: "subTitle",
			Type: "string",
		},
		{
			Name: "mainStat",
			Type: "string",
		},
		{
			Name: "secondaryStat",
			Type: "number",
		},
		{
			Name:  "arc__failed",
			Type:  "number",
			Color: "red",
		},
		{
			Name:  "arc__passed",
			Type:  "number",
			Color: "green",
		},
		{
			Name: "detail__role",
			Type: "string",
		},
	}

	nodes := []Node{
		{
			ID:            "1",
			Title:         "front-end",
			SubTitle:      "sub",
			ArcFailed:     0.7,
			ArcPassed:     0.3,
			MainStat:      "mainStat",
			SecondaryStat: 10,
		},
		{
			ID:        "2",
			Title:     "cart",
			SubTitle:  "instance:#2",
			ArcFailed: 0.5,
			ArcPassed: 0.5,
			MainStat:  "mainStat",
		},
		{
			ID:        "3",
			Title:     "order",
			SubTitle:  "instance:#3",
			ArcFailed: 0.3,
			ArcPassed: 0.7,
			MainStat:  "mainStat",
		},
		{
			ID:        "4",
			Title:     "user",
			SubTitle:  "instance:#1",
			ArcFailed: 0.5,
			ArcPassed: 0.5,
			MainStat:  "mainStat",
		},
		{
			ID:        "5",
			Title:     "catalogue",
			SubTitle:  "instance:#5",
			ArcFailed: 0.5,
			ArcPassed: 0.5,
			MainStat:  "mainStat",
		},
		{
			ID:         "6",
			Title:      "shipping",
			SubTitle:   "instance:#5",
			ArcFailed:  0.5,
			ArcPassed:  0.5,
			MainStat:   "mainStat",
			DetailRole: "hoge",
		},
	}

	edges := []Edge{
		{
			ID:           "1",
			Source:       "1",
			Target:       "2",
			MainStat:     "10.7KB/s",
			ResponseTime: "7ms",
		},
		{
			ID:       "2",
			Source:   "1",
			Target:   "3",
			MainStat: "53",
		},
		{
			ID:       "3",
			Source:   "1",
			Target:   "4",
			MainStat: "5",
		},
		{
			ID:       "4",
			Source:   "1",
			Target:   "5",
			MainStat: "70",
		},
		{
			ID:       "5",
			Source:   "3",
			Target:   "4",
			MainStat: "100",
		},
		{
			ID:       "6",
			Source:   "3",
			Target:   "6",
			MainStat: "100",
		},
	}

	e.GET("/api/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "ok")
	})

	e.GET("/api/graph/fields", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string][]Field{
			"edges_fields": edgeFields,
			"nodes_fields": nodeFields,
		})
	})

	e.GET("/api/graph/data", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"nodes": nodes,
			"edges": edges,
		})
	})

	e.Logger.Fatal(e.Start(":8080"))
}
