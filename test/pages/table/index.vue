<template>
  <div>
    <h1>table page</h1>
    <button v-on:click="fetch">fetch</button>
    COUNT: {{ length }}
    <table border="1">
      <tr>
        <th>age</th>
        <th>count</th>
      </tr>
      <tr v-for="item in items" :key="item.age">
        <th>{{ item.age }}</th>
        <th>{{ item.users }}</th>
      </tr>
    </table>
  </div>
</template>

<script lang="ts">
import Vue from 'vue'

function getJson() {
  const ages = Array.from(
    new Set(
      new Array(parseInt(Math.random() * 100))
          .fill(null)
          .map(() => parseInt(Math.random() * 100))
    )
  )

  return ages.map(age => ({
    age,
    users: Math.random() * 1000
  }))
}


export default Vue.extend({
  data: () => ({
    items: []
  }),

  computed: {
    length() {
      return this.items.length
    }
  },

  created() {
    this.fetch()
  },

  methods: {
    fetch() {
      this.items = getJson()
    }
  }
})
</script>