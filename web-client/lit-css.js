import litcss from 'rollup-plugin-lit-css';
// import config from './rollup.config.js'
import aliasPlugin from '@rollup/plugin-alias';

import { dirname, resolve } from 'node:path';
import { fileURLToPath } from 'node:url';
import { rollup } from 'rollup';

// export default {
//   ...config,
//   plugins: [
//     litcss({ include: ['node_modules/@material/typography/dist/mdc.typography.css'], exclude: undefined }
//     )
//   ]
// }




const additionalPlugins = [ []]; // ...alias ? [aliasPlugin({ entries: alias })] :

const input = 'node_modules/@material/typography/dist/mdc.typography.css';

const bundle = await rollup({
  input,
  external: ['lit'],
  plugins: [
    litcss({ uglify: false }),
    ...additionalPlugins,
  ],
});

const { output: [{ code }] } = await bundle.generate({ format: 'es' });
console.log(code);
