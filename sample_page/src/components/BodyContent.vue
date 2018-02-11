<template>
  <div id="body">
    <h1>Welcome to Mypage</h1>
    <hr />
    <div @click="clickLink">{{link}}</div>
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

    <span class="icon">
      <span class="qiita-icon" v-if="qiitaActiveIcon"><img src="../assets/qiita.png" width="400" height="400"></span>
      <span class="github-icon" v-if="githubActiveIcon"><img src="../assets/GitHub.png" width="400" height="400"></span>
      <span class="hatenablog-icon" v-if="hatenablogActiveIcon"><img src="../assets/hatena.png" width="400" height="400"></span>
    </span>

    <div @click="clickWho" class="who-list">{{who}}</div>
    <span v-if="linkWho">
      <span class="who-name">NAME: ritarock</span>
      <span class="who-mail">MAIL: wkryz1204@gmail.com</span>
    </span>
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
  font-size: 100%;
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
}
.github-list {
  position: absolute;
  top: 20px;
  left: 10px;
}
.hatenablog-list {
  position: absolute;
  top: 40px;
  left: 10px;
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
  position: absolute;
  top: 200px;
  left: 0px;
}
.who-name {
  position: absolute;
  top: 225px;
  left: 10px;
}
.who-mail {
  position: absolute;
  top: 245px;
  left: 10px;
}

</style>
