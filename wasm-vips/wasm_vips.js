const Vips = require('wasm-vips');
const fs = require('fs');

// /todo dumo csv

async function main() {
  const vips = await Vips();

  let im = vips.Image.newFromFile('../images/input.jpg');

  im = im.resize(0.0625);

  // Finally, write the result to a buffer
  const outBuffer = im.writeToBuffer('.jpg');

  fs.writeFileSync('../images/outputWasm.jpeg', outBuffer);
}

main().then((o) => console.log(o));
