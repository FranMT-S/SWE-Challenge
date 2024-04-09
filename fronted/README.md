# Vue Fronted - Enron Mail Mail Reader
This is a fronted that is used to consume an API that uses zincsearch as a database engine. The objective is to read indexed emails that come from the database of [Enron Corp](http://www.cs.cmu.edu/~enron/enron_mail_20110402.tgz)


This project works in conjunction with two other projects. The[Backend](https://github.com/FranMT-S/chi-zinc-server)  who provides the query endpoints to the database and the [indexer](https://github.com/FranMT-S/Challenge-Go) used to index the data and upload it to the database

## Project setup
```
npm install
```

### Compiles and hot-reloads for development
```
npm run serve
```

### Compiles and minifies for production
```
npm run build
```

### Lints and fixes files
```
npm run lint
```

### Customize configuration
See [Configuration Reference](https://cli.vuejs.org/config/).
