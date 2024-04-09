export interface Mail {
    _id: string
    Message_ID: string
    Date: Date
    From: string
    To: string
    Subject: string
    Cc: string
    Mime_Version: string
    Content_Type: string
    Content_Transfer_Encoding: string
    Bcc: string
    X_From: string
    X_To: string
    X_cc: string
    X_bcc: string
    X_Folder: string
    X_Origin: string
    X_FileName: string
    Content: string
}


export type MailSummary = Pick<Mail, "To" | "From" | "Subject" | "Date">
export type SubMail = Partial<Mail>

export type MailReceivedResponse = Pick<Mail, "_id" | "Date" | "From" | "Subject">



export interface ResponseData<T>  {
	status: number 
	msg:    string 
	data:  T
}


export interface ResponseHits  {
	total:  {
		value: number 
	} 

	hits:Hit[]
}

export interface Hit{
	_index:  string 
    _id:    string 
    _source: MailSummary    
}


export interface ResponseError  {
	status:      string    
	msg:    string    
	error: string    
}


