const fs = require('fs');
const crypto = require('crypto');
const execSync = require('child_process').execSync;

const indexPath = 'www/index.pcss';
const outputCssPath = 'www/index.css';
const hashFilePath = 'www/index_pcss.hash';

function calculateHash(filePath) {
  const content = fs.readFileSync(filePath, 'utf-8');
  return crypto.createHash('md5').update(content).digest('hex');
}

function updateHashFile(hash) {
  fs.writeFileSync(hashFilePath, hash);
}

if (fs.existsSync(hashFilePath)) {
  const storedHash = fs.readFileSync(hashFilePath, 'utf-8').trim();
  const currentHash = calculateHash(indexPath);
  if (storedHash === currentHash) {
    console.log('Hashes match. Skipping rebuild.');
  } else {
    console.log('Hashes don\'t match. Rebuilding CSS.')
    execSync(`tailwindcss -i ${indexPath} -o ${outputCssPath}`);
    console.log('CSS rebuild completed.');
    updateHashFile(currentHash);
  }
} else {
  console.log('Hash file not found. Rebuilding CSS.');
  execSync(`tailwindcss -i ${indexPath} -o ${outputCssPath}`);
  console.log('CSS rebuild completed.');
  const currentHash = calculateHash(indexPath);
  updateHashFile(currentHash);
}