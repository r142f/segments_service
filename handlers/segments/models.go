package segments

import (
	"segments_service/config"
)

type Segment struct {
	Name string `json:"name" example:"DISCOUNT_30"`
}

func InsertSegment(segment *Segment) error {
	_, err := config.DB.Exec("INSERT INTO Segments (SegmentName) VALUES ($1);", segment.Name)
	return err
}

func DeleteSegmentByName(segmentName string) error {
	_, err := config.DB.Exec("DELETE FROM Segments WHERE SegmentName=$1;", segmentName)
	return err
}

func SelectSegmentIdBySegmentName(segmentName string) (segmentId int, err error) {
	row := config.DB.QueryRow("SELECT SegmentId FROM Segments where SegmentName=$1;", segmentName)
	if err = row.Err(); err != nil {
		return
	}

	err = row.Scan(&segmentId)

	return
}
