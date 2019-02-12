package main

import (
	"fmt"

	"golang.org/x/net/context"
	"google.golang.org/api/sheets/v4"

	"github.com/pkg/errors"
)

// 新規スプレッドシート作成
func genSpreadSheet(service *sheets.Service, ctx context.Context) (*sheets.Spreadsheet, error) {
	// 新規に生成するスプレッドシートの設定
	spreadsheet := &sheets.Spreadsheet{
		Properties: &sheets.SpreadsheetProperties{
			Title:    "New Spreadsheet", // スプレッドシートの名前
			Locale:   "ja_JP",           // ロケール
			TimeZone: "Asia/Tokyo",      // タイムゾーン
		},
	}

	// スプレッドシートを新規作成
	createResponse, err := service.Spreadsheets.Create(spreadsheet).Context(ctx).Do()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return createResponse, nil
}

// スプレッドシートを編集
func editSpreadSheet(service *sheets.Service, createResponse *sheets.Spreadsheet, m map[string][]interface{}) error {
	// シート1のA1セルを起点にして3行分書き込む
	writeRange := "シート1!A1"
	valueRange := &sheets.ValueRange{
		Values: [][]interface{}{
			[]interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},                      // 1行目
			[]interface{}{200, 400, 500, 450, 300, 700, 400, 200, 350, 500, 900, 800}, // 2行目
			m["code"], // 3行目
		},
	}
	_, err := service.Spreadsheets.Values.Update(createResponse.SpreadsheetId, writeRange, valueRange).ValueInputOption("RAW").Do()
	if err != nil {
		return errors.Errorf("Unable to retrieve data from sheet. %v", err)
	}

	// 枠線の範囲を設定
	gridRange := &sheets.GridRange{
		SheetId:          0,
		StartRowIndex:    0,
		EndRowIndex:      2,
		StartColumnIndex: 0,
		EndColumnIndex:   12,
	}

	// 枠線の色(RGBとAlpha値)を設定
	borderColor := &sheets.Color{
		Red:   0,
		Green: 0,
		Blue:  0,
		Alpha: 0,
	}

	// 枠線単体を設定する構造体に幅とスタイルを設定して色設定を紐付け
	border := &sheets.Border{
		Style: "SOLID",
		Width: 1,
		Color: borderColor,
	}

	// 枠線更新リクエスト構造体に枠線の設定を紐付け
	updateBordersRequest := &sheets.UpdateBordersRequest{
		Range:           gridRange,
		Top:             border,
		Bottom:          border,
		Left:            border,
		Right:           border,
		InnerHorizontal: border,
		InnerVertical:   border,
	}

	// スプレッドシート更新リクエスト構造体に枠線更新リクエストを紐付け
	batchUpdateRequest := &sheets.BatchUpdateSpreadsheetRequest{
		Requests: []*sheets.Request{
			&sheets.Request{
				UpdateBorders: updateBordersRequest,
			},
		},
	}

	// 更新をスプレッドシートに反映
	batchUpdateResponse, err := service.Spreadsheets.BatchUpdate(createResponse.SpreadsheetId, batchUpdateRequest).Do()
	if err != nil {
		return errors.Errorf("Unable to batch update from sheet. %v", err)
	}

	fmt.Printf("Batch update response: %v", batchUpdateResponse)

	return nil
}
