import Vue from "vue";
import Vuetify from "vuetify/lib";

import colors from "vuetify/lib/util/colors";

Vue.use(Vuetify);

export default new Vuetify({
  theme: {
    themes: {
      light: {
        primary: colors.blue.lighten1,
        secondary: colors.cyan.lighten1,
        accent: colors.cyan.accent2,
      },
      dark: {
        primary: colors.blue.darken4,
        secondary: colors.cyan.darken4,
        accent: colors.cyan.accent1,
      },
    },
  },
});
