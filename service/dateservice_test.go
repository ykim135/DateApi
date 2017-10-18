package service

import (
	"reflect"
	"testing"
	"time"
)

func TestDateService(t *testing.T) {
	expectedOutputJan := time.Date(2017, 1, 2, 0, 0, 0, 0, time.UTC)
	expectedOutputFeb := time.Date(2017, 2, 6, 0, 0, 0, 0, time.UTC)
	expectedOutputMar := time.Date(2017, 3, 6, 0, 0, 0, 0, time.UTC)
	expectedOutputApr := time.Date(2017, 4, 3, 0, 0, 0, 0, time.UTC)
	expectedOutputMay := time.Date(2017, 5, 1, 0, 0, 0, 0, time.UTC)
	expectedOutputJun := time.Date(2017, 6, 5, 0, 0, 0, 0, time.UTC)
	expectedOutputJul := time.Date(2017, 7, 3, 0, 0, 0, 0, time.UTC)
	expectedOutputAug := time.Date(2017, 8, 7, 0, 0, 0, 0, time.UTC)
	expectedOutputSep := time.Date(2017, 9, 4, 0, 0, 0, 0, time.UTC)
	expectedOutputOct := time.Date(2017, 10, 2, 0, 0, 0, 0, time.UTC)
	expectedOutputNov := time.Date(2017, 11, 6, 0, 0, 0, 0, time.UTC)
	expectedOutputDec := time.Date(2017, 12, 4, 0, 0, 0, 0, time.UTC)

	firstMondayJan := GetFirstMonday(1)
	firstMondayFeb := GetFirstMonday(2)
	firstMondayMar := GetFirstMonday(3)
	firstMondayApr := GetFirstMonday(4)
	firstMondayMay := GetFirstMonday(5)
	firstMondayJun := GetFirstMonday(6)
	firstMondayJul := GetFirstMonday(7)
	firstMondayAug := GetFirstMonday(8)
	firstMondaySep := GetFirstMonday(9)
	firstMondayOct := GetFirstMonday(10)
	firstMondayNov := GetFirstMonday(11)
	firstMondayDec := GetFirstMonday(12)

	if firstMondayJan != expectedOutputJan {
		t.Errorf("%v is wrong! Should be %v.", firstMondayJan, expectedOutputJan)
	}
	if firstMondayFeb != expectedOutputFeb {
		t.Errorf("%v is wrong! Should be %v.", firstMondayFeb, expectedOutputFeb)
	}
	if firstMondayMar != expectedOutputMar {
		t.Errorf("%v is wrong! Should be %v.", firstMondayMar, expectedOutputMar)
	}
	if firstMondayApr != expectedOutputApr {
		t.Errorf("%v is wrong! Should be %v.", firstMondayApr, expectedOutputApr)
	}
	if firstMondayMay != expectedOutputMay {
		t.Errorf("%v is wrong! Should be %v.", firstMondayMay, expectedOutputMay)
	}
	if firstMondayJun != expectedOutputJun {
		t.Errorf("%v is wrong! Should be %v.", firstMondayJun, expectedOutputJun)
	}
	if firstMondayJul != expectedOutputJul {
		t.Errorf("%v is wrong! Should be %v.", firstMondayJul, expectedOutputJul)
	}
	if firstMondayAug != expectedOutputAug {
		t.Errorf("%v is wrong! Should be %v.", firstMondayAug, expectedOutputAug)
	}
	if firstMondaySep != expectedOutputSep {
		t.Errorf("%v is wrong! Should be %v.", firstMondaySep, expectedOutputSep)
	}
	if firstMondayOct != expectedOutputOct {
		t.Errorf("%v is wrong! Should be %v.", firstMondayOct, expectedOutputOct)
	}
	if firstMondayNov != expectedOutputNov {
		t.Errorf("%v is wrong! Should be %v.", firstMondayNov, expectedOutputNov)
	}
	if firstMondayDec != expectedOutputDec {
		t.Errorf("%v is wrong! Should be %v.", firstMondayDec, expectedOutputDec)
	}
}

func TestGetAllFirstMonday(t *testing.T) {
	expectedOutput := []time.Time{
		time.Date(2017, 1, 2, 0, 0, 0, 0, time.UTC),
		time.Date(2017, 2, 6, 0, 0, 0, 0, time.UTC),
		time.Date(2017, 3, 6, 0, 0, 0, 0, time.UTC),
		time.Date(2017, 4, 3, 0, 0, 0, 0, time.UTC),
		time.Date(2017, 5, 1, 0, 0, 0, 0, time.UTC),
		time.Date(2017, 6, 5, 0, 0, 0, 0, time.UTC),
		time.Date(2017, 7, 3, 0, 0, 0, 0, time.UTC),
		time.Date(2017, 8, 7, 0, 0, 0, 0, time.UTC),
		time.Date(2017, 9, 4, 0, 0, 0, 0, time.UTC),
		time.Date(2017, 10, 2, 0, 0, 0, 0, time.UTC),
		time.Date(2017, 11, 6, 0, 0, 0, 0, time.UTC),
		time.Date(2017, 12, 4, 0, 0, 0, 0, time.UTC),
	}

	var allFirstMonday []time.Time
	var month = time.Month(1)

	output := GetAllFirstMonday(allFirstMonday, month)

	if !reflect.DeepEqual(output, expectedOutput) {
		t.Errorf("get all monday wrong! Got %v, should be %v", output, expectedOutput)
	}
}
