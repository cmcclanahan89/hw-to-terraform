package azuremap

// import (
// 	"context"
// 	"fmt"
// 	"log"

// 	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
// 	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute"
// 	"github.com/Azure/azure-sdk-for-go/sdk/to"
// )

// func GetSku() {
// 	// Authenticate
// 	cred, err := azidentity.NewDefaultAzureCredential(nil)
// 	if err != nil {
// 		log.Fatalf("Failed to create credential: %v", err)
// 	}

// 	// Create a Compute Management client
// 	subscriptionID := ""
// 	client, err := armcompute.NewResourceSkusClient(subscriptionID, cred, nil)
// 	if err != nil {
// 		log.Fatalf("Failed to create client: %v", err)
// 	}

// 	pager := client.NewListPager(&armcompute.ResourceSkusClientListOptions{
// 		Filter: to.Ptr("location eq 'eastus'"),
// 	})

// 	for pager.More() {
// 		page, err := pager.NextPage(context.Background())
// 		if err != nil {
// 			log.Fatalf("Failed to get next page: %v", err)
// 		}
// 		for _, sku := range page.Value {
// 			fmt.Printf("Resource Type: %s, Name: %s, Tier: %s, Locations: %v\n",
// 				*sku.ResourceType, *sku.Name, *sku.Tier, sku.Locations)
// 		}
// 	}
//
