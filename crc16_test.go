package crc16_test

import (
	"testing"

	"github.com/getrak/crc16"
	"github.com/stretchr/testify/assert"
)

var testData = []byte("123456789")

func testSelectedCRC(t *testing.T, params crc16.Params, expected uint16) {
	table := crc16.MakeTable(params)
	assert.NotNil(t, table)

	crc := crc16.Checksum(testData, table)
	assert.Equal(t, expected, crc)
}

func TestCRC16_ARC(t *testing.T) {
	testSelectedCRC(t, crc16.CRC16_ARC, 0xBB3D)
}

func BenchmarkCRC16_ARC(b *testing.B) {
	table := crc16.MakeTable(crc16.CRC16_ARC)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		crc16.Checksum(testData, table)
	}
}

func TestCRC16_AUG_CCIT(t *testing.T) {
	testSelectedCRC(t, crc16.CRC16_AUG_CCITT, 0xE5CC)
}

func BenchmarkCRC16_AUG_CCITT(b *testing.B) {
	table := crc16.MakeTable(crc16.CRC16_AUG_CCITT)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		crc16.Checksum(testData, table)
	}
}

func TestCRC16_BUYPASS(t *testing.T) {
	testSelectedCRC(t, crc16.CRC16_BUYPASS, 0xFEE8)
}

func TestCRC16_CCITT_FALSE(t *testing.T) {
	testSelectedCRC(t, crc16.CRC16_CCITT_FALSE, 0x29B1)
}

func BenchmarkCRC16_CCITT_FALSE(b *testing.B) {
	table := crc16.MakeTable(crc16.CRC16_CCITT_FALSE)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		crc16.Checksum(testData, table)
	}
}

func TestCRC16_CDMA2000(t *testing.T) {
	testSelectedCRC(t, crc16.CRC16_CDMA2000, 0x4C06)
}

func TestCRC16_DDS_110(t *testing.T) {
	testSelectedCRC(t, crc16.CRC16_DDS_110, 0x9ECF)
}

func TestCRC16_DECT_R(t *testing.T) {
	testSelectedCRC(t, crc16.CRC16_DECT_R, 0x007E)
}

func TestCRC16_DECT_X(t *testing.T) {
	testSelectedCRC(t, crc16.CRC16_DECT_X, 0x007F)
}

func TestCRC16_DNP(t *testing.T) {
	testSelectedCRC(t, crc16.CRC16_DNP, 0xEA82)
}

func TestCRC16_EN_13757(t *testing.T) {
	testSelectedCRC(t, crc16.CRC16_EN_13757, 0xC2B7)
}

func TestCRC16_GENIBUS(t *testing.T) {
	testSelectedCRC(t, crc16.CRC16_GENIBUS, 0xD64E)
}

func TestCRC16_MAXIM(t *testing.T) {
	testSelectedCRC(t, crc16.CRC16_MAXIM, 0x44C2)
}

func TestCRC16_MCRF4XX(t *testing.T) {
	testSelectedCRC(t, crc16.CRC16_MCRF4XX, 0x6F91)
}

func TestCRC16_RIELLO(t *testing.T) {
	testSelectedCRC(t, crc16.CRC16_RIELLO, 0x63D0)
}

func TestCRC16_T10_DIF(t *testing.T) {
	testSelectedCRC(t, crc16.CRC16_T10_DIF, 0xD0DB)
}

func TestCRC16_TELEDISK(t *testing.T) {
	testSelectedCRC(t, crc16.CRC16_TELEDISK, 0x0FB3)
}

func TestCRC16_TMS37157(t *testing.T) {
	testSelectedCRC(t, crc16.CRC16_TMS37157, 0x26B1)
}

func TestCRC16_USB(t *testing.T) {
	testSelectedCRC(t, crc16.CRC16_USB, 0xB4C8)
}

func TestCRC16_CRC_A(t *testing.T) {
	testSelectedCRC(t, crc16.CRC16_CRC_A, 0xBF05)
}

func TestCRC16_KERMIT(t *testing.T) {
	testSelectedCRC(t, crc16.CRC16_KERMIT, 0x2189)
}

func BenchmarkCRC16_KERMIT(b *testing.B) {
	table := crc16.MakeTable(crc16.CRC16_KERMIT)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		crc16.Checksum(testData, table)
	}
}

func TestCRC16_MODBUS(t *testing.T) {
	testSelectedCRC(t, crc16.CRC16_MODBUS, 0x4B37)
}

func TestCRC16_X_25(t *testing.T) {
	testSelectedCRC(t, crc16.CRC16_X_25, 0x906E)
}

func BenchmarkCRC16_X_25(b *testing.B) {
	table := crc16.MakeTable(crc16.CRC16_X_25)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		crc16.Checksum(testData, table)
	}
}

func TestCRC16_XMODEM(t *testing.T) {
	testSelectedCRC(t, crc16.CRC16_XMODEM, 0x31C3)
}

func BenchmarkCRC16_XMODEM(b *testing.B) {
	table := crc16.MakeTable(crc16.CRC16_XMODEM)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		crc16.Checksum(testData, table)
	}
}

func TestCRC16_CMS(t *testing.T) {
	testSelectedCRC(t, crc16.CRC16_CMS, 0xAEE7)
}

func TestCRC16_GSM(t *testing.T) {
	testSelectedCRC(t, crc16.CRC16_GSM, 0xCE3C)
}

func TestCRC_A(t *testing.T) {
	testSelectedCRC(t, crc16.CRC_A, 0xBF05)
}

func TestCRC16_LJ1200(t *testing.T) {
	testSelectedCRC(t, crc16.CRC16_LJ1200, 0xBDF4)
}

func TestCRC16_NRSC_5(t *testing.T) {
	testSelectedCRC(t, crc16.CRC16_NRSC_5, 0xA066)
}

func TestCRC16_OPENSAFETY_A(t *testing.T) {
	testSelectedCRC(t, crc16.CRC16_OPENSAFETY_A, 0x5D38)
}

func TestCRC16_OPENSAFETY_B(t *testing.T) {
	testSelectedCRC(t, crc16.CRC16_OPENSAFETY_B, 0x20FE)
}

func TestCRC16_PROFIBUS(t *testing.T) {
	testSelectedCRC(t, crc16.CRC16_PROFIBUS, 0xA819)
}
