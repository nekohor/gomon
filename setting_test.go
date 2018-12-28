package gomon

import (
    "testing"
    "log"
)

func TestGetDateArray(t *testing.T) {
    s := Setting{} 
    s.StartDate = "21081213"
    s.EndDate = "21081217"
    log.Println(s.GetDateArray())

}