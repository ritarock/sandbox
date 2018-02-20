<template>
  <div id="body">
    <h1>Mypage</h1>
    <hr />

    <div @click="clickLink" style="font-size:120%;">
      <md-button class="md-raised md-primary">{{link}}</md-button>
    </div>
    <span v-if="listLink" class="link-list">
      <span @mouseenter="qiitaShowIcon" @mouseleave="qiitaHiddenIcon" class="qiita-list">
        <a href="https://qiita.com/ritarock">qiita</a>
      </span>
      <span @mouseenter="githubShowIcon" @mouseleave="githubHiddenIcon" class="github-list">
        <a href="https://github.com/ritarock">github</a>
      </span>
      <span @mouseenter="hatenablogShowIcon" @mouseleave="hatenablogHiddenIcon" class="hatenablog-list">
        <a href="http://ritarock.hatenablog.com/">hatenablog</a>
      </span>
    </span>

    <div @click="clickWho" class="who-list" style="font-size:120%;">
      <md-button class="md-raised md-primary">{{who}}</md-button>
    </div>
    <span v-if="linkWho">
      <span class="who-name">NAME: ritarock</span>
      <span class="who-mail">MAIL: wkryz1204@gmail.com</span>
    </span>

    <span class="icon">
      <span class="qiita-icon" v-if="qiitaActiveIcon"><img src="../assets/qiita.png" width="400" height="400"></span>
      <span class="github-icon" v-if="githubActiveIcon"><img src="../assets/GitHub.png" width="400" height="400"></span>
      <span class="hatenablog-icon" v-if="hatenablogActiveIcon"><img src="../assets/hatena.png" width="400" height="400"></span>
    </span>

    <span id="footer"><a href="https://github.com/ritarock/sand-box/tree/master/sample_page">This page is made of vuejs</a></span>

  </div>
</template>

<script>
export default {
  name: 'BodyContent',
  data: function(){
    return{
      link: 'LINKS ▶︎',
      who: 'WHO? ▶︎',
      listLink: false,
      linkWho: false,
      qiitaActiveIcon: false,
      githubActiveIcon: false,
      hatenablogActiveIcon: false
    }
  },
  methods: {
    clickLink: function(){
      if (this.listLink) {
        this.link = 'LINKS ▶︎'
        this.listLink = false
      } else {
        this.link = 'LINKS ▼'
        this.listLink = true
      }
    },
    clickWho: function(){
      if (this.linkWho) {
        this.who = 'WHO? ▶︎'
        this. linkWho = false
      } else {
        this.who = 'WHO? ▼'
        this.linkWho = true
      }
    },
    qiitaShowIcon: function(){
      return this.qiitaActiveIcon = true
    },
    qiitaHiddenIcon: function(){
      return this.qiitaActiveIcon = false
    },
    githubShowIcon: function(){
      return this.githubActiveIcon = true
    },
    githubHiddenIcon: function(){
      return this.githubActiveIcon = false
    },
    hatenablogShowIcon: function(){
      return this.hatenablogActiveIcon = true
    },
    hatenablogHiddenIcon: function(){
      return this.hatenablogActiveIcon = false
    }
  }
}
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
