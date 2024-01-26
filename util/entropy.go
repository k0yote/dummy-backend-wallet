package util

import (
	"context"

	kms "cloud.google.com/go/kms/apiv1"
	"cloud.google.com/go/kms/apiv1/kmspb"
)

func GenerateRandomBytes(config Config, length int) ([]byte, error) {
	ctx := context.Background()
	client, err := kms.NewKeyManagementClient(ctx)
	if err != nil {
		return nil, err
	}
	defer client.Close()

	// Build the request.
	req := &kmspb.GenerateRandomBytesRequest{
		Location:        config.KmsResourceLocation,
		LengthBytes:     int32(length),
		ProtectionLevel: kmspb.ProtectionLevel_HSM,
	}

	result, err := client.GenerateRandomBytes(ctx, req)
	if err != nil {
		return nil, err
	}

	return result.Data, nil
}
