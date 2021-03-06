package main

import (
	"log"

	"runtime/debug"

	"github.com/donNewtonAlpha/AbstractOLT/api"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

// Authentication holds the login/password
type Authentication struct {
	Login    string
	Password string
}

// GetRequestMetadata gets the current request metadata
func (a *Authentication) GetRequestMetadata(context.Context, ...string) (map[string]string, error) {
	return map[string]string{
		"login":    a.Login,
		"password": a.Password,
	}, nil
}

// RequireTransportSecurity indicates whether the credentials requires transport security
func (a *Authentication) RequireTransportSecurity() bool {
	return true
}

func main() {
	var conn *grpc.ClientConn
	creds, err := credentials.NewClientTLSFromFile("cert/server.crt", "AbstractOLT.dev.atl.foundry.att.com")
	if err != nil {
		log.Fatalf("could not load tls cert: %s", err)
	}
	// Setup the login/pass
	auth := Authentication{
		Login:    "john",
		Password: "doe",
	}
	conn, err = grpc.Dial(":7777", grpc.WithTransportCredentials(creds), grpc.WithPerRPCCredentials(&auth))
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := api.NewAddChassisClient(conn)

	response, err := c.CreateChassis(context.Background(), &api.AddChassisMessage{CLLI: "my cilli", VCoreIP: "192.168.0.1", VCorePort: 9191})
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}
	log.Printf("Response from server: %s", response.GetDeviceID())
	d := api.NewAddOLTChassisClient(conn)
	newResponse, err := d.CreateOLTChassis(context.Background(), &api.AddOLTChassisMessage{CLLI: "my cilli", SlotIP: "12.2.2.0", SlotPort: 9191, Hostname: "SlotOne", Type: api.AddOLTChassisMessage_edgecore})
	if err != nil {
		debug.PrintStack()
		log.Fatalf("Error when calling SayHello: %s", err)
	}
	log.Printf("Response from server: %s", newResponse.GetDeviceID())
}
