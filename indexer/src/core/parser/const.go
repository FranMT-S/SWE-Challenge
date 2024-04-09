package parser

// Used to analize header of file.
const (
	MESSAGE_ID                = "Message-ID:"
	DATE                      = "Date:"
	FROM                      = "From:"
	TO                        = "To:"
	SUBJECT                   = "Subject:"
	CC                        = "Cc:"
	MIME_VERSION              = "Mime-Version:"
	CONTENT_TYPE              = "Content-Type:"
	CONTENT_TRANSFER_ENCODING = "Content-Transfer-Encoding:"
	BCC                       = "Bcc:"
	X_FROM                    = "X-From:"
	X_TO                      = "X-To:"
	X_CC                      = "X-cc:"
	X_BCC                     = "X-bcc:"
	X_FOLDER                  = "X-Folder:"
	X_ORIGIN                  = "X-Origin:"
	X_FILENAME                = "X-FileName:"
	CONTENT                   = "Content"

	K_FATHER = "Father" // used int the struct lineMail
)
