# Challenge - Indexer Enron Mail

This project is responsible for indexing emails that come from the [Enron Corp](http://www.cs.cmu.edu/~enron/enron_mail_20110402.tgz), Click for download.

# 1. Run Indexer
  ```bash
  SWE Challeng
  │
  ├── indexer
  │   ├── db
  │   │   └── Place enron emails here
  │   └── ... 
  ├── backend-chi-zinc
  │   └── ... 
  ├── fronted
  │   └── ... 
  ├── .gitignore
  └── README.md
  ```
  1. ``` docker compose up ```
  2. ```go run main.go```
  3. write db
  4. config thread write 10 and then 10 (or whatever you want)
  5. if you want disabled  profiling writing 'N'

## Credentials Zincsearch
**user**: admin  
**password**: Complexpass#123

# 2. Run Backend
Execute Command:
``` Go run main.go```

# 3. Run Fronted
  ```
  npm install
  ```
  ```
  npm run serve
  ```


# Chi-zinc-server


# End points
## Get Mails

Return a mailing list

 
 <summary><code>GET</code> <code><b>/</b></code> <code>http://localhost:3000/v1/api/from={from}&max={max}</code></summary>

The parameters are:
- **from (int)**: index from where the search would start
- **max (int)**: the total number of emails that will be returned

Response Success return a [ResponseHits](#response-hits) :
`Code:200`
```  
interface ResponseHits {
	"status":  int,
	"msg":  string,
	"data":  {
		"total":{
			"value":  int
		},
		"hits":  Hit[]
	}
}
```
[Hit](#hit)

Failed response returns [Response Error Interface](#response-error) 
`Code:400`

## Find Mails

Find emails that match the requested query

 
 <summary><code>GET</code> <code><b>/</b></code> <code>http://localhost:3000/v1/api/from={from}&max={max}&terms={terms}</code></summary>

The parameters are:
- **from (int)**: index from where the search would start
- **max (int)**: the total number of emails that will be returned
- **terms (string)**: the query used for search

The searches in Terms are composed this way:

1)  `%20` instead of blank space = search for any match of the terms.
2)  `+` used to returns all data where both terms appear.
3)  `-` used to returns all data where the terms do not appear.
4) `*` used to returns all the data where it starts with the term.

#### example terms:
 - `susan`  find all matches of susan in all fields
 - `susan%20bianca` (instead of "susan bianca")  find all matches of susan or bianca in all fields
 - `-susan`  all matches where susan is not in all fields
 - `susan.bailey%20+bianca.ornelas`  all matches where this susan and bianca.ornelas in all fields
 - `susan*`  all matches starting with susan in all fields
 - `-susan*` all matches you start that do not start with susan in all fields
 - `From:susan`   all susan matches in the From field
 - `-From:susan`   all non-susan matches in the field
 - `From:susan*`  all matches in From that start with susan
 - `-From:susan*`  all matches in From that do not start with susan
 - `+From:susan.bailey%20+To:bianca.ornelas`  all matches in From de susan.bailey and in To de bianca.ornelas

Response Success return a [ResponseHits](#response-hits) :
`Code:200`
```  
interface ResponseHits {
	"status":  int,
	"msg":  string,
	"data":  {
		"total":{
			"value":  int
		},
		"hits":  Hit[]
	}
}
```
[Hit](#hit)

Failed response returns [Response Error Interface](#response-error) 
`Code:400`

## Get Mail

Return a mail

 
 <summary><code>GET</code> <code><b>/</b></code> <code>http://localhost:3000/v1/api/{id}</code></summary>

The parameters are:
- **id (string)**: ID of the requested email

Response Success return a response with a [Mail](#mail) :
```
interface ResponseMail {
	"status":  int,
	"msg":  string,
	"data": Mail
}
```

`Code:200`
 

# Interfaces Responses

## Response Hits

```
interface ResponseHits {
	"status":  int,
	"msg":  string,
	"data":  {
		"total":{
			"value":  int
		},
		"hits":  Hit[]
	}
}

```

## Hit

One Source is equivalent to Mail Resummary
```
interface Hit {
    _index:  string;
    _id:     string;
    _source: Source;
}

interface Source {
    To:      string;
    From:    string;
    Subject: string;
    Date:    Date;
}
```

## Mail
```
interface Mail {
    Message_ID:                string;
    Date:                      Date;
    From:                      string;
    To:                        string;
    Subject:                   string;
    Cc:                        string;
    Mime_Version:              string;
    Content_Type:              string;
    Content_Transfer_Encoding: string;
    Bcc:                       string;
    X_From:                    string;
    X_To:                      string;
    X_cc:                      string;
    X_bcc:                     string;
    X_Folder:                  string;
    X_Origin:                  string;
    X_FileName:                string;
    Content:                   string;
}
```

## Response Error

```
interface ResponseError {
    status: number;
    msg:    string;
    error:  string;
}

```

