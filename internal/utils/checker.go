package utils

import (
	"crypto/rand"
	"encoding/binary"
	"math/big"
	"mime/multipart"
	"strconv"
	"time"
)

func IsPDF(file *multipart.FileHeader) bool {
	f, err := file.Open()

	if err != nil {
		return false
	}

	defer f.Close()
	// PDF files start with "%PDF"
	buf := make([]byte, 4)
	_, err = f.Read(buf)

	if err != nil {
		return false
	}

	return string(buf) == "%PDF"
}

// return the current time in GMT + 0 timezone
func GMTTime() time.Time {
	return time.Now().In(time.FixedZone("GMT", 0))
}

func CheckLastIDLimit(lastID, limit, typeOfQuery string) (int, int) {
	lastIDInt, limitInt := 0, 0

	if typeOfQuery == "chat" {
		lastIDInt, limitInt = 9999, 50
	} else {
		lastIDInt, limitInt = 0, 50
	}

	if lastID == "" || limit == "" || len(lastID) > 9 {
		return lastIDInt, limitInt
	}

	limitInt, _ = strconv.Atoi(limit)
	lastIDInt, _ = strconv.Atoi(lastID)

	return lastIDInt, limitInt
}

// check gmt + 0
func CheckGMTTime(t time.Time) bool {
	_, offset := t.Zone()
	return offset == 0
}

func RandomOTP() int {
	// Use crypto/rand for secure random number generation
	n, err := rand.Int(rand.Reader, big.NewInt(900000))
	if err != nil {
		// Fallback to time-based seed if crypto/rand fails (shouldn't happen)
		seed := time.Now().UnixNano()
		n = big.NewInt(seed % 900000)
	}
	return int(n.Int64()) + 100000
}

func RandomUsername() string {
	const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	const length = 8
	b := make([]byte, length)

	for i := range b {
		n, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			// Fallback: use time-based seed converted to bytes
			var seedBytes [8]byte
			binary.BigEndian.PutUint64(seedBytes[:], uint64(time.Now().UnixNano()))
			n = big.NewInt(int64(seedBytes[i%8] % uint8(len(charset))))
		}
		b[i] = charset[n.Int64()]
	}
	return string(b)
}
