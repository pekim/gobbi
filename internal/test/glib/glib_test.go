package glib

//func TestCleanBuild(t *testing.T) {
//	assert.Equal(t, glib.AsciiType(16), glib.ASCII_GRAPH)
//}
//
//func TestFunctionCallWithIntegers(t *testing.T) {
//	assert.Equal(t, int32(4), glib.AsciiDigitValue('4'))
//}
//
//func TestFunctionCallWithStringsNotFreeReturnedString(t *testing.T) {
//	assert.Equal(t, "edcba", glib.Strreverse("abcde"))
//	assert.Equal(t, "abcde", glib.InternString("abcde"))
//}
//
//func TestFunctionCallWithReturnedStringToFree(t *testing.T) {
//	assert.Equal(t, "ABCDE", glib.AsciiStrup("aBcDe", -1))
//}
//
//func TestFunctionCallWithAliasParamAndReturn(t *testing.T) {
//	s := "abc"
//	q := glib.QuarkFromString(s)
//	assert.Equal(t, s, glib.QuarkToString(q))
//}
//
//func TestFunctionReturningError(t *testing.T) {
//	_, _, err := glib.FileOpenTmp("bad/should/not/contain/a/slash")
//	assert.NotNil(t, err)
//
//	glibError := err.(*glib.Error)
//	assert.True(t,
//		strings.Contains(glibError.Message, "should not contain"),
//		"error message not correct")
//
//	// Chack the Error's domain Quark.
//	domain := glib.QuarkToString(glibError.Domain)
//	assert.Equal(t, "g-file-error-quark", domain)
//}
//
//func TestFunctionReturningNoError(t *testing.T) {
//	_, _, err := glib.FilenameFromUri("file://hostname/filename")
//	assert.Nil(t, err)
//}
//
//func TestFunctionReturningOutParam(t *testing.T) {
//	filename, hostname, _ := glib.FilenameFromUri("file://hostname/filename")
//	assert.Equal(t, "/filename", filename)
//	assert.Equal(t, "hostname", hostname)
//}
//
//func TestRecordConstructor(t *testing.T) {
//	date := glib.DateNewDmy(02, 06, 1962)
//	assert.NotNil(t, date)
//	//assert.NotNil(t, date.native)
//}
//
//func TestRecordReceivers(t *testing.T) {
//	date := glib.DateNewDmy(02, 06, 1962)
//	assert.Equal(t, glib.DateDay(2), date.GetDay())
//	assert.Equal(t, glib.DateMonth(6), date.GetMonth())
//	assert.Equal(t, glib.DateYear(1962), date.GetYear())
//	assert.Equal(t, uint32(153), date.GetDayOfYear())
//}
//
//func TestFormatStringFunction(t *testing.T) {
//	formattedLength := glib.Printf("%d %s", 42, "qaz")
//	assert.Equal(t, int32(6), formattedLength)
//}
