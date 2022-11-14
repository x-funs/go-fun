package fun

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMd5(t *testing.T) {
	assert.Equal(t, "d41d8cd98f00b204e9800998ecf8427e", Md5(""))
	assert.Equal(t, "df10ef8509dc176d733d59549e7dbfaf", Md5("123456abc"))
	assert.Equal(t, "21232f297a57a5a743894a0e4a801fc3", Md5("admin"))
	assert.Equal(t, "e10adc3949ba59abbe56e057f20f883e", Md5("123456"))
	assert.Equal(t, "49ba59abbe56e057", Md5Bit16("123456"))
	assert.Equal(t, "a32b4da32d9a67a5", Md5Bit16("df"))
}

func TestSha(t *testing.T) {
	assert.Equal(t, "da39a3ee5e6b4b0d3255bfef95601890afd80709", Sha1(""))
	assert.Equal(t, "a172ffc990129fe6f68b50f6037c54a1894ee3fd", Sha1("123456abc"))
	assert.Equal(t, "931145d4ddd1811be545e4ac88a81f1fdbfaf0779c437efba16b884595274d11", Sha256("123456abc"))
	assert.Equal(t, "2545507ada3a26999b11ec0324ae76e0cdef629c4a28b24be3965d24e1c406180a8ef79626c77fb3f556bfd59ab54920", Sha384("123456abc"))
	assert.Equal(t, "8756869d440a13e93979197b5d7839c823de87c2b115bce0dee62030af3b5b63114a517f1ab02509bfefa88527369ae0ad7946990f27dcb37711a7d6fb9bc893", Sha512("123456abc"))
}

func TestBase64(t *testing.T) {
	assert.Equal(t, "", Base64Encode(""))
	assert.Equal(t, "MTIzNDU2YWJj", Base64Encode("123456abc"))
	assert.Equal(t, "aHR0cHM6Ly93d3cuYmFpZHUuY29tL3M/aWU9dXRmLTgmZj04JnJzdl9icD0xJnRuPWJhaWR1", Base64Encode("https://www.baidu.com/s?ie=utf-8&f=8&rsv_bp=1&tn=baidu"))
	assert.Equal(t, "aHR0cHM6Ly93d3cuYmFpZHUuY29tL3M_aWU9dXRmLTgmZj04JnJzdl9icD0xJnRuPWJhaWR1", Base64UrlEncode("https://www.baidu.com/s?ie=utf-8&f=8&rsv_bp=1&tn=baidu"))

	assert.Equal(t, "123456abc", Base64Decode("MTIzNDU2YWJj"))
	assert.Equal(t, "https://www.baidu.com/s?ie=utf-8&f=8&rsv_bp=1&tn=baidu", Base64Decode("aHR0cHM6Ly93d3cuYmFpZHUuY29tL3M/aWU9dXRmLTgmZj04JnJzdl9icD0xJnRuPWJhaWR1"))
	assert.Equal(t, "https://www.baidu.com/s?ie=utf-8&f=8&rsv_bp=1&tn=baidu", Base64UrlDecode("aHR0cHM6Ly93d3cuYmFpZHUuY29tL3M_aWU9dXRmLTgmZj04JnJzdl9icD0xJnRuPWJhaWR1"))
}
