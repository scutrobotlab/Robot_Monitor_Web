<template>
  <v-app id="inspire">
    <DrawerList ref="DrawerList" />

    <v-app-bar app dark color="primary">
      <v-app-bar-nav-icon @click.stop="switchDrawer()"></v-app-bar-nav-icon>
      <v-toolbar-title>坠好用的上位机</v-toolbar-title>
      <v-spacer></v-spacer>
      <v-switch v-model="$vuetify.theme.dark" hide-details inset color="black" label="深色模式"></v-switch>
    </v-app-bar>

    <v-content>
      <router-view></router-view>
    </v-content>

    <v-footer app color="primary">
      <span class="white--text">华工机器人实验室</span>
    </v-footer>
  </v-app>
</template>

<script>
import DrawerList from "@/components/DrawerList.vue";
export default {
  components: {
    DrawerList
  },
  props: {
    source: String
  },
  methods: {
    switchDrawer() {
      this.$refs.DrawerList.switchDrawer();
    }
  }
};
</script>
