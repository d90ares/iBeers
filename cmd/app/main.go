// package main

// import (
//    "context"
//    "fmt"
//    "google.golang.org/api/option"
//    "log"
//    "os"

//    drive "google.golang.org/api/drive/v3"
// )

// // Constants
// const (
//    ServiceAccount = "C:\\Youtube\\dev\\ServiceAccountCred.json"                  // Please set the json file of Service account.
//    filename       = "C:\\Development\\FreeLance\\GoogleSamples\\Data\\dummy.txt" // Filename
//    SCOPE          = drive.DriveScope
// )

// func main() {

//    ctx := context.Background()
//    srv, err := drive.NewService(ctx, option.WithCredentialsFile(ServiceAccount), option.WithScopes(SCOPE))
//    if err != nil {
//       log.Fatalf("Warning: Unable to create drive Client %v", err)
//    }

//    file, err := os.Open(filename)
//    info, _ := file.Stat()
//    if err != nil {
//       log.Fatalf("Warning: %v", err)
//    }

//    if err != nil {
//       log.Fatalln(err)
//    }
//    defer file.Close()

//    // Create File metadata
//    f := &drive.File{Name: info.Name()}

//    // Create and upload the file
//    res, err := srv.Files.
//       Create(f).
//       Media(file). //context.Background(), file, fileInf.Size(), baseMimeType).
//       ProgressUpdater(func(now, size int64) { fmt.Printf("%d, %d\r", now, size) }).
//       Do()

//    if err != nil {
//       log.Fatalln(err)
//    }

//    fmt.Printf("New file id: %s\n", res.Id)
// }