import {
  layer,
  map,
  type NumberKeyValue,
  rule,
  withMapper,
  writeToProfile,
  hyperLayer,
  toApp,
  mapSimultaneous,
} from 'karabiner.ts'

writeToProfile('Default', [
  rule('Caps Lock → Hyper').manipulators([
    map('caps_lock')
      .parameters({ 'basic.to_if_alone_timeout_milliseconds': 250 })
      .toHyper()
      .toIfAlone('escape'),
  ]),

  rule('Delete forward')
    .description('If equals and backspace pressed together, delete forward.')
    .manipulators([mapSimultaneous(['equal_sign', 'delete_or_backspace']).to('delete_forward')]),

  hyperLayer('o', 'hyper-o').manipulators({
    t: toApp('Warp'),
    v: toApp('Visual Studio Code - Insiders'),
    s: toApp('Spotify'),
    m: toApp('Messages'),
  }),
])
