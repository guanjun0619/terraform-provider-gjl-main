package demo

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDemo() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceDemoCreate,
		ReadContext:   resourceDemoRead,
		UpdateContext: resourceDemoUpdate,
		DeleteContext: resourceDemoDelete,

		Schema: map[string]*schema.Schema{
			"instance_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "名称",
			},
			"disk_size": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "磁盘大小",
			},
			"tags": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "标签",
			},
		},
	}
}

func resourceDemoCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	conf := m.(*Configuration)
	endpoint := conf.endpoint

	client := &http.Client{}
	postData := url.Values{}
	postData.Add("instance_name", d.Get("instance_name").(string))
	postData.Add("disk_size", d.Get("disk_size").(string))
	postData.Add("tags", d.Get("tags").(string))
	request, err := http.NewRequest("POST", endpoint, strings.NewReader(postData.Encode()))
	if err != nil {
		return diag.Errorf("err", err)
	}

	response, err := client.Do(request)
	if err != nil {
		return diag.Errorf("err", err)
	}

	defer response.Body.Close()
	d.SetId("test_id")
	return resourceDemoRead(ctx, d, m)

}

func resourceDemoRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	conf := meta.(*Configuration)
	endpoint := conf.endpoint
	client := &http.Client{}
	request, err := http.NewRequest("GET", fmt.Sprintf("%sget?id=%s", endpoint, d.Id()), nil)
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
	if len(tempMap) == 0 {
		d.Set("instance_name", "guanguan")
		d.Set("disk_size", "100")
		d.Set("tags", "tags")
		return nil
	}
	d.Set("instance_name", tempMap["instance_name"])
	d.Set("disk_size", tempMap["disk_size"])
	d.Set("tags", tempMap["tags"])
	return nil
}

func resourceDemoUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	conf := m.(*Configuration)
	endpoint := conf.endpoint
	postData := url.Values{}

	if d.HasChange("instance_name") {
		postData.Add("instance_name", d.Get("instance_name").(string))
	}
	if d.HasChange("disk_size") {
		postData.Add("disk_size", d.Get("disk_size").(string))
	}
	if d.HasChange("tags") {
		postData.Add("tags", d.Get("test").(string))
	}

	client := &http.Client{}
	request, err := http.NewRequest("PUT", fmt.Sprintf("%supdate?id=%s", endpoint, d.Id()),
		strings.NewReader(postData.Encode()))
	if err != nil {
		return diag.Errorf("err", err)
	}
	_, err = client.Do(request)
	if err != nil {
		return diag.Errorf("err", err)
	}
	return resourceDemoRead(ctx, d, m)
}

func resourceDemoDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	conf := m.(*Configuration)
	endpoint := conf.endpoint
	client := &http.Client{}
	request, err := http.NewRequest("DELETE", fmt.Sprintf("%sdelete?id=%s", endpoint, d.Id()), nil)
	if err != nil {
		return diag.Errorf("err", err)
	}
	_, err = client.Do(request)
	if err != nil {
		return diag.Errorf("err", err)
	}
	return nil
}
