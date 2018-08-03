// Use this client when there is extra layer of security like Oauth or TLS based auth.
package main

import (
	"context"
	"flag"
	"fmt"
	google_protobuf "github.com/golang/protobuf/ptypes/wrappers"
	tf_core_framework "github.com/itnilesh/tensor-flow-inception-client/proto/tensorflow/core/framework"
	pb "github.com/itnilesh/tensor-flow-inception-client/proto/tensorflow_serving/apis"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
	"io/ioutil"
	"log"
	"path/filepath"
	"time"
)

// ClientConfig - Information needed to build grpc client.
type ClientConfig struct {
	ServingAddress string
	AuthToken      string
	ImageFilePath  string
	CertFilePath   string
}

// NewClientConfig - Create new client config from flags.
func NewClientConfig(servingAddress string,
	authToken string,
	imageFilePath string,
	certFilePath string) (*ClientConfig, error) {
	clientCfg := ClientConfig{
		ServingAddress: servingAddress,
		AuthToken:      authToken,
		ImageFilePath:  imageFilePath,
		CertFilePath:   certFilePath}
	return &clientCfg, nil
}

// main - driver method
func main() {

	servingAddress := flag.String("serving-address", "localhost:9000", "The tensorflow serving address")
	authToken := flag.String("auth-token", "", "OAuth token sent as Authorization header value")
	imageFilePath := flag.String("image-file-path", "", "The tensorflow inception input image")
	certFilePath := flag.String("cert-file-path", "", "TLS cert path")

	flag.Parse()

	fmt.Printf(" passed flags are >>  --serving-address %s  --auth-token %s --image-file-path %s --cert-file-path %s ", *servingAddress, *authToken, *imageFilePath, *certFilePath)
	if len(*servingAddress) == 0 {
		log.Fatalln("Please pass reauired flag --serving-address")
	}

	if len(*authToken) == 0 {
		log.Println("Warn: --auth-token flag  is not passed")
	}

	if len(*imageFilePath) == 0 {
		log.Fatalln("Please reauired flag --image-file-path")
	}

	if len(*certFilePath) == 0 {
		log.Println("Warn: --cert-file-path flag is not passed")
	}

	clientCfg, err := NewClientConfig(*servingAddress,
		*authToken,
		*imageFilePath,
		*certFilePath)
	if err != nil {
		log.Fatalln(err)
	}

	// Create connection
	conn, err := getConn(clientCfg)
	if err != nil {
		log.Fatalln(err)
	}

	defer conn.Close()

	// Create New GRPC request for inception model
	client := pb.NewPredictionServiceClient(conn)

	// Create headers
	md := createHeaders(clientCfg)
	outCtx := metadata.NewOutgoingContext(context.Background(), *md)

	request, err := newPredictRequest(clientCfg)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Calling predict..")
	start := time.Now()
	resp, err := client.Predict(outCtx, request)
	elapsed := time.Since(start)
	log.Printf("Predict took %s", elapsed)
	log.Println("Got response")

	if err != nil {
		log.Fatalln(err)
	}

	log.Println(resp)

}

func getConn(clientCfg *ClientConfig) (*grpc.ClientConn, error) {

	// create insecure client
	if len(clientCfg.CertFilePath) == 0 {
		conn, err := grpc.Dial(clientCfg.ServingAddress, grpc.WithInsecure())
		return conn, err
	}

	// create secure client
	certFullPath, err := filepath.Abs(clientCfg.CertFilePath)
	if err != nil {
		return nil, err
	}
	log.Println(certFullPath)
	creds, err := credentials.NewClientTLSFromFile(certFullPath, "")
	if err != nil {
		return nil, err
	}

	conn, err := grpc.Dial(clientCfg.ServingAddress, grpc.WithTransportCredentials(creds))
	if err != nil {
		return nil, err
	}

	return conn, err

}

func newPredictRequest(clientCfg *ClientConfig) (*pb.PredictRequest, error) {
	imgPath, err := filepath.Abs(clientCfg.ImageFilePath)
	if err != nil {
		return nil, err
	}

	imageBytes, err := ioutil.ReadFile(imgPath)
	if err != nil {
		return nil, err
	}

	request := &pb.PredictRequest{
		ModelSpec: &pb.ModelSpec{
			Name:          "inception",
			SignatureName: "predict_images",
			Version: &google_protobuf.Int64Value{
				Value: int64(1),
			},
		},
		Inputs: map[string]*tf_core_framework.TensorProto{
			"images": &tf_core_framework.TensorProto{
				Dtype: tf_core_framework.DataType_DT_STRING,
				TensorShape: &tf_core_framework.TensorShapeProto{
					Dim: []*tf_core_framework.TensorShapeProto_Dim{
						&tf_core_framework.TensorShapeProto_Dim{
							Size: int64(1),
						},
					},
				},
				StringVal: [][]byte{imageBytes},
			},
		},
	}

	return request, nil
}

func createHeaders(clientCfg *ClientConfig) *metadata.MD {
	md := metadata.Pairs(
		"Authorization", clientCfg.AuthToken,
	)
	return &md
}
