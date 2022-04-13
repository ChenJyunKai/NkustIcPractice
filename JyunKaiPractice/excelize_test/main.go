package main

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/xuri/excelize/v2"
)

type Table []struct {
	Foodname   string `json:"foodname"`
	Totalcount int    `json:"totalcount"`
	Price      int    `json:"price"`
	Subtotal   int    `json:"subtotal"`
	Detail     []struct {
		Name   string `json:"name"`
		Count  int    `json:"count"`
		Remark string `json:"remark"`
	} `json:"detail"`
}

func main() {
	f := excelize.NewFile()
	//背景留白
	f.SetSheetViewOptions("Sheet1", -1, excelize.ShowGridLines(false))
	//樣式
	top := excelize.Border{Type: "top", Style: 1, Color: "000000"}
	left := excelize.Border{Type: "left", Style: 1, Color: "000000"}
	right := excelize.Border{Type: "right", Style: 1, Color: "000000"}
	bottom := excelize.Border{Type: "bottom", Style: 1, Color: "000000"}
	nostyle, _ := f.NewStyle(&excelize.Style{
		Font:      &excelize.Font{Size: 12, Family: "新細明體"},
		Border:    []excelize.Border{top, left, right, bottom},
		Alignment: &excelize.Alignment{Vertical: "center", Horizontal: "center"},
	})
	titlestyle, _ := f.NewStyle(&excelize.Style{
		Font:      &excelize.Font{Size: 10, Family: "新細明體", Bold: true},
		Border:    []excelize.Border{top, left, right, bottom},
		Alignment: &excelize.Alignment{Vertical: "center", Horizontal: "center"},
	})
	foodnamestyle, _ := f.NewStyle(&excelize.Style{
		Font:      &excelize.Font{Size: 15, Family: "新細明體"},
		Border:    []excelize.Border{top, left, right, bottom},
		Alignment: &excelize.Alignment{Vertical: "center", Horizontal: "left"},
	})
	otherstyle, _ := f.NewStyle(&excelize.Style{
		Font:      &excelize.Font{Size: 15, Family: "新細明體"},
		Border:    []excelize.Border{top, left, right, bottom},
		Alignment: &excelize.Alignment{Vertical: "center", Horizontal: "right"},
	})
	detailstyle, _ := f.NewStyle(&excelize.Style{
		Font:      &excelize.Font{Size: 12, Family: "新細明體"},
		Border:    []excelize.Border{right},
		Alignment: &excelize.Alignment{Vertical: "center", Horizontal: "left"},
	})
	//便當
	f.SetSheetRow("Sheet1", "B1", &[]interface{}{"商品", "數量", "單價", "小計"})
	f.SetCellStyle("Sheet1", "B1", "E1", titlestyle)
	f.SetColWidth("Sheet1", "B", "B", 30)
	var row_count = 1
	var pricetotal = 0

	bandon := &Table{}
	bandon_json := []byte(`[{"foodname":"正宗池上便當","totalcount":2,"price":80,"subtotal":160,"detail":[{"name":"kyna","count":1,"remark":""},{"name":"何勇","count":1,"remark":""}]},{"foodname":"宮保雞丁便當","totalcount":1,"price":80,"subtotal":80,"detail":[{"name":"王偉龍","count":1,"remark":""}]},{"foodname":"黃金排骨便當","totalcount":6,"price":90,"subtotal":540,"detail":[{"name":"BRAD","count":1,"remark":""},{"name":"張銘杰","count":1,"remark":""},{"name":"周俊立","count":1,"remark":""},{"name":"蔡政達","count":1,"remark":""},{"name":"謝榮長","count":1,"remark":""},{"name":"林宜儒","count":1,"remark":""}]},{"foodname":"正宗池上便當","totalcount":2,"price":85,"subtotal":170,"detail":[{"name":"欣宜","count":1,"remark":""},{"name":"陳立容","count":1,"remark":""}]}]`)
	json.Unmarshal(bandon_json, &bandon)

	for i, value := range *bandon {
		pricetotal += value.Subtotal
		row_count++
		Merge_f := "A" + strconv.Itoa(row_count)
		colum := "B" + strconv.Itoa(row_count)
		colum_f := "C" + strconv.Itoa(row_count)
		colum_e := "E" + strconv.Itoa(row_count)
		f.SetCellValue("Sheet1", Merge_f, i+1)
		f.SetCellStyle("Sheet1", Merge_f, Merge_f, nostyle)
		f.SetSheetRow("Sheet1", colum, &[]interface{}{value.Foodname, value.Totalcount, value.Price, value.Subtotal})
		f.SetCellStyle("Sheet1", colum, colum, foodnamestyle)
		f.SetCellStyle("Sheet1", colum_f, colum_e, otherstyle)
		for _, value2 := range value.Detail {
			row_count++
			colum := "B" + strconv.Itoa(row_count)
			cellvalue := value2.Name + " *" + strconv.Itoa(value2.Count)
			merge_e := "E" + strconv.Itoa(row_count)
			f.SetCellValue("Sheet1", colum, cellvalue)
			f.SetCellStyle("Sheet1", colum, colum, detailstyle)
			f.MergeCell("Sheet1", colum, merge_e)
		}
		Merge_e := "A" + strconv.Itoa(row_count)
		f.MergeCell("Sheet1", Merge_f, Merge_e)
	}
	row_count++
	last_f := "B" + strconv.Itoa(row_count)
	last_e := "E" + strconv.Itoa(row_count)
	outbase := "A" + strconv.Itoa(row_count)
	Pricetotal := "總計:" + strconv.Itoa(pricetotal)
	f.SetCellValue("Sheet1", last_f, Pricetotal)
	f.SetCellStyle("Sheet1", last_f, last_f, otherstyle)
	f.MergeCell("Sheet1", last_f, last_e)
	f.SetCellStyle("Sheet1", outbase, outbase, otherstyle)

	//飲料
	f.SetSheetRow("Sheet1", "H1", &[]interface{}{"商品", "數量", "單價", "小計"})
	f.SetCellStyle("Sheet1", "H1", "K1", titlestyle)
	f.SetColWidth("Sheet1", "H", "H", 25)
	row_count = 1
	pricetotal = 0
	outbase = "G" + strconv.Itoa(row_count)
	f.SetCellStyle("Sheet1", outbase, outbase, otherstyle)

	drink := &Table{}
	drink_json := []byte(`[{"foodname":"布丁奶茶 L","totalcount":1,"price":50,"subtotal":50,"detail":[{"name":"吳誌軒","count":1,"remark":"微糖，去冰"}]},{"foodname":"伯爵奶茶 L","totalcount":1,"price":50,"subtotal":50,"detail":[{"name":"楊淑婷","count":1,"remark":"微糖.去冰"}]},{"foodname":"珍珠奶茶 L","totalcount":1,"price":50,"subtotal":50,"detail":[{"name":"蔡寶嫺","count":1,"remark":":3分糖、去冰"}]},{"foodname":"凍頂烏龍茶 L","totalcount":1,"price":30,"subtotal":30,"detail":[{"name":"周俊立","count":1,"remark":"無糖,微冰"}]},{"foodname":"頂級茉莉綠茶 L","totalcount":1,"price":25,"subtotal":25,"detail":[{"name":"何勇","count":1,"remark":":一分糖 去冰"}]},{"foodname":"琥珀紅茶 L","totalcount":1,"price":25,"subtotal":25,"detail":[{"name":"吳誌軒","count":1,"remark":"微糖微冰"}]},{"foodname":"蘆薈好蜜 瓶","totalcount":1,"price":85,"subtotal":85,"detail":[{"name":"應月萍","count":1,"remark":"半糖正常冰"}]}]`)
	json.Unmarshal(drink_json, &drink)
	for i, value := range *drink {
		pricetotal += value.Subtotal
		row_count++
		Merge_f := "G" + strconv.Itoa(row_count)
		colum := "H" + strconv.Itoa(row_count)
		colum_f := "I" + strconv.Itoa(row_count)
		colum_e := "K" + strconv.Itoa(row_count)
		f.SetCellValue("Sheet1", Merge_f, i+1)
		f.SetCellStyle("Sheet1", Merge_f, Merge_f, nostyle)
		f.SetSheetRow("Sheet1", colum, &[]interface{}{value.Foodname, value.Totalcount, value.Price, value.Subtotal})
		f.SetCellStyle("Sheet1", colum, colum, foodnamestyle)
		f.SetCellStyle("Sheet1", colum_f, colum_e, otherstyle)
		for _, value2 := range value.Detail {
			row_count++
			colum := "H" + strconv.Itoa(row_count)
			cellvalue := value2.Name + " *" + strconv.Itoa(value2.Count) + " :" + value2.Remark
			merge_e := "K" + strconv.Itoa(row_count)
			f.SetCellValue("Sheet1", colum, cellvalue)
			f.SetCellStyle("Sheet1", colum, colum, detailstyle)
			f.MergeCell("Sheet1", colum, merge_e)
		}
		Merge_e := "G" + strconv.Itoa(row_count)
		f.MergeCell("Sheet1", Merge_f, Merge_e)
	}
	row_count++
	last_f = "H" + strconv.Itoa(row_count)
	last_e = "K" + strconv.Itoa(row_count)
	outbase = "G" + strconv.Itoa(row_count)
	Pricetotal = "總計:" + strconv.Itoa(pricetotal)
	f.SetCellValue("Sheet1", last_f, Pricetotal)
	f.SetCellStyle("Sheet1", last_f, last_f, otherstyle)
	f.MergeCell("Sheet1", last_f, last_e)
	f.SetCellStyle("Sheet1", outbase, outbase, otherstyle)

	if err := f.SaveAs("Table.xlsx"); err != nil {
		fmt.Println(err)
	}
}
