import Vips from 'wasm-vips';
import fs from 'fs';
import { hrtime } from 'process';

async function main() {
  const vips = await Vips();

  const inputs = [
    { path: '../images/input.jpg', factor: 0.5, format: '.jpeg' },
    {
      path: '../images/input.jpg',
      factor: 0.1,
      format: '.jpeg',
    },
    {
      path: '../images/input.jpg',
      factor: 0.05,
      format: '.jpeg',
    },
    {
      path: '../images/input.jpg',
      factor: 0.01,
      format: '.jpeg',
    },
    { path: '../images/input.png', factor: 0.5, format: '.png' },
    {
      path: '../images/input.png',
      factor: 0.1,
      format: '.png',
    },
    {
      path: '../images/input.png',
      factor: 0.05,
      format: '.png',
    },
    {
      path: '../images/input.png',
      factor: 0.01,
      format: '.png',
    },
    { path: '../images/input.webp', factor: 0.5, format: '.webp' },
    {
      path: '../images/input.webp',
      factor: 0.1,
      format: '.webp',
    },
    {
      path: '../images/input.webp',
      factor: 0.05,
      format: '.webp',
    },
    {
      path: '../images/input.webp',
      factor: 0.01,
      format: '.webp',
    },
  ];

  const jpegLatencies = [];
  const pngLatencies = [];
  const webpLatencies = [];

  for (const input of inputs) {
    const latencies = [];

    for (let i = 0; i < 20; i++) {
      const start = hrtime.bigint();

      let im = vips.Image.newFromFile(input.path);

      im = im.resize(input.factor);
      const outBuffer = im.writeToBuffer(input.format);
      fs.writeFileSync(`../images/output.wasm_vips.${input.factor}${input.format}`, outBuffer);
      latencies.push(Number(hrtime.bigint() - start) / 1000000);
    }

    if (input.format === '.jpeg') {
      jpegLatencies.push(latencies.reduce((l1, l2) => l1 + l2) / latencies.length);
    } else if (input.format === '.png') {
      pngLatencies.push(latencies.reduce((l1, l2) => l1 + l2) / latencies.length);
    } else {
      webpLatencies.push(latencies.reduce((l1, l2) => l1 + l2) / latencies.length);
    }
  }
  console.log('JPEG', jpegLatencies);
  console.log('PNG', pngLatencies);
  console.log('WEBP', webpLatencies);
}

main().then((o) => console.log(o));
