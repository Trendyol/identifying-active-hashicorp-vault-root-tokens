package main

import (
	"encoding/json"
	"fmt"
	"github.com/hashicorp/vault/api"
	"github.com/olekukonko/tablewriter"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

const dateLayout = "Mon, 2 Jan 2006 15:04:05 MST"

func main() {
	vaultAddr := os.Getenv("VAULT_ADDR")

	if vaultAddr == "" {
		log.Fatal("The VAULT_ADDR environment must be set.")
	}

	vaultToken := os.Getenv("VAULT_TOKEN")

	if vaultToken == "" {
		log.Fatal("The VAULT_TOKEN environment must be set.")
	}
	config := api.DefaultConfig()
	_ = os.Setenv("VAULT_SKIP_VERIFY", "true")

	client, err := api.NewClient(config)

	if err != nil {
		log.Fatalf("error occurred, detail: %+v", err)
	}

	s, err := client.Logical().List("auth/token/accessors")
	if err != nil {
		log.Fatalf("error occurred, detail: %+v", err)
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Display Name", "Creation Time", "Expiration Time", "Policies", "Token Accessor"})

	for _, accessor := range s.Data["keys"].([]interface{}) {
		lookupAccessor, err := client.Auth().Token().LookupAccessor(accessor.(string))
		if err != nil {
			log.Fatalf("error occurred, detail: %+v", err)
		}
		displayName := lookupAccessor.Data["display_name"]
		creationTime, _ := unixTimeStampToTime(lookupAccessor.Data["creation_time"].(json.Number).String())
		expireTime := lookupAccessor.Data["expire_time"]
		policies := lookupAccessor.Data["policies"]
		var policiesSlice []string
		for _, policy := range policies.([]interface{}) {
			policiesSlice = append(policiesSlice, policy.(string))
		}
		parse, _ := time.Parse(time.RFC3339Nano, fmt.Sprintf("%v", expireTime))
		table.Append([]string{fmt.Sprintf("%v", displayName), creationTime.Format(dateLayout), parse.Format(dateLayout), strings.Join(policiesSlice, ","), fmt.Sprintf("%v", accessor)})
	}
	table.Render()
}

func unixTimeStampToTime(ms string) (time.Time, error) {
	msInt, err := strconv.ParseInt(ms, 10, 64)
	if err != nil {
		return time.Time{}, err
	}

	return time.Unix(msInt, 0), nil
}
