package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
	"unicode"

	pb "3/getBeef/getBeef"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

var url = "https://baconipsum.com/api/?type=meat-and-filler&paras=99&format=text"
var method = "GET"
var client = &http.Client{}
var srvPort = ":8080"
var grpcPort = ":50051"
var tcp = "tcp"

type server struct {
	pb.UnimplementedGetBeefServer
}

func (s *server) GetBeefSummary(ctx context.Context, req *pb.GetBeefSummaryRequest) (*pb.GetBeefSummaryReply, error) {
	return &pb.GetBeefSummaryReply{Beef: GetBeefSummary()}, nil
}

func main() {
	var r = gin.Default()
	var srv = &http.Server{
		Addr:    srvPort,
		Handler: r,
	}
	var grpcServer = grpc.NewServer()
	r.GET("/beef/summary", HandlerGetBeefSummary)

	go func() {
		log.Println(fmt.Sprintf("Server is running on port %s", srvPort))
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	go func() {
		listener, err := net.Listen(tcp, grpcPort)
		if err != nil {
			log.Fatalf("Failed to listen : %v", err)
		}

		pb.RegisterGetBeefServer(grpcServer, &server{})

		log.Println(fmt.Sprintf("GRPC is running on port %s", grpcPort))
		if err := grpcServer.Serve(listener); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}()

	gracefulStop := make(chan os.Signal, 1)
	signal.Notify(gracefulStop, syscall.SIGTERM)
	signal.Notify(gracefulStop, syscall.SIGINT)

	<-gracefulStop
	log.Println("Shutdown Server ...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown: ", err)
	}
	grpcServer.Stop()
	log.Println("Server exiting")
}

func HandlerGetBeefSummary(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"beef": GetBeefSummary(),
	})
	return
}

func GetBeefSummary() (sumBeef map[string]int64) {
	sumBeef = map[string]int64{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	var tokens = tokenize(string(body))
	for _, v := range tokens {
		sumBeef[v] += 1
	}
	return sumBeef
}

func tokenize(text string) []string {
	var delimiterFunc = func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c) && c != '-'
	}
	return strings.FieldsFunc(strings.ToLower(text), delimiterFunc)
}
