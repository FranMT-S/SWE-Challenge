package parser

import (
	"strings"
)

// ILineReader provides methods to read and analyze a line
//   - Read: is in charge of analyzing the line
//   - getMapData: is responsible for returning a map that contains the email data
type ILineReader[T string | *lineMail] interface {
	Read(line T)
	getMapData() map[string]string
}

// lineByLineReader read and analyze line by line and create a map with mail data
//   - beforeLecture: it is used for fields with multilines
type lineByLineReader struct {
	mailMap       map[string]string
	beforeLecture string
}

// newLineByLineReader return a lineByLineReader object that implement ILineReader
func newLineByLineReader() *lineByLineReader {
	return &lineByLineReader{mailMap: make(map[string]string)}
}

// getMapData return a map that contains the email data
func (lineReader lineByLineReader) getMapData() map[string]string {
	return lineReader.mailMap
}

// Read read line by line  searching the email fields
// X-Filename is the field that marks the final of header
func (lineReader *lineByLineReader) Read(line string) {
	if lineReader.mailMap[X_FILENAME] != "" {
		lineReader.mailMap[CONTENT] += line
	} else if strings.HasPrefix(line, X_FROM) && lineReader.mailMap[X_FROM] == "" {
		lineReader.mailMap[X_FROM] = line[len(X_FROM):]
		lineReader.beforeLecture = X_FROM
	} else if strings.HasPrefix(line, X_TO) && lineReader.mailMap[X_TO] == "" {
		lineReader.mailMap[X_TO] = line[len(X_TO):]
		lineReader.beforeLecture = X_TO
	} else if strings.HasPrefix(line, X_CC) && lineReader.mailMap[X_CC] == "" {
		lineReader.mailMap[X_CC] = line[len(X_CC):]
		lineReader.beforeLecture = X_CC
	} else if strings.HasPrefix(line, X_BCC) && lineReader.mailMap[X_BCC] == "" {
		lineReader.mailMap[X_BCC] = line[len(X_BCC):]
		lineReader.beforeLecture = X_BCC
	} else if strings.HasPrefix(line, X_FOLDER) && lineReader.mailMap[X_FOLDER] == "" {
		lineReader.mailMap[X_FOLDER] = line[len(X_FOLDER):]
		lineReader.beforeLecture = X_FOLDER
	} else if strings.HasPrefix(line, X_ORIGIN) && lineReader.mailMap[X_ORIGIN] == "" {
		lineReader.mailMap[X_ORIGIN] = line[len(X_ORIGIN):]
		lineReader.beforeLecture = X_ORIGIN
	} else if strings.HasPrefix(line, X_FILENAME) && lineReader.mailMap[X_FILENAME] == "" {
		lineReader.mailMap[X_FILENAME] = line[len(X_FILENAME):]
		lineReader.beforeLecture = X_FILENAME
	} else if strings.HasPrefix(line, MESSAGE_ID) && lineReader.mailMap[MESSAGE_ID] == "" {
		lineReader.mailMap[MESSAGE_ID] = line[len(MESSAGE_ID):]
		lineReader.beforeLecture = MESSAGE_ID
	} else if strings.HasPrefix(line, DATE) && lineReader.mailMap[DATE] == "" {
		lineReader.mailMap[DATE] = line[len(DATE):]
		lineReader.beforeLecture = DATE
	} else if strings.HasPrefix(line, FROM) && lineReader.mailMap[FROM] == "" {
		lineReader.mailMap[FROM] = line[len(FROM):]
		lineReader.beforeLecture = FROM
	} else if strings.HasPrefix(line, TO) && lineReader.mailMap[TO] == "" {
		lineReader.mailMap[TO] = line[len(TO):]
		lineReader.beforeLecture = TO
	} else if strings.HasPrefix(line, SUBJECT) && lineReader.mailMap[SUBJECT] == "" {
		lineReader.mailMap[SUBJECT] = line[len(SUBJECT):]
		lineReader.beforeLecture = SUBJECT
	} else if strings.HasPrefix(line, CC) && lineReader.mailMap[CC] == "" {
		lineReader.mailMap[CC] = line[len(CC):]
		lineReader.beforeLecture = CC
	} else if strings.HasPrefix(line, BCC) && lineReader.mailMap[BCC] == "" {
		lineReader.mailMap[BCC] = line[len(BCC):]
		lineReader.beforeLecture = BCC
	} else if strings.HasPrefix(line, MIME_VERSION) && lineReader.mailMap[MIME_VERSION] == "" {
		lineReader.mailMap[MIME_VERSION] = line[len(MIME_VERSION):]
		lineReader.beforeLecture = MIME_VERSION
	} else if strings.HasPrefix(line, CONTENT_TYPE) && lineReader.mailMap[CONTENT_TYPE] == "" {
		lineReader.mailMap[CONTENT_TYPE] = line[len(CONTENT_TYPE):]
		lineReader.beforeLecture = CONTENT_TYPE
	} else if strings.HasPrefix(line, CONTENT_TRANSFER_ENCODING) && lineReader.mailMap[CONTENT_TRANSFER_ENCODING] == "" {
		lineReader.mailMap[CONTENT_TRANSFER_ENCODING] = line[len(CONTENT_TRANSFER_ENCODING):]
		lineReader.beforeLecture = CONTENT_TRANSFER_ENCODING
	} else if lineReader.beforeLecture != "" {
		lineReader.mailMap[lineReader.beforeLecture] += line
	}
}

