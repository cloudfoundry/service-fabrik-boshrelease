'use strict';

const http = require('http');
const exec = require('child_process').exec;

http.createServer(function (req, res) {
  exec('bash /var/vcap/jobs/keepalived/bin/get_keepalived_state.sh', function (error, stdout, stderr) {
    var status = 500;
    if (error !== null) {
      console.log('error', error);
      if (stderr) {
        console.log('stderr', stderr);
      }
      status = 500;
    } else {
      if (stdout !== null && stdout.toString().trim() === 'up') {
        status = 200;
      } else {
        status = 500;
      }
    }
    res.writeHead(status, {
      'Content-Type': 'text/plain'
    });
    if (status === 200) {
      res.write('Up');
    } else {
      res.write('Down');
    }
    res.end();
  });
}).listen(9595);

console.log(`${Date.now()}: Listening on 9595`);