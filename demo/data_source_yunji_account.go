package demo

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"io/ioutil"
	"net/http"
)

func dataSourceYunjiAccount() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceYunjiAccountRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func dataSourceYunjiAccountRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	conf := m.(*Configuration)
	endpoint := conf.endpoint
	client := &http.Client{}
	name := d.Get("name").(string)
	request, err := http.NewRequest("GET", fmt.Sprintf("%sdata_source?name=%s", endpoint, name), nil)
	if err != nil {
		return diag.Errorf("err", err)
	}
	response, err := client.Do(request)
	if err != nil {
		return diag.Errorf("err", err)
	}
	body, readErr := ioutil.ReadAll(response.Body)
	if readErr != nil {
		return diag.Errorf("err", readErr)
	}
	defer response.Body.Close()
	var tempMap map[string]interface{}
	json.Unmarshal(body, &tempMap)
	d.Set("name", tempMap["name"])
	d.SetId(tempMap["id"].(string))

	return nil
}
