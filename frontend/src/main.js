import Vue from "vue";
import App from "./App.vue";
import vuetify from "./plugins/vuetify";
import toasted from "./plugins/toasted";
import router from "./router";

Vue.config.productionTip = false;

new Vue({
  vuetify,
  toasted,
  router,
  render: (h) => h(App),
}).$mount("#app");
