import {
  layer,
  map,
  type NumberKeyValue,
  rule,
  withMapper,
  writeToProfile,
  toApp,
  mapSimultaneous,
} from 'karabiner.ts'
import { hyperRule } from './rules'
import { hyperLayer } from './layers'

writeToProfile('Default', [
  hyperRule,

  rule('Delete forward')
    .description('If equals and backspace pressed together, delete forward.')
    .manipulators([mapSimultaneous(['equal_sign', 'delete_or_backspace']).to('delete_forward')]),

  hyperLayer,
])
