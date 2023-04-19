const fs = require('fs');
const { promisify } = require('util');
const readFile = promisify(fs.readFile);
const writeFile = promisify(fs.writeFile);
const https = require('https');

const API_KEY = 'AIzaSyDWV3QGYTAf4ScolWz4EWvrpLBm3KGVJH0'; 
const getAddressLatLng = (address) => {
  
  return new Promise((resolve, reject) => {
    https.get(`https://maps.googleapis.com/maps/api/geocode/json?address=${encodeURI(address)}&key=${API_KEY}`, (res) => {
      let data = '';
      res.on('data', (chunk) => {
        data += chunk;
      });
      res.on('end', () => {
        const result = JSON.parse(data);
        if (result.status === 'OK') {
          const location = result.results[0].geometry.location;
          resolve({ lat: location.lat, lon: location.lng });
        } else {
          resolve(null);
        }
      });
    }).on('error', (err) => {
      reject(err);
    });
  });
};

const processFile = async (filename) => {
  const data = await readFile(filename, 'utf8');
  const lines = data.split('\n');
  let output = '';
  for (let i = 0; i < lines.length; i++) {
    const line = lines[i];
    if (line.includes('Property Address')) {
      const address = line.split(': ')[1];
      const latLon = await getAddressLatLng(address);
      if (latLon === null) {
        output += `${line}\n[No results found]\n`;
      } else {
        const { lat, lon } = latLon;
        output += `${line}\n[${lon}, ${lat}]\n`;
      }
    } else {
      output += line + '\n';
    }
  }
  await writeFile('latlonproperties.txt', output, 'utf8');
};

processFile('properties.txt');
