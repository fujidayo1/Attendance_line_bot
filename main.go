package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

func main() {
	credential := option.WithCredentialsFile("credentials/secret.json")
	srv, err := sheets.NewService(context.TODO(), credential)
	if err != nil {
		log.Fatal(err)
	}

	data := []*sheets.ValueRange{
		{
			Range: "A4:B5",
			Values: [][]interface{}{
				{"kyuuri", "verygood"},
				{"suika", "verygood"},
			},
			MajorDimension: "ROWS",
		},
	}

	reqest := &sheets.BatchUpdateValuesRequest{
		ValueInputOption: "USER_ENTERED",
		Data:             data,
	}

	ctx := context.Background()

	res, err := srv.Spreadsheets.Values.BatchUpdate(id.SpreadsheetID(), reqest).Context(ctx).Do()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%#v\n", res)

}
