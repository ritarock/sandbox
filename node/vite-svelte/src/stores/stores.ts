import { writable, derived } from 'svelte/store'

export const input = writable('')

export const output = derived(
  input,
  $input => `${$input}`
)
