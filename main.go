package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// Response is of type APIGatewayProxyResponse since we're leveraging the
// AWS Lambda Proxy Request functionality (default behavior)
//
// https://serverless.com/framework/docs/providers/aws/events/apigateway/#lambda-proxy-integration
type res events.APIGatewayProxyResponse
type req events.APIGatewayProxyRequest
type ctx context.Context

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	fmt.Println("Received Query", request.QueryStringParameters)
	// body := req.QueryStringParameters
	return events.APIGatewayProxyResponse{Body: request.Body, StatusCode: 200}, nil
	// if err := req.ParseForm(); err != nil {
	// 	log.Printf("Error parsing form: %s", err)
	// 	return
	// }
	// fu := req.Form.Get("feed")
	// feedUrl, err := strconv.Atoi(fu)
	// if err != nil {
	// 	log.Printf("Error parsing limit: %s", err)
	// 	return
	// }
	// log.Printf(feedUrl)
	// fp := gofeed.NewParser()
	// feed, _ := fp.ParseURL("http://feeds.twit.tv/twit.xml")
	// fmt.Println(feed.Title)

	// body, err := json.Marshal(map[string]interface{}{
	// 	"message": "Okay so your other function also executed successfully!",
	// })
	// if err != nil {
	// 	return res{StatusCode: 404}, err
	// }
	// var buf bytes.Buffer
	// json.HTMLEscape(&buf, body)

	// resp := res{
	// 	StatusCode:      200,
	// 	IsBase64Encoded: false,
	// 	Body:            buf.String(),
	// 	Headers: map[string]string{
	// 		"Content-Type": "application/json",
	// 	},
	// }

	// return resp, nil
}

func main() {
	lambda.Start(Handler)
}
