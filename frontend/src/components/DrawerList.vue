<template>
  <v-list dense>
    <v-list-item link v-for="item in items" :key="item.text" v-bind:to="item.link">
      <v-list-item-action>
        <v-icon>{{item.icon}}</v-icon>
      </v-list-item-action>
      <v-list-item-content>
        <v-list-item-title>{{item.text}}</v-list-item-title>
      </v-list-item-content>
    </v-list-item>
  </v-list>
</template>

<script>
export default {
  data: () => ({
    items: [
      {
        link: "/",
        icon: "mdi-home",
        text: "开始"
      },
      {
        link: "/variables",
        icon: "mdi-variable",
        text: "变量"
      },
      {
        link: "/chart",
        icon: "mdi-chart-bell-curve-cumulative",
        text: "图表"
      },
      {
        link: "/about",
        icon: "mdi-information",
        text: "关于"
      }
    ]
  })
};
</script>
