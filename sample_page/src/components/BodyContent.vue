<template>
  <div id="body">
    <h1>メンテナンス中</h1>
    <hr />

    <span id="footer"><a href="https://github.com/ritarock/sand-box/tree/master/sample_page">This page is made of vuejs</a></span>

  </div>
</template>

<script>
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>

#links {
  color: black;
}
h1{
  font-size: 120%;
  font-weight: normal;
  text-align: center;
}
.link-list {
  position: relative;
  left: 0px;
  background-color:red;
}
.qiita-list {
  position: absolute;
  top: 0px;
  left: 10px;
  font-size:120%;
}
.github-list {
  position: absolute;
  top: 20px;
  left: 10px;
  font-size:120%;
}
.hatenablog-list {
  position: absolute;
  top: 40px;
  left: 10px;
  font-size:120%;
}
.qiita-icon{
  position: absolute;
  top: 100px;
  left: 450px
}
.github-icon{
  position: absolute;
  top: 100px;
  left: 450px
}
.hatenablog-icon{
  position: absolute;
  top: 120px;
  left: 450px
}
.who-list {
  position: relative;
  top: 200px;
  left: 0px;
}
.who-name {
  position: absolute;
  top: 300px;
  left: 10px;
  font-size:120%;
}
.who-mail {
  position: absolute;
  top: 320px;
  left: 10px;
  font-size:120%;
}
body , html {
  height: 100%;
}
#container {
  width: 100%;
  position: relative;
  height: auto !important;
  height: 100%;
  min-height: 100%;
}
#contents {
  padding-bottom: 100px;
}
#footer {
  position: fixed;
  bottom: 0;
  width: 100%;
  height: 100px;
  font-size: 85%;
  font-weight: normal;
  text-align: center;
}
.md-raised {
  color: gray;
}

</style>
