<script lang="ts">
import { input, output } from "../stores/stores";
let result = "";
let indentSize = 2;

function handleClick() {
  const input = $output;
  result = converter(input, 1);
}

function converter(input: string, indent: number) {
  let output = "";
  indentSize = indent * 2;
  const jsonObject = JSON.parse(input);
  const keys = Object.keys(jsonObject);
  for (let key in keys) {
    let name = keys[key];
    let value = jsonObject[keys[key]];
    let arrFlg = (value instanceof Array);
    let objFlg = (value instanceof Object);

    if (arrFlg) {
      output += " ".repeat(indentSize) + `- ${name}:\n`;
      value.forEach((e) => {
        output += " ".repeat(indentSize * 2) + `- '${e}'\n`;
      });
      continue;
    }
    if (objFlg) {
      output += " ".repeat(indentSize) + `- ${name}:\n`;
      let tmp = indentSize;
      output += converter(JSON.stringify(value), 2);
      indentSize = tmp;
      continue;
    }
    output += " ".repeat(indentSize) + `- ${name}:\n`;
    output += " ".repeat(indentSize * 2) + `- '${value}'\n`;
  }
  return output;
}

</script>

<div>
  <p>json to yaml</p>
  <textarea rows="10" cols="50" bind:value={$input}></textarea>
  <textarea rows="10" cols="50" value={result}></textarea>
  <p>
    <button on:click={handleClick}>convert</button>
  </p>
</div>
