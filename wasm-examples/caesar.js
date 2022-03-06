const fs = require('fs');
const wasmBuffer = fs.readFileSync('./caesar.wasm');

WebAssembly.instantiate(wasmBuffer).then((wasmModule) => {
  const { memory, caesarEncrypt, caesarDecrypt } = wasmModule.instance.exports;

  const plaintext = 'helloworld';
  const myKey = 3;

  const encode = function stringToIntegerArray(string, array) {
    const alphabet = 'abcdefghijklmnopqrstuvwxyz';
    for (let i = 0; i < string.length; i++) {
      array[i] = alphabet.indexOf(string[i]);
    }
  };

  const decode = function intergerArrayToString(array) {
    const alphabet = 'abcdefghijklmnopqrstuvwxyz';
    let string = '';
    for (let i = 0; i < array.length; i++) {
      string += alphabet[array[i]];
    }
    return string;
  };

  const myArray = new Int32Array(memory.buffer, 0, plaintext.length);

  // That second argument, 0, means our array begins at the very beginning of our shared memory.

  encode(plaintext, myArray);

  caesarEncrypt(myArray.byteOffset, myArray.length, myKey);

  console.log(myArray);
  console.log(decode(myArray));

  caesarDecrypt(myArray.byteOffset, myArray.length, myKey);
  console.log(myArray);
  console.log(decode(myArray));
});