// Reader that uses a lineMail object to ensure the use of goroutines
//   - line: a object lineMail
//   - headLineFlag: mark the final of header when X-FileName is encountered
type lineByLineReaderAsync struct {
	line         *lineMail
	headLineFlag int // mark the final of header when X-FileName is encountered
}

// newLineByLineReaderAsync return a lineByLineReaderAsync object that implement ILineReader.
//
// Reader that uses a lineMail object to ensure the use of goroutines.
func newLineByLineReaderAsync() *lineByLineReaderAsync {
	return &lineByLineReaderAsync{headLineFlag: -1}
}

// Read Receives a lineMail object and analize the line of file.
//
// this function mutates the lineMail.
func (lineReader *lineByLineReaderAsync) Read(line *lineMail) {
	if lineReader.headLineFlag > 0 && lineReader.headLineFlag < line.lineNumber {
		line.data = line.lineToAnalize
		line.field = CONTENT
	} else if strings.HasPrefix(line.lineToAnalize, X_FROM) {
		line.data = line.lineToAnalize[len(X_FROM):]
		line.field = X_FROM
	} else if strings.HasPrefix(line.lineToAnalize, X_TO) {
		line.data = line.lineToAnalize[len(X_TO):]
		line.field = X_TO
	} else if strings.HasPrefix(line.lineToAnalize, X_CC) {
		line.data = line.lineToAnalize[len(X_CC):]
		line.field = X_CC
	} else if strings.HasPrefix(line.lineToAnalize, X_BCC) {
		line.data = line.lineToAnalize[len(X_BCC):]
		line.field = X_BCC
	} else if strings.HasPrefix(line.lineToAnalize, X_FOLDER) {
		line.data = line.lineToAnalize[len(X_FOLDER):]
		line.field = X_FOLDER
	} else if strings.HasPrefix(line.lineToAnalize, X_ORIGIN) {
		line.data = line.lineToAnalize[len(X_ORIGIN):]
		line.field = X_ORIGIN
	} else if strings.HasPrefix(line.lineToAnalize, X_FILENAME) {
		line.data = line.lineToAnalize[len(X_FILENAME):]
		line.field = X_FILENAME
		lineReader.headLineFlag = line.lineNumber
	} else if strings.HasPrefix(line.lineToAnalize, MESSAGE_ID) {
		line.data = line.lineToAnalize[len(MESSAGE_ID):]
		line.field = MESSAGE_ID
	} else if strings.HasPrefix(line.lineToAnalize, DATE) {
		line.data = line.lineToAnalize[len(DATE):]
		line.field = DATE
	} else if strings.HasPrefix(line.lineToAnalize, FROM) {
		line.data = line.lineToAnalize[len(FROM):]
		line.field = FROM
	} else if strings.HasPrefix(line.lineToAnalize, TO) {
		line.data = line.lineToAnalize[len(TO):]
		line.field = TO
	} else if strings.HasPrefix(line.lineToAnalize, SUBJECT) {
		line.data = line.lineToAnalize[len(SUBJECT):]
		line.field = SUBJECT
	} else if strings.HasPrefix(line.lineToAnalize, CC) {
		line.data = line.lineToAnalize[len(CC):]
		line.field = CC
	} else if strings.HasPrefix(line.lineToAnalize, BCC) {
		line.data = line.lineToAnalize[len(BCC):]
		line.field = BCC
	} else if strings.HasPrefix(line.lineToAnalize, MIME_VERSION) {
		line.data = line.lineToAnalize[len(MIME_VERSION):]
		line.field = MIME_VERSION
	} else if strings.HasPrefix(line.lineToAnalize, CONTENT_TYPE) {
		line.data = line.lineToAnalize[len(CONTENT_TYPE):]
		line.field = CONTENT_TYPE
	} else if strings.HasPrefix(line.lineToAnalize, CONTENT_TRANSFER_ENCODING) {
		line.data = line.lineToAnalize[len(CONTENT_TRANSFER_ENCODING):]
		line.field = CONTENT_TRANSFER_ENCODING
	} else {
		line.data = line.lineToAnalize
		line.field = K_FATHER
	}

}

// getMapData return a map that contains the email data
func (lineReader *lineByLineReaderAsync) getMapData() map[string]string {
	temp := lineReader.line
	mailMap := make(map[string]string)
	for {
		if temp == nil {
			break
		}

		if lineReader.headLineFlag < temp.lineNumber {
			mailMap[CONTENT] = temp.lineToAnalize + mailMap[CONTENT]
		} else {
			field := temp.getField()
			mailMap[field] = temp.data + mailMap[field]
		}

		temp = temp.lineFather
	}
	return mailMap
}
