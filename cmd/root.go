/*
Copyright © 2024 Janos Hardi
*/
package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"golang.org/x/exp/slices"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	t "github.com/yaanno/tenders-cli/types"
)

var Country string

var rootCmd = &cobra.Command{
	Use:   "tenders-cli",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&Country, "country", "c", "", "(required), available options: hu")
	rootCmd.MarkPersistentFlagRequired("country")
	viper.BindPFlag("country", rootCmd.PersistentFlags().Lookup("country"))
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(getCmd)
}

var listCmd = &cobra.Command{
	Use:       "list",
	Short:     "Get a list of tenders",
	ValidArgs: []string{"tender"},
	Args:      cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
	RunE: func(cmd *cobra.Command, args []string) error {
		country := viper.GetString("country")
		if !slices.Contains([]string{"hu", "ro", "pl"}, country) {
			return fmt.Errorf("invalid country code: %s", country)
		}
		res, err := tenderListRequest(country)
		if err != nil {
			return err
		}
		fmt.Println(res.Total)
		return nil
	},
}

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a specific tender",
	Args:  cobra.MatchAll(cobra.ExactArgs(2)),
	RunE: func(cmd *cobra.Command, args []string) error {
		country := viper.GetString("country")
		tenderId := args[1]
		if !slices.Contains([]string{"hu", "ro", "pl"}, country) {
			return fmt.Errorf("invalid country code: %s", country)
		}
		res, err := tenderGetRequest(country, tenderId)
		if err != nil {
			return err
		}
		fmt.Println(res.Title)
		return nil
	},
}

func tenderListRequest(country string) (t.PaginatedTenderList, error) {
	url := fmt.Sprintf("https://tenders.guru/api/%s/tenders", country)
	response, err := http.Get(url)
	tender_list := t.PaginatedTenderList{}

	if err != nil {
		return tender_list, fmt.Errorf("unable to reach url \n%w", err)
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return tender_list, fmt.Errorf("unable to read response body: %s", err)
	}

	jsonErr := json.Unmarshal(body, &tender_list)

	if jsonErr != nil {
		return tender_list, fmt.Errorf("unable to transform json: %s", err)
	}

	return tender_list, err
}

func tenderGetRequest(country string, tenderId string) (t.Tender, error) {
	tender := t.Tender{}

	method := "GET"
	url := fmt.Sprintf("https://tenders.guru/api/%s/tenders/%s", country, tenderId)
	response, err := makeRequest(method, url, nil, nil)

	if err != nil {
		return tender, fmt.Errorf("unable to reach url \n%w", err)
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return tender, fmt.Errorf("unable to read response body: %s", err)
	}

	jsonErr := json.Unmarshal(body, &tender)

	if jsonErr != nil {
		fmt.Println(body)
		return tender, fmt.Errorf("unable to transform json: %s", err)
	}

	return tender, err

}

func makeRequest(method, url string, headers map[string]string, body []byte) (*http.Response, error) {
	client := &http.Client{}

	req, err := http.NewRequest(method, url, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("unable to reach url \n%w", err)
	}

	return resp, nil
}
