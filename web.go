/*package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"os"
	
)

func main() {
	http.HandleFunc("/", hello)
	fmt.Println("listening...")
	err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if err != nil {
		panic(err)
	
			}
}

func hello(res http.ResponseWriter, req *http.Request) {
	b, err := ioutil.ReadFile("domains.txt") // just pass the file name
    	if err != nil {
        fmt.Print(err)
			}
	str := string(b)
	fmt.Fprintln(res, str) }*/
	
package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"os"	
	"github.com/aws/aws-sdk-go/aws"
        "github.com/aws/aws-sdk-go/aws/session"
        "github.com/aws/aws-sdk-go/service/s3"
)

func main() {
	http.HandleFunc("/", hello)
	fmt.Println("listening...")
	err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if err != nil {
		panic(err)
	
			}
}

func hello(res http.ResponseWriter, req *http.Request) {
	
sess := session.Must(session.NewSession())
// Create a downloader with the session and default options
downloader := s3manager.NewDownloader(sess)

// Create a file to write the S3 Object contents to.
f, err := os.Create(filename)
if err != nil {
    return fmt.Errorf("failed to create file %q, %v", filename, err)
}

// Write the contents of S3 Object to the file
n, err := downloader.Download(f, &s3.GetObjectInput{
    Bucket: aws.String(bucket),
    Key:    aws.String(key),
})
if err != nil {
    return fmt.Errorf("failed to download file, %v", err)
}
fmt.Printf("file downloaded, %d bytes\n", n)
	fmt.Fprintln(res, n)	
}

