package utils

import (
	"fmt"
	"log"
	"strings"
	"time"
)

func IsValidDateFormat(input string) bool {
	dateFormat := "2006-01-02"
	_, err := time.Parse(dateFormat, input)
	return err == nil
}

func IsValidTimeFormat(input string) bool {
	timeFormat := "15:04"
	_, err := time.Parse(timeFormat, input)
	return err == nil
}

func IsTodayEqualTo(targetDate string) bool {
	today := time.Now().Format("2006-01-02")
	return today == targetDate
}

func GetHariIndonesia(englishDay string) string {
	hariMapping := map[string]string{
		"Monday":    "Senin",
		"Tuesday":   "Selasa",
		"Wednesday": "Rabu",
		"Thursday":  "Kamis",
		"Friday":    "Jumat",
		"Saturday":  "Sabtu",
		"Sunday":    "Minggu",
	}

	return hariMapping[englishDay]
}

func GetBulanIndonesia(month string) string {
	bulanMapping := map[string]string{
		"01": "Januari",
		"02": "Februari",
		"03": "Maret",
		"04": "April",
		"05": "Mei",
		"06": "Juni",
		"07": "Juli",
		"08": "Agustus",
		"09": "September",
		"10": "Oktober",
		"11": "November",
		"12": "Desember",
	}

	return bulanMapping[month]
}

func IsDayEqualTo(dayString string) bool {
	loc, err := time.LoadLocation("Asia/Makassar")
	if err != nil {
		log.Println(err.Error())
	}

	timeInWIT := time.Now().In(loc)
	hariInggris := timeInWIT.Weekday().String()
	hariIndonesia := GetHariIndonesia(hariInggris)
	log.Println(hariIndonesia)

	return strings.EqualFold(hariIndonesia, dayString)
}

func IsBeforeTime(targetTime string) bool {
	parsedTime, err := time.Parse("15:04", targetTime)
	if err != nil {
		return false
	}

	now := time.Now()
	witaLocation, err := time.LoadLocation("Asia/Makassar")
	if err != nil {
		return false
	}

	witaTime := now.In(witaLocation)

	timeFormat := time.Date(witaTime.Year(), witaTime.Month(), witaTime.Day(), parsedTime.Hour(), parsedTime.Minute(), 0, 0, witaTime.Location())
	fmt.Println(witaTime, timeFormat)
	return witaTime.Before(timeFormat)
}

func IsAfterTime(targetTime string) bool {
	parsedTime, err := time.Parse("15:04", targetTime)
	if err != nil {
		return false
	}

	now := time.Now()
	witaLocation, err := time.LoadLocation("Asia/Makassar")
	if err != nil {
		return false
	}
	witaTime := now.In(witaLocation)

	timeFormat := time.Date(witaTime.Year(), witaTime.Month(), witaTime.Day(), parsedTime.Hour(), parsedTime.Minute(), 0, 0, witaTime.Location())

	return witaTime.After(timeFormat)
}

func IsValidDate(format, value string) bool {
	_, err := time.Parse(format, value)
	return err == nil
}

func JumlahHariPadaBulan(bulan int, tahun int) int {
	tanggalPertamaBulanBerikutnya := time.Date(tahun, time.Month(bulan)+1, 1, 0, 0, 0, 0, time.UTC)
	tanggalTerakhirBulanSekarang := tanggalPertamaBulanBerikutnya.AddDate(0, 0, -1)
	jumlahHari := tanggalTerakhirBulanSekarang.Day()
	return jumlahHari
}