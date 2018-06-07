'use strict';
 
const http = require('http');
const {
  spawnSync
} = require('child_process');
 
http.createServer(function (req, res) {
let ls = spawnSync('bash', ['/var/vcap/jobs/keepalived/bin/get_keepalived_state.sh']);
   console.log(`stderr: ${ls.stderr.toString()}`);
   console.log(`stdout: ${ls.stdout.toString().trim()}`);
   console.log(`stdout: ${ls.stdout.toString().trim() === 'up'}`);
   if(ls.stdout.toString().trim() === 'up'){
     res.writeHead(200, {'Content-Type': 'text/html'});
     res.write('Up');
   }else{
     res.writeHead(500, {'Content-Type': 'text/html'});
     res.write('Down');
   }
  res.end();
}).listen(9595);
 
console.log('Listening on 9595');
