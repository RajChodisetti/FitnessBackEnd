const { spawn } = require('child_process');

function run(cmd, args) {
  const p = spawn(cmd, args, { stdio: 'inherit' });
  p.on('exit', (code) => {
    if (code !== null && code !== 0) {
      console.error(`${cmd} exited with code ${code}`);
    }
    process.exit(code);
  });
  return p;
}

const backend = run('go', ['run', './server']);
const frontend = run('next', ['dev']);

function shutdown() {
  backend.kill('SIGINT');
  frontend.kill('SIGINT');
}

process.on('SIGINT', shutdown);
process.on('SIGTERM', shutdown);

