const fs = require('fs');
const wasmBin = fs.readFileSync('./sort.wasm');

WebAssembly.instantiate(wasmBin).then((wasmMod) => {
  const { insertionSort, memory } = wasmMod.instance.exports;
  const myJSArray = [4, 3, 2, 1, 5];
  console.log(myJSArray.length);
  const myWASMArray = new Uint32Array(memory.buffer, 0, myJSArray.length);

  myWASMArray.set(myJSArray);

  console.log(myWASMArray);

  // calling WASM function (passing in memory)
  console.log('byte offset :::', myWASMArray.byteOffset);
  insertionSort(myWASMArray.byteOffset, myJSArray.length);

  console.log(myWASMArray);
});
