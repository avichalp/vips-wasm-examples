import { generate } from 'csv-generate';
import { stringify } from 'csv-stringify';

import { hrtime } from 'process';
import fs from 'fs';

const wasmBuffer = fs.readFileSync('./fibo.wasm');

function fibonacci(n, a = 0, b = 1) {
  return n < 1 ? a : fibonacci(n - 1, a + b, a);
}

WebAssembly.instantiate(wasmBuffer, {
  imports: {
    imported_func: function (arg) {
      console.log(arg);
    },
  },
}).then((wasmModule) => {
  const data = [];

  for (let i = 1; i <= 45; i++) {
    const jsStart = hrtime.bigint();
    fibonacci(i);
    const jsEnd = hrtime.bigint();

    const wasmStart = hrtime.bigint();
    wasmModule.instance.exports.fibonacci(i);
    const wasmEnd = hrtime.bigint();

    data.push([i, wasmEnd - wasmStart, jsEnd - jsStart]);
  }

  let columns = {
    request: 'request',
    wasm: 'wasm',
    js: 'js',
  };

  stringify(data, { header: true, columns: columns }, (err, output) => {
    if (err) throw err;
    fs.writeFile('fibo.csv', output, (err) => {
      if (err) throw err;
      console.log('fibo.csv saved.');
    });
  });
});
