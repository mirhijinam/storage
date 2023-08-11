package main

import (
	"fmt"
	"log"

	"github.com/mirhijinam/storage/internal/storage"
)

func main() {
	st := storage.NewStorage()
	fmt.Println("storage created: ", *st)

	uploadedFile, err := st.Upload("test.txt", []byte{'h'})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("file is uploaded: ", uploadedFile)

	restoredFile, err := st.GetByID(uploadedFile.ID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("file is restored: ", restoredFile)

}
