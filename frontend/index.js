import * as esbuild from 'esbuild'
import * as sass from 'sass-embedded'
import { PurgeCSS } from 'purgecss'
import * as fs from 'node:fs'
import * as fsP from 'node:fs/promises'

await esbuild.build({
  entryPoints: ['./src/crypto.ts'],
  bundle: true,
  minify: true,
  platform: 'browser',
  outfile: '../release/assets/js/index.js',
})

console.log("ESBuild complete")

fsP.mkdir("./tmp")

const bootstrap = sass.compile("./scss/theme.scss", {silenceDeprecations: ['mixed-decls','color-functions','global-builtin', 'import']})

fs.writeFile("./tmp/bootstrap.css", bootstrap.css, function(err) {
    if(err) {
        fsP.rm("./tmp", {recursive: true,force: true})
        return console.log(err);
    }
}); 

const bootstrapPurged = await new PurgeCSS().purge({
    content: ['../views/*.templ', '../components/*.templ'],
    css: ['./tmp/bootstrap.css'],
  })

  fsP.rm("./tmp", {recursive: true,force: true})

  fs.writeFile("../release/assets/css/bootstrap.css", bootstrapPurged[0].css, function(err) {
    if(err) {
        return console.log(err);
    }

console.log("purgeCSS complete")

}); 