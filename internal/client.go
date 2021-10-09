package client

import (
	"context"
	"io/ioutil"
	"net/http"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/route53"
	"github.com/aws/aws-sdk-go-v2/service/route53/types"
)

func GetIp(url string) (ip string, err error) {
	client := http.Client{}

	response, err := client.Get(url)
	if err != nil {
		return
	}

	defer response.Body.Close()
	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}

	ip = string(bytes)
	return
}

func GetDNSIP(zoneId, recordName *string) (ip string, err error) {
	// Load the Shared AWS Configuration (~/.aws/config)
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return
	}

	r53 := route53.NewFromConfig(cfg)

	listRequest := route53.ListResourceRecordSetsInput{
		HostedZoneId:    zoneId,
		StartRecordType: types.RRTypeA,
		StartRecordName: recordName,
		MaxItems:        aws.Int32(1),
	}

	listResult, err := r53.ListResourceRecordSets(context.TODO(), &listRequest)
	if err != nil {
		return
	}

	if *listRequest.MaxItems == 1 {
		ip = *listResult.ResourceRecordSets[0].ResourceRecords[0].Value
	}
	return
}

func SetDNSIP(zoneId, recordName *string, ttl *int64, ip string) (err error) {

	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return
	}

	r53 := route53.NewFromConfig(cfg)

	change := route53.ChangeResourceRecordSetsInput{
		HostedZoneId: zoneId,
		ChangeBatch: &types.ChangeBatch{
			Changes: []types.Change{
				{
					Action: types.ChangeActionUpsert,
					ResourceRecordSet: &types.ResourceRecordSet{
						Name: recordName,
						Type: types.RRTypeA,
						TTL:  ttl,
						ResourceRecords: []types.ResourceRecord{
							{
								Value: aws.String(ip),
							},
						},
					},
				},
			},
		},
	}

	_, err = r53.ChangeResourceRecordSets(context.TODO(), &change)

	return
}
