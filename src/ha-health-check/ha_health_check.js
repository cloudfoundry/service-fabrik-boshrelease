'use strict';

const http = require('http');
const {
  spawnSync
} = require('child_process');

http.createServer(function (req, res) {
  const ls = spawnSync('bash', ['/var/vcap/jobs/keepalived/bin/get_keepalived_state.sh']);
  if (ls.stdout.toString().trim() === 'up') {
    res.writeHead(200, { 'Content-Type': 'text/html' });
    res.write('Up');
  } else {
    res.writeHead(500, { 'Content-Type': 'text/html' });
    res.write('Down');
  }
  res.end();
}).listen(9595);

console.log(`${Date.now}: Listening on 9595`);
