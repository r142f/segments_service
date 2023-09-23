package main

import (
	"segments_service/config"
	"segments_service/handlers/records"
	"segments_service/handlers/segments"
	"segments_service/handlers/users_segments"
	"fmt"
	"net/http"

	_ "segments_service/docs"
	_ "github.com/lib/pq"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title segments_service
// @host localhost:8080
func main() {
	config.Init(false)
	defer config.DB.Close()

	http.HandleFunc("/createSegment", segments.CreateSegment)
	http.HandleFunc("/deleteSegment", segments.DeleteSegment)
	http.HandleFunc("/updateUserSegments", users_segments.UpdateUserSegments)
	http.HandleFunc("/userSegments", users_segments.SegmentsByUser)
	http.HandleFunc("/generateReport", records.Report)
	http.HandleFunc(fmt.Sprintf("/%v/", config.REPORTS_DIRNAME), records.Reports)
	http.HandleFunc("/docs/", httpSwagger.WrapHandler)

	http.ListenAndServe(":8080", nil)
}
