import { opine } from "https://deno.land/x/opine@2.0.0/mod.ts";

const app = opine()

app.get("/", function(_req, res) {
  res.send("Hello World")
})

app.listen(
  3000,
  () => console.log("server start")
)
