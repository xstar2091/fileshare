package main

import "filesync/src/handler"

func main() {
    sync := handler.CreateSyncHandler()
    sync.Work()
}
