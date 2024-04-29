import { map, rule } from 'karabiner.ts'

export const hyperRule = rule('Caps Lock → Hyper').manipulators([
  map('caps_lock', undefined, ['shift', 'command'])
    .parameters({ 'basic.to_if_alone_timeout_milliseconds': 250 })
    .toHyper()
    .toIfAlone('escape'),
])

if (import.meta.main) {
  console.log(JSON.stringify(hyperRule.build(), null, 2))
}
