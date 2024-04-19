import {
  layer,
  map,
  type NumberKeyValue,
  rule,
  withMapper,
  writeToProfile,
  hyperLayer,
  toApp,
} from 'karabiner.ts'

writeToProfile('Default', [
  rule('Caps Lock → Hyper').manipulators([
    map('caps_lock').toHyper().toIfAlone('escape'),
  ]),
  hyperLayer('o', 'hyper-o').manipulators({
    t: toApp('Warp'),
    v: toApp('Visual Studio Code - Insiders')
  }),
])
