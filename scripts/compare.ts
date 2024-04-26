import { $ } from 'bun'
import { watch } from 'chokidar'

import { hyperRule } from '../src/rules'
import { hyperLayer } from '../src/layers'

const tsFile = `out_ts.json`
const goFile = `out_go.json`

const generateTS = () => {
  const data = [hyperRule, hyperLayer].map((r) => r.build())
  return Bun.write(tsFile, JSON.stringify(data, null, 2))
}

const generateGo = async () => Bun.write(goFile, await $`go run main.go`.text())

const watcher = watch([`src/**/*.ts`, `cmd/**/*.go`, `internal/**/*.go`], {
  awaitWriteFinish: true,
  ignoreInitial: true,
})

watcher.on('all', async (event, path) => {
  if (path.endsWith('.ts')) {
    console.log(`Generating TypeScript`)
    await generateTS()
  } else if (path.endsWith('.go')) {
    console.log(`Generating Go`)
    await generateGo()
  }
})

await generateTS()
await generateGo()

await $`ksdiff -l karabiner-config ${tsFile} ${goFile}`
