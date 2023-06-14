package cloudwatch_test

import (
	"testing"

	"github.com/K-Phoen/grabana/target/cloudwatch"
	"github.com/stretchr/testify/require"
)

func TestQueriesCanBeCreated(t *testing.T) {
	req := require.New(t)

	query := &cloudwatch.CloudwatchQueryParams{
		Dimensions: map[string]string{
			"QueueName": "test-queue",
		},
		Statistics: []string{"Sum"},
		Namespace:  "AWS/SQS",
		MetricName: "NumberOfMessagesReceived",
		Period:     "30",
		Region:     "us-east-1",
		Expr:       "1-($A/$B)",
	}

	target := cloudwatch.New(*query)

	req.Equal(query.Dimensions, target.Builder.Dimensions)
	req.Equal(query.Statistics, target.Builder.Statistics)
	req.Equal(query.Namespace, target.Builder.Namespace)
	req.Equal(query.MetricName, target.Builder.MetricName)
	req.Equal(query.Period, target.Builder.Period)
	req.Equal(query.Region, target.Builder.Region)
	req.Equal(query.Expr, target.Builder.Expr)
}

func TestRefCanBeConfigured(t *testing.T) {
	req := require.New(t)

	query := &cloudwatch.CloudwatchQueryParams{}

	target := cloudwatch.New(*query, cloudwatch.Ref("A"))

	req.Equal("A", target.Builder.RefID)
}

func TestTargetCanBeHidden(t *testing.T) {
	req := require.New(t)

	query := &cloudwatch.CloudwatchQueryParams{}

	target := cloudwatch.New(*query, cloudwatch.Hide())

	req.True(target.Builder.Hide)
}
