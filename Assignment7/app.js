var createError = require('http-errors');
var express = require('express');
var path = require('path');
var cookieParser = require('cookie-parser');
var logger = require('morgan');

const bodyParser = require('body-parser');

//compression used to compress the data
const compression = require('compression');

//helmet module from preventing various attavck
const helmet = require('helmet');
const cors = require('cors'); //Enable All CORS Requests

var app = express();

// view engine setup
app.set('views', path.join(__dirname, 'views'));
app.set('view engine', 'jade');

app.use(cors());
app.use(logger('dev'));

app.use(bodyParser.json({ limit: '50mb' }));
app.use(bodyParser.urlencoded({ limit: '50mb', extended: true }));
app.use(cookieParser());

//compression module for compression
app.use(compression());

//helmet module for revent from attack
app.use(helmet());
app.use(helmet.referrerPolicy({ policy: 'same-origin' }));
app.use(express.static(path.join(__dirname, 'public')));

//check for keys in headers
app.use((req, res, next) => {
  if (!req.headers.lang || req.headers.lang == "en") {
    global.messages = require('./locales/en');
  }
  next();
});

app.use("/v1/students/", require("./src/version1/student/student-router"));
// catch 404 and forward to error handler
app.use(function(req, res, next) {
  next(createError(404));
});

// error handler
app.use(function(err, req, res, next) {
  // set locals, only providing error in development
  res.locals.message = err.message;
  res.locals.error = req.app.get('env') === 'development' ? err : {};
  // render the error page
  res.status(err.status || 500);
  // res.render('error');
  res.send({
    message: err.message,
    error: {}
  });
});

module.exports = app;
