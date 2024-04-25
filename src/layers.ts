import { hyperLayer as karabinerHyperLayer, toApp } from 'karabiner.ts'

export const hyperLayer = karabinerHyperLayer('o', 'hyper-o').manipulators({
  t: toApp('Warp'),
  v: toApp('Visual Studio Code - Insiders'),
  s: toApp('Spotify'),
  m: toApp('Messages'),
  b: toApp('Arc'),
})
